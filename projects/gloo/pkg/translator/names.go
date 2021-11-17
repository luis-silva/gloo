package translator

import (
	"fmt"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

func routeConfigName(listener *v1.Listener) string {
	return listener.GetName() + "-routes"
}

func matchedRouteConfigName(listener *v1.Listener, matcher *v1.Matcher) string {
	return fmt.Sprintf("%s-%s", routeConfigName(listener), matcher.String())
}
