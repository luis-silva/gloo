package clientip

import (
	envoycore "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

func NewPlugin() *Plugin {
	return &Plugin{}
}

var _ plugins.Plugin = new(Plugin)
var _ plugins.ListenerPlugin = new(Plugin)

type Plugin struct {
}

func (p *Plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *Plugin) ProcessListener(params plugins.Params, in *v1.Listener, out *envoy_config_listener_v3.Listener) error {
	if in.GetOptions() == nil {
		return nil
	}
	switch listenerType := in.GetListenerType().(type) {
	case *v1.Listener_HttpListener:
		if listenerType.HttpListener == nil {
			return nil
		}
		// add a matcher to the filter chain for each vh that has client IP matching
		for _, vh := range listenerType.HttpListener.GetVirtualHosts() {
			if len(vh.GetOptions().GetClientIpsToMatch()) > 0 {
				for i, fc := range out.GetFilterChains() {
					if fc.GetName() == vh.GetName() { // determine fc that matches vh
						ranges := []*envoycore.CidrRange{}
						for _, ip := range vh.GetOptions().GetClientIpsToMatch() {
							ranges = append(ranges, &envoycore.CidrRange{
								AddressPrefix: ip,
								PrefixLen: &wrappers.UInt32Value{
									Value: 32,
								},
							})
						}
						out.FilterChains[i].FilterChainMatch.SourcePrefixRanges = ranges
					}
				}
			}
		}
	}
	return nil
}





