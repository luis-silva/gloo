package translator

import (
	"crypto/md5"
	"fmt"
	"io"
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
	sort.Slice(matcher.GetSourcePrefixRanges(), func(i, j int) bool {
		if matcher.GetSourcePrefixRanges()[i].GetAddressPrefix() != matcher.GetSourcePrefixRanges()[j].GetAddressPrefix() {
			return matcher.GetSourcePrefixRanges()[i].GetAddressPrefix() < matcher.GetSourcePrefixRanges()[j].GetAddressPrefix()
		}
		return matcher.GetSourcePrefixRanges()[i].GetPrefixLen().GetValue() < matcher.GetSourcePrefixRanges()[j].GetPrefixLen().GetValue()
	})

	sort.Strings(matcher.GetSslConfig().GetAlpnProtocols())
	sort.Strings(matcher.GetSslConfig().GetSniDomains())
	sort.Strings(matcher.GetSslConfig().GetVerifySubjectAltName())

	h := md5.New()
	io.WriteString(h, matcher.String())

	return fmt.Sprintf("%x", h.Sum(nil))
}
