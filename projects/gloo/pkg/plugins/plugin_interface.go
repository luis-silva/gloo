package plugins

import (
	"bytes"
	"context"
	"sort"
	"strings"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
)

type InitParams struct {
	Ctx      context.Context
	Settings *v1.Settings
}

type Plugin interface {
	Name() string
	Init(params InitParams) error
}

type Params struct {
	Ctx      context.Context
	Snapshot *v1snap.ApiSnapshot
}

type VirtualHostParams struct {
	Params
	Proxy        *v1.Proxy
	Listener     *v1.Listener
	HttpListener *v1.HttpListener
}

type RouteParams struct {
	VirtualHostParams
	VirtualHost *v1.VirtualHost
}

type RouteActionParams struct {
	RouteParams
	Route *v1.Route
}

/*
	Upstream Plugins
*/

// UpstreamPlugin is called after the envoy Cluster has been created for the input Upstream, and allows
// the cluster to be edited before being sent to envoy via CDS
type UpstreamPlugin interface {
	Plugin
	ProcessUpstream(params Params, in *v1.Upstream, out *envoy_config_cluster_v3.Cluster) error
}

// Endpoint is called after the envoy ClusterLoadAssignment has been created for the input Upstream, and allows
// the endpoints to be edited before being sent to envoy via EDS
// If one wishes to also modify the corresponding envoy Cluster the above UpstreamPlugin interface should be used.
type EndpointPlugin interface {
	Plugin
	ProcessEndpoints(params Params, in *v1.Upstream, out *envoy_config_endpoint_v3.ClusterLoadAssignment) error
}

/*
	Routing Plugins
*/

type RoutePlugin interface {
	Plugin
	ProcessRoute(params RouteParams, in *v1.Route, out *envoy_config_route_v3.Route) error
}

// note: any route action plugin can be implemented as a route plugin
// suggestion: if your plugin requires configuration from a RoutePlugin field, implement the RoutePlugin interface
type RouteActionPlugin interface {
	Plugin
	ProcessRouteAction(params RouteActionParams, inAction *v1.RouteAction, out *envoy_config_route_v3.RouteAction) error
}

type WeightedDestinationPlugin interface {
	Plugin
	ProcessWeightedDestination(
		params RouteParams,
		in *v1.WeightedDestination,
		out *envoy_config_route_v3.WeightedCluster_ClusterWeight,
	) error
}

/*
	Listener Plugins
*/

type ListenerPlugin interface {
	Plugin
	ProcessListener(params Params, in *v1.Listener, out *envoy_config_listener_v3.Listener) error
}

type StagedNetworkFilter struct {
	NetworkFilter *envoy_config_listener_v3.Filter
	Stage         FilterStage
}

type StagedNetworkFilterList []StagedNetworkFilter

func (s StagedNetworkFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, and (to ensure stability) index
func (s StagedNetworkFilterList) Less(i, j int) bool {
	switch FilterStageComparison(s[i].Stage, s[j].Stage) {
	case -1:
		return true
	case 1:
		return false
	}
	if s[i].NetworkFilter.GetName() < s[j].NetworkFilter.GetName() {
		return true
	}
	if s[i].NetworkFilter.GetName() > s[j].NetworkFilter.GetName() {
		return false
	}
	if s[i].NetworkFilter.String() < s[j].NetworkFilter.String() {
		return true
	}
	if s[i].NetworkFilter.String() > s[j].NetworkFilter.String() {
		return false
	}
	// ensure stability
	return i < j
}

func (s StagedNetworkFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type TcpFilterChainPlugin interface {
	Plugin
	CreateTcpFilterChains(params Params, parentListener *v1.Listener, in *v1.TcpListener) ([]*envoy_config_listener_v3.FilterChain, error)
}

// HttpConnectionManager Plugins
type HttpConnectionManagerPlugin interface {
	Plugin
	ProcessHcmNetworkFilter(params Params, parentListener *v1.Listener, listener *v1.HttpListener, out *envoyhttp.HttpConnectionManager) error
}

type HttpFilterPlugin interface {
	Plugin
	HttpFilters(params Params, listener *v1.HttpListener) ([]StagedHttpFilter, error)
}

type VirtualHostPlugin interface {
	Plugin
	ProcessVirtualHost(params VirtualHostParams, in *v1.VirtualHost, out *envoy_config_route_v3.VirtualHost) error
}

type StagedHttpFilter struct {
	HttpFilter *envoyhttp.HttpFilter
	Stage      FilterStage
}

type StagedHttpFilterList []StagedHttpFilter

func (s StagedHttpFilterList) Len() int {
	return len(s)
}

// filters by Relative Stage, Weighting, Name, Config Type-Url, Config Value, and (to ensure stability) index.
// The assumption is that if two filters are in the same stage, their order doesn't matter, and we
// just need to make sure it is stable.
func (s StagedHttpFilterList) Less(i, j int) bool {
	if compare := FilterStageComparison(s[i].Stage, s[j].Stage); compare != 0 {
		return compare < 0
	}

	if compare := strings.Compare(s[i].HttpFilter.GetName(), s[j].HttpFilter.GetName()); compare != 0 {
		return compare < 0
	}

	if compare := strings.Compare(s[i].HttpFilter.GetTypedConfig().GetTypeUrl(), s[j].HttpFilter.GetTypedConfig().GetTypeUrl()); compare != 0 {
		return compare < 0
	}

	if compare := bytes.Compare(s[i].HttpFilter.GetTypedConfig().GetValue(), s[j].HttpFilter.GetTypedConfig().GetValue()); compare != 0 {
		return compare < 0
	}

	// ensure stability
	return i < j
}

func (s StagedHttpFilterList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = StagedHttpFilterList{}

// WellKnownFilterStages are represented by an integer that reflects their relative ordering
type WellKnownFilterStage int

// If new well known filter stages are added, they should be inserted in a position corresponding to their order
const (
	FaultStage     WellKnownFilterStage = iota // Fault injection // First Filter Stage
	CorsStage                                  // Cors stage
	WafStage                                   // Web application firewall stage
	AuthNStage                                 // Authentication stage
	AuthZStage                                 // Authorization stage
	RateLimitStage                             // Rate limiting stage
	AcceptedStage                              // Request passed all the checks and will be forwarded upstream
	OutAuthStage                               // Add auth for the upstream (i.e. aws λ)
	RouteStage                                 // Request is going to upstream // Last Filter Stage
)

type FilterStage struct {
	RelativeTo WellKnownFilterStage
	Weight     int
}

// FilterStageComparison helps implement the sort.Interface Less function for use in other implementations of sort.Interface
// returns -1 if less than, 0 if equal, 1 if greater than
// It is not sufficient to return a Less bool because calling functions need to know if equal or greater when Less is false
func FilterStageComparison(a, b FilterStage) int {
	if a.RelativeTo < b.RelativeTo {
		return -1
	} else if a.RelativeTo > b.RelativeTo {
		return 1
	}
	if a.Weight < b.Weight {
		return -1
	} else if a.Weight > b.Weight {
		return 1
	}
	return 0
}

func BeforeStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, -1)
}
func DuringStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, 0)
}
func AfterStage(wellKnown WellKnownFilterStage) FilterStage {
	return RelativeToStage(wellKnown, 1)
}
func RelativeToStage(wellKnown WellKnownFilterStage, weight int) FilterStage {
	return FilterStage{
		RelativeTo: wellKnown,
		Weight:     weight,
	}
}

/*
	Generation plugins
*/
type ResourceGeneratorPlugin interface {
	Plugin
	GeneratedResources(params Params,
		inClusters []*envoy_config_cluster_v3.Cluster,
		inEndpoints []*envoy_config_endpoint_v3.ClusterLoadAssignment,
		inRouteConfigurations []*envoy_config_route_v3.RouteConfiguration,
		inListeners []*envoy_config_listener_v3.Listener,
	) ([]*envoy_config_cluster_v3.Cluster, []*envoy_config_endpoint_v3.ClusterLoadAssignment, []*envoy_config_route_v3.RouteConfiguration, []*envoy_config_listener_v3.Listener, error)
}

// A PluginRegistry is used to provide Plugins to relevant translators
// Historically, all plugins were passed around as an argument, and each translator
// would iterate over all plugins, and only apply the relevant ones.
// This interface enables translators to only know of the relevant plugins
type PluginRegistry interface {
	GetPlugins() []Plugin
	GetListenerPlugins() []ListenerPlugin
	GetTcpFilterChainPlugins() []TcpFilterChainPlugin
	GetHttpFilterPlugins() []HttpFilterPlugin
	GetHttpConnectionManagerPlugins() []HttpConnectionManagerPlugin
	GetVirtualHostPlugins() []VirtualHostPlugin
	GetResourceGeneratorPlugins() []ResourceGeneratorPlugin
	GetUpstreamPlugins() []UpstreamPlugin
	GetEndpointPlugins() []EndpointPlugin
	GetRoutePlugins() []RoutePlugin
	GetRouteActionPlugins() []RouteActionPlugin
	GetWeightedDestinationPlugins() []WeightedDestinationPlugin
}

// A PluginRegistryFactory generates a PluginRegistry
// It is executed each translation loop, ensuring we have up to date configuration of all plugins
type PluginRegistryFactory func(ctx context.Context) PluginRegistry
