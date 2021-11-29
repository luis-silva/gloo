package tg

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	envoyapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoycore "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type plugin struct{}

func NewPlugin() *plugin {
	return &plugin{}
}

func (*plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoyapi.Cluster) error {
	// check that the upstream is our type (GCE)
	if _, ok := in.UpstreamType.(*v1.Upstream_Tg); !ok {
		// not gce, return early
		return nil
	}
	// tell Envoy to use EDS to get endpoints for this cluster
	out.ClusterDiscoveryType = &envoyapi.Cluster_Type{
		Type: envoyapi.Cluster_EDS,
	}
	// tell envoy to use ADS to resolve Endpoints
	out.EdsClusterConfig = &envoyapi.Cluster_EdsClusterConfig{
		EdsConfig: &envoycore.ConfigSource{
			ConfigSourceSpecifier: &envoycore.ConfigSource_Ads{
				Ads: &envoycore.AggregatedConfigSource{},
			},
		},
	}
	return nil
}

func (*plugin) WatchEndpoints(writeNamespace string, upstreamsToTrack v1.UpstreamList, opts clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {
	// use the context from the opts we were passed
	ctx := opts.Ctx

	// get the client for interacting with AWS ELBs
	elbv2Client, err := initializeClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	// initialize the channel on which we will send endpoint results to Gloo Edge
	results := make(chan v1.EndpointList)

	// initialize a channel on which we can send polling errors to Gloo Edge
	errorsDuringUpdate := make(chan error)

	// in a goroutine, continue updating endpoints at an interval
	// until the context is done
	go func() {
		// once this goroutine exits, we should close our output channels
		defer close(results)
		defer close(errorsDuringUpdate)

		// poll indefinitely
		for {
			select {
			case <-ctx.Done():
				// context was cancelled, stop polling
				return
			default:
				endpoints, err := getLatestEndpoints(elbv2Client, upstreamsToTrack)
				if err != nil {
					// send the error to Gloo Edge for logging
					errorsDuringUpdate <- err
				} else {
					// send the latest set of endpoints to Gloo Edge
					results <- endpoints
				}

				// sleep 10s between polling
				time.Sleep(time.Second * 10)
			}
		}
	}()

	// return the channels to Gloo Edge
	return results, errorsDuringUpdate, nil
}

// it is sufficient to return nil here
func (*plugin) Init(params plugins.InitParams) error {
	return nil
}

// though required by the plugin interface, this function is not necesasary for our plugin
func (*plugin) DiscoverUpstreams(watchNamespaces []string, writeNamespace string, opts clients.WatchOpts, discOpts discovery.Opts) (chan v1.UpstreamList, chan error, error) {
	return nil, nil, nil
}

// though required by the plugin interface, this function is not necesasary for our plugin
func (*plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	return false, nil
}

// initialize client for talking to Google Compute Engine API
func initializeClient(ctx context.Context) (*elbv2.ELBV2, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return elbv2.New(sess), nil
}

// one call results in a list of endpoints for our upstreams
func getLatestEndpoints(instancesClient *elbv2.ELBV2, upstreams v1.UpstreamList) (v1.EndpointList, error) {

	// initialize a new list of endpoints
	var result v1.EndpointList

	// for each upstream, retrieve its endpoints
	// for _, us := range upstreams {
	//   // check that the upstream uses the GCE Spec
	// 	gceSpec := us.GetGce()
	// 	if gceSpec == nil {
	// 		// skip ELB TG upstreams
	// 		continue
	// 	}

	// 	// get the Google Compute VM Instances for the project/zone
	// 	instancesForUpstream, err := instancesClient.List(
	// 		gceSpec.ProjectId,
	// 		gceSpec.Zone,
	// 	).Do()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	// iterate over each instance
	// 	// add its address as an endpoint if its labels match
	// 	for _, instance := range instancesForUpstream.Items {

	// 		if !shouldSelectInstance(gceSpec.Selector, instance.Labels) {
	// 			// the selector doesn't match this instance, skip it
	// 			continue
	// 		}

	// 		if len(instance.NetworkInterfaces) == 0 {
	// 			// skip vms that don't have an allocated IP address
	// 			continue
	// 		}

	// 		// use the first network ip of the vm for our endpoint
	// 		address := instance.NetworkInterfaces[0].NetworkIP

	// 		// get the port from the upstream spec
	// 		port := gceSpec.Port

	// 		// provide a pointer back to the upstream this
	// 		// endpoint was created for
	// 		upstreamRef := us.Metadata.Ref()

	// 		endpointForInstance := &v1.Endpoint{
	// 			Metadata: core.Metadata{
	// 				Namespace: us.Metadata.Namespace,
	// 				Name:      instance.Name,
	// 				Labels:    instance.Labels,
	// 			},
	// 			Address:   address,
	// 			Port:      port,
	// 			// normally if more than one upstream shares an endpoint
	// 			// we would provide a list here
	// 			Upstreams: []*core.ResourceRef{&upstreamRef},
	// 		}

	// 		// add the endpoint to our list
	// 		result = append(result, endpointForInstance)
	// }

	return result, nil
}

// inspect the labels for a match
func shouldSelectInstance(selector, instanceLabels map[string]string) bool {
	if len(instanceLabels) == 0 {
		// only an empty selector can match empty labels
		return len(selector) == 0
	}

	for k, v := range selector {
		instanceVal, ok := instanceLabels[k]
		if !ok {
			// the selector key is missing from the instance labels
			return false
		}
		if v != instanceVal {
			// the label value in the selector does not match
			// the label value from the instance
			return false
		}
	}
	// we didn't catch a mismatch by now, they match
	return true
}
