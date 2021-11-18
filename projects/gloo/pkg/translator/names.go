package translator

import (
	"fmt"
	"sort"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

func routeConfigName(listener *v1.Listener) string {
	return listener.GetName() + "-routes"
}

func matchedRouteConfigName(listener *v1.Listener, matcher *v1.Matcher) string {
	return fmt.Sprintf("%s-%s", routeConfigName(listener), matcherID(matcher))
}

func matcherID(matcher *v1.Matcher) string {
	sort.Slice(matcher.SourcePrefixRanges, func(i, j int) bool {
		if matcher.SourcePrefixRanges[i].AddressPrefix != matcher.SourcePrefixRanges[j].AddressPrefix {
			return matcher.SourcePrefixRanges[i].AddressPrefix < matcher.SourcePrefixRanges[j].AddressPrefix
		}
		return matcher.SourcePrefixRanges[i].PrefixLen.Value < matcher.SourcePrefixRanges[j].PrefixLen.Value
	})

	sort.Strings(matcher.SslConfig.AlpnProtocols)
	sort.Strings(matcher.SslConfig.SniDomains)
	sort.Strings(matcher.SslConfig.VerifySubjectAltName)

	return matcher.String()
}