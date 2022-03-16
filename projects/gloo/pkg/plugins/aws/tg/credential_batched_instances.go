package tg

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	glootg "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/tg"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

const AvailabilityZoneAnnotationKey = "AvailabilityZone"

// In order to minimize calls to the AWS API, we group calls by credentials and apply tag filters locally.
// This function groups upstreams by credentials, calls the AWS API, maps the TargetGroups to upstreams, and returns the
// endpoints associated with the provided upstream list
// NOTE: MUST filter the upstreamList to ONLY TargetGroup upstreams before calling this function
func getLatestEndpoints(ctx context.Context, lister TargetGroupLister, secrets v1.SecretList, writeNamespace string, upstreamList v1.UpstreamList) (v1.EndpointList, error) {
	// we want unique creds so we can query api once per unique cred
	// we need to make sure we maintain the association between those unique creds and the upstreams that share them
	// so that when we get the TargetGroups associated with the creds, we will know which upstreams have access to those
	// TargetGroups.
	credGroups, err := getCredGroupsFromUpstreams(upstreamList)
	if err != nil {
		return nil, err
	}
	// call the ELBV2 APIs once for each set of credentials and apply the output to the credential groups
	if err := getTargetGroupsForCredentialGroups(ctx, lister, secrets, credGroups); err != nil {
		return nil, err
	}
	// produce the endpoints list
	var allEndpoints v1.EndpointList
	for _, credGroup := range credGroups {
		for _, upstream := range credGroup.upstreams {
			targetsForUpstream := filterTargetGroupsForUpstreams(ctx, lister, upstream, credGroup, secrets)
			for _, target := range targetsForUpstream {
				if endpoint := upstreamTargetsToEndpoint(ctx, writeNamespace, upstream, target); endpoint != nil {
					allEndpoints = append(allEndpoints, endpoint)
				}
			}
		}
	}
	return allEndpoints, nil
}

// credentialGroup exists to support batched calls to the AWS API
// one credentialGroup should be made for each unique credentialSpec
type credentialGroup struct {
	// a unique credential spec
	credentialSpec *CredentialSpec
	// all the upstreams that share the CredentialSpec
	upstreams v1.UpstreamList
	// all the targets visible to the given credentials
	targets []*ExtendedTargetGroup
	// one filter map exists for each targetgroup in order to support client-side filtering
	filterMaps []FilterMap
}

// Initializes the credentialGroups
// Credential groups are returned as a map to enforce the "one credentialGroup per unique credential" property that is
// required in order to realize the benefits of batched AWS API calls.
// NOTE: assumes that upstreams are TargetGroups upstreams
func getCredGroupsFromUpstreams(upstreams v1.UpstreamList) (map[CredentialKey]*credentialGroup, error) {
	credGroups := make(map[CredentialKey]*credentialGroup)
	for _, upstream := range upstreams {
		cred := NewCredentialSpecFromTargetGroupUpstreamSpec(upstream.GetAwsTg())
		key := cred.GetKey()
		if _, ok := credGroups[key]; ok {
			credGroups[key].upstreams = append(credGroups[key].upstreams, upstream)
		} else {
			credGroups[key] = &credentialGroup{
				upstreams:      v1.UpstreamList{upstream},
				credentialSpec: cred,
			}
		}
	}
	return credGroups, nil
}

// calls the AWS API and attaches the output to the the provided list of credentialGroups. Modifications include:
// - adds the TargetGroups for each credentialGroup's credential
// - adds tag filters for each TargetGroup for later use when refining the list of TargetGroups that an upstream has
// permission to describe to the list of targets that the upstream should route to
func getTargetGroupsForCredentialGroups(ctx context.Context, lister TargetGroupLister, secrets v1.SecretList, credGroups map[CredentialKey]*credentialGroup) error {
	for _, credGroup := range credGroups {
		targets, err := lister.ListForCredentials(ctx, credGroup.credentialSpec, secrets)
		if err != nil {
			return err
		}
		credGroup.targets = targets
		credGroup.filterMaps = generateFilterMaps(targets)
	}
	return nil
}

// applies filter logic equivalent to the tag filter logic used in AWS's DescribeInstances API
// NOTE: assumes that upstreams are TargetGroup upstreams
func filterTargetGroupsForUpstreams(ctx context.Context, lister TargetGroupLister, upstream *v1.Upstream, credGroup *credentialGroup, secrets v1.SecretList) []Target {
	var targets []Target
	logger := contextutils.LoggerFrom(ctx)
	// sweep through each filter map, if all the upstream's filters are matched, add the corresponding target to the list
	for i, fm := range credGroup.filterMaps {
		candidateTarget := credGroup.targets[i]
		logger.Debugw("considering targetgroup for upstream", "upstream", upstream.GetMetadata().Ref().Key(), "tags", candidateTarget.Tags, "tg-arn", candidateTarget.TargetGroupArn)
		matchesAll := true
	ScanFilters: // label so that we can break out of the for loop rather than the switch
		for _, filter := range upstream.GetAwsTg().GetFilters() {
			switch filterSpec := filter.GetSpec().(type) {
			case *glootg.TagFilter_Key:
				if _, ok := fm[awsKeyCase(filterSpec.Key)]; !ok {
					matchesAll = false
					break ScanFilters
				}
			case *glootg.TagFilter_KvPair_:
				if val, ok := fm[awsKeyCase(filterSpec.KvPair.GetKey())]; !ok || val != filterSpec.KvPair.GetValue() {
					matchesAll = false
					break ScanFilters
				}
			}
		}
		if matchesAll {
			svc, err := GetELBV2Client(credGroup.credentialSpec, secrets)
			if err != nil {
				return nil
			}
			endpoints, err := lister.ListHealthyTasks(candidateTarget.TargetGroupArn, ctx, svc)
			if err != nil {
				return nil
			}

			targets = append(targets, endpoints...)

			logger.Debugw("target for upstream accepted", "upstream", upstream.GetMetadata().Ref().Key(), "tags", candidateTarget.Tags)
		} else {
			logger.Debugw("target for upstream filtered out", "upstream", upstream.GetMetadata().Ref().Key(), "tags", candidateTarget.Tags)
		}
	}
	return targets
}

// NOTE: assumes that upstreams are targetgroup upstreams
func upstreamTargetsToEndpoint(ctx context.Context, writeNamespace string, upstream *v1.Upstream, target Target) *v1.Endpoint {
	ipAddr := target.Ip
	if ipAddr == nil {
		contextutils.LoggerFrom(ctx).Warnw("no ip found for config",
			zap.Any("upstreamRef", upstream.GetMetadata().Ref()))
		return nil
	}
	port := target.Port
	ref := upstream.GetMetadata().Ref()
	// for easier debugging, add the target availability zone to the xds output
	targetInfo := make(map[string]string)
	targetInfo[AvailabilityZoneAnnotationKey] = aws.StringValue(target.Az)
	endpoint := v1.Endpoint{
		Upstreams: []*core.ResourceRef{ref},
		Address:   aws.StringValue(ipAddr),
		Port:      uint32(*port),
		Metadata: &core.Metadata{
			Name:        generateName(ref, aws.StringValue(ipAddr)),
			Namespace:   writeNamespace,
			Annotations: targetInfo,
		},
	}
	contextutils.LoggerFrom(ctx).Debugw("target from upstream",
		zap.Any("upstream", upstream),
		zap.Any("target", target),
		zap.Any("endpoint", endpoint))
	return &endpoint
}

// a FilterMap is created for each TargetGroup so we can efficiently filter the TargetGroups associated with a given
// upstream's filter spec
// filter maps are generated from tag lists, the keys are the tag keys, the values are the tag values
type FilterMap map[string]string

func generateFilterMap(target *ExtendedTargetGroup) FilterMap {
	m := make(FilterMap)
	for _, t := range target.Tags {
		m[awsKeyCase(aws.StringValue(t.Key))] = aws.StringValue(t.Value)
	}
	return m
}

func generateFilterMaps(targets []*ExtendedTargetGroup) []FilterMap {
	var maps []FilterMap
	for _, target := range targets {
		maps = append(maps, generateFilterMap(target))
	}
	return maps
}

// AWS tag keys are not case-sensitive so cast them all to lowercase
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-policy-structure.html#amazon-ec2-keys
func awsKeyCase(input string) string {
	return strings.ToLower(input)
}
