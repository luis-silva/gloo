package tg

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/aws/aws-sdk-go/service/elbv2"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

// TargetGroupLister is a simple interface for calling the AWS API.
// This allows us to easily mock the API in our tests.
type TargetGroupLister interface {
	ListForCredentials(ctx context.Context, cred *CredentialSpec, secrets v1.SecretList) ([]*ExtendedTargetGroup, error)
	ListHealthyTasks(arn *string, ctx context.Context, svc *elbv2.ELBV2) ([]Task, error)
}

type targetGroupLister struct {
}

func NewTargetGroupLister() *targetGroupLister {
	return &targetGroupLister{}
}

var _ TargetGroupLister = &targetGroupLister{}

type Task struct {
	Ip   *string
	Port *int64
	Az   *string
}

type ExtendedTargetGroup struct {
	*elbv2.TargetGroup
	Tags  []*elbv2.Tag
	Tasks []*Task
}

func (c *targetGroupLister) ListForCredentials(ctx context.Context, cred *CredentialSpec, secrets v1.SecretList) ([]*ExtendedTargetGroup, error) {
	svc, err := GetELBV2Client(cred, secrets)
	if err != nil {
		return nil, GetClientError(err)
	}

	tgs, err := c.ListWithClient(ctx, svc)
	if err != nil {
		return nil, err
	}

	tds, err := c.ListTagsWithClient(tgs, ctx, svc)
	if err != nil {
		return nil, err
	}

	tags := c.ConvertTagDescriptionsToMap(tds)

	return c.MergeTargetGroupTags(tgs, tags), nil
}

func (c *targetGroupLister) ListWithClient(ctx context.Context, svc *elbv2.ELBV2) ([]*elbv2.TargetGroup, error) {

	var results []*elbv2.DescribeTargetGroupsOutput
	// pass a filter to only get running instances.
	input := &elbv2.DescribeTargetGroupsInput{}
	err := svc.DescribeTargetGroupsPagesWithContext(ctx, input, func(r *elbv2.DescribeTargetGroupsOutput, more bool) bool {
		results = append(results, r)
		return true
	})
	if err != nil {
		return nil, DescribeInstancesError(err)
	}

	var result []*elbv2.TargetGroup
	for _, dio := range results {
		result = append(result, GetInstancesFromDescription(dio)...)
	}

	contextutils.LoggerFrom(ctx).Debugw("ec2Upstream result", zap.Any("value", result))
	return result, nil
}

func (c *targetGroupLister) ListTagsWithClient(tgs []*elbv2.TargetGroup, ctx context.Context, svc *elbv2.ELBV2) ([]*elbv2.TagDescription, error) {

	var queryarns []*string
	var arns []*string
	var tagDescriptions []*elbv2.TagDescription

	for _, tg := range tgs {
		arns = append(arns, tg.TargetGroupArn)
	}

	for len(arns) > 0 {
		if len(arns) > 20 {
			queryarns = arns[0:20]
			arns = arns[20:]
		} else {
			queryarns = arns
			arns = []*string{}
		}

		// var results []*elbv2.DescribeTagsOutput
		// pass a filter to only get running instances.
		input := &elbv2.DescribeTagsInput{
			ResourceArns: queryarns,
		}
		results, err := svc.DescribeTagsWithContext(ctx, input)
		if err != nil {
			return nil, DescribeInstancesError(err)
		}

		tagDescriptions = append(tagDescriptions, results.TagDescriptions...)

		contextutils.LoggerFrom(ctx).Debugw("ec2Upstream result", zap.Any("value", results))
	}

	return tagDescriptions, nil
}

func (c *targetGroupLister) ConvertTagDescriptionsToMap(tds []*elbv2.TagDescription) map[string][]*elbv2.Tag {

	var tagMap map[string][]*elbv2.Tag
	for _, td := range tds {
		tagMap[*td.ResourceArn] = td.Tags
	}

	return tagMap
}

func (c *targetGroupLister) ListHealthyTasks(arn *string, ctx context.Context, svc *elbv2.ELBV2) ([]Task, error) {

	var healthyTargets []Task
	input := &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: arn,
	}
	results, err := svc.DescribeTargetHealthWithContext(ctx, input)
	if err != nil {
		return nil, DescribeInstancesError(err)
	}

	for _, result := range results.TargetHealthDescriptions {
		if *result.TargetHealth.State == elbv2.TargetHealthStateEnumHealthy {
			target := Task{
				Ip:   result.Target.Id,
				Port: result.Target.Port,
				Az:   result.Target.AvailabilityZone,
			}
			healthyTargets = append(healthyTargets, target)
		}
	}

	return healthyTargets, nil
}

func (c *targetGroupLister) MergeTargetGroupTags(tgs []*elbv2.TargetGroup, tags map[string][]*elbv2.Tag) []*ExtendedTargetGroup {

	var etgs []*ExtendedTargetGroup
	for _, tg := range tgs {
		etgs = append(etgs, &ExtendedTargetGroup{tg, tags[*tg.TargetGroupArn], nil})
	}

	return etgs
}

var (
	GetClientError = func(err error) error {
		return eris.Wrapf(err, "unable to get aws client")
	}

	DescribeInstancesError = func(err error) error {
		return eris.Wrapf(err, "unable to describe instances")
	}
)
