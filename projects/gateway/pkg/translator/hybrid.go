package translator

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

type HybridTranslator struct {
	WarnOnRouteShortCircuiting bool
}

func (t *HybridTranslator) GenerateListeners(ctx context.Context, proxyName string, snap *v1.ApiSnapshot, filteredGateways []*v1.Gateway, reports reporter.ResourceReports) []*gloov1.Listener {
	var result []*gloov1.Listener
	for _, gateway := range filteredGateways {
		if gateway.GetHybridGateway() == nil {
			continue
		}

		listener := makeListener(gateway)
		hybridListener := &gloov1.HybridListener{}

		for _, matchedGateway := range gateway.GetHybridGateway().GetMatchedGateways() {
			matcher := &gloov1.Matcher{
				SslConfig:          matchedGateway.GetMatcher().GetSslConfig(),
				SourcePrefixRanges: matchedGateway.GetMatcher().GetSourcePrefixRanges(),
			}

			switch gt := matchedGateway.GetGatewayType().(type) {
			case *v1.MatchedGateway_HttpGateway:
				// logic mirrors HttpTranslator.GenerateListeners
				if len(snap.VirtualServices) == 0 {
					snapHash := hashutils.MustHash(snap)
					contextutils.LoggerFrom(ctx).Debugf("%v had no virtual services", snapHash)
					continue // TODO: do we want the entire listener to fail or just this matched listener?
				}

				virtualServices := getVirtualServicesForMatchedHttpGateway(matchedGateway, gateway, snap.VirtualServices, reports)
				applyGlobalVirtualServiceSettings(ctx, virtualServices)
				validateVirtualServiceDomains(gateway, virtualServices, reports)
				httpListener := t.desiredHttpListenerForHybrid(gateway, proxyName, virtualServices, snap, reports)

				hybridListener.MatchedListeners = append(hybridListener.GetMatchedListeners(), &gloov1.MatchedListener{
					Matcher: matcher,
					ListenerType: &gloov1.MatchedListener_HttpListener{
						HttpListener: httpListener,
					},
				})
			case *v1.MatchedGateway_TcpGateway:
				// logic mirrors TcpTranslator.GenerateListeners
				hybridListener.MatchedListeners = append(hybridListener.GetMatchedListeners(), &gloov1.MatchedListener{
					Matcher: matcher,
					ListenerType: &gloov1.MatchedListener_TcpListener{
						TcpListener: &gloov1.TcpListener{
							Options:  gt.TcpGateway.GetOptions(),
							TcpHosts: gt.TcpGateway.GetTcpHosts(),
						},
					},
				})
			}
		}

		listener.ListenerType = &gloov1.Listener_HybridListener{
			HybridListener: hybridListener,
		}

		if err := appendSource(listener, gateway); err != nil {
			// should never happen
			reports.AddError(gateway, err)
		}

		result = append(result, listener)
	}
	return result
}

// logic mirrors HttpTranslator.desiredListenerForHttp()
func (t *HybridTranslator) desiredHttpListenerForHybrid(gateway *v1.Gateway, proxyName string, virtualServicesForGateway v1.VirtualServiceList, snapshot *v1.ApiSnapshot, reports reporter.ResourceReports) *gloov1.HttpListener {
	var virtualHosts []*gloov1.VirtualHost

	for _, virtualService := range virtualServicesForGateway.Sort() {
		if virtualService.GetVirtualHost() == nil {
			virtualService.VirtualHost = &v1.VirtualHost{}
		}
		vh, err := t.virtualServiceToVirtualHost(virtualService, gateway, proxyName, snapshot, reports)
		if err != nil {
			reports.AddError(virtualService, err)
			continue
		}
		virtualHosts = append(virtualHosts, vh)
	}

	var httpPlugins *gloov1.HttpListenerOptions
	if httpGateway := gateway.GetHttpGateway(); httpGateway != nil {
		httpPlugins = httpGateway.GetOptions()
	}
	httpListener := &gloov1.HttpListener{
		VirtualHosts: virtualHosts,
		Options:      httpPlugins,
	}

	return httpListener
}

// logic mirrors HttpTranslator.virtualServiceToVirtualHost()
func (t *HybridTranslator) virtualServiceToVirtualHost(vs *v1.VirtualService, gateway *v1.Gateway, proxyName string, snapshot *v1.ApiSnapshot, reports reporter.ResourceReports) (*gloov1.VirtualHost, error) {
	converter := NewRouteConverter(NewRouteTableSelector(snapshot.RouteTables), NewRouteTableIndexer())
	t.mergeDelegatedVirtualHostOptions(vs, snapshot.VirtualHostOptions, reports)

	// note: in the future it may be necessary to create unique routes per matcher within a hybrid gateway
	// in order to apply settings, such as rate limiting, on a per-matched gateway basis
	routes, err := converter.ConvertVirtualService(vs, gateway, proxyName, snapshot, reports)
	if err != nil {
		// internal error, should never happen
		return nil, err
	}

	vh := &gloov1.VirtualHost{
		Name:    VirtualHostName(vs),
		Domains: vs.GetVirtualHost().GetDomains(),
		Routes:  routes,
		Options: vs.GetVirtualHost().GetOptions(),
	}

	validateRoutes(vs, vh, reports)

	if t.WarnOnRouteShortCircuiting {
		validateRouteShortCircuiting(vs, vh, reports)
	}

	if err := appendSource(vh, vs); err != nil {
		// should never happen
		return nil, err
	}

	return vh, nil
}

// finds delegated VirtualHostOption Objects and merges the options into the virtual service
func (t *HybridTranslator) mergeDelegatedVirtualHostOptions(vs *v1.VirtualService, options v1.VirtualHostOptionList, reports reporter.ResourceReports) {
	optionRefs := vs.GetVirtualHost().GetOptionsConfigRefs().GetDelegateOptions()
	for _, optionRef := range optionRefs {
		vhOption, err := options.Find(optionRef.GetNamespace(), optionRef.GetName())
		if err != nil {
			reports.AddError(vs, err)
			continue
		}
		if vs.GetVirtualHost().GetOptions() == nil {
			vs.GetVirtualHost().Options = vhOption.GetOptions()
			continue
		}
		vs.GetVirtualHost().Options, err = mergeVirtualHostOptions(vs.GetVirtualHost().GetOptions(), vhOption.GetOptions())
		if err != nil {
			reports.AddError(vs, err)
		}
	}
}

// logic mirrors getVirtualServicesForGateway()
func getVirtualServicesForMatchedHttpGateway(matchedGateway *v1.MatchedGateway, parentGateway *v1.Gateway, virtualServices v1.VirtualServiceList, reports reporter.ResourceReports) v1.VirtualServiceList {
	var virtualServicesForGateway v1.VirtualServiceList
	if matchedGateway.GetHttpGateway() == nil {
		return virtualServicesForGateway
	}
	for _, vs := range virtualServices {
		contains, err := HttpGatewayContainsVirtualService(matchedGateway.GetHttpGateway(), vs, matchedGateway.GetMatcher().GetSslConfig() != nil)
		if err != nil {
			reports.AddError(parentGateway, err)
			continue
		}
		if contains {
			virtualServicesForGateway = append(virtualServicesForGateway, vs)
		}
	}

	return virtualServicesForGateway
}

// logic mirrors GatewayContainsVirtualService()
func HttpGatewayContainsVirtualService(httpGateway *v1.HttpGateway, virtualService *v1.VirtualService, ssl bool) (bool, error) {
	if ssl != hasSsl(virtualService) {
		return false, nil
	}

	if httpGateway.GetVirtualServiceExpressions() != nil {
		return virtualServiceValidForSelectorExpressions(virtualService, httpGateway.GetVirtualServiceExpressions(),
			httpGateway.GetVirtualServiceNamespaces())
	}
	if httpGateway.GetVirtualServiceSelector() != nil {
		return virtualServiceMatchesLabels(virtualService, httpGateway.GetVirtualServiceSelector(),
			httpGateway.GetVirtualServiceNamespaces()), nil
	}
	// use individual refs to collect virtual services
	virtualServiceRefs := httpGateway.GetVirtualServices()

	if len(virtualServiceRefs) == 0 {
		return virtualServiceNamespaceValidForGateway(httpGateway.GetVirtualServiceNamespaces(), virtualService), nil
	}

	vsRef := virtualService.GetMetadata().Ref()

	for _, ref := range virtualServiceRefs {
		if ref.Equal(vsRef) {
			return true, nil
		}
	}

	return false, nil
}
