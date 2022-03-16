package xdsinspection

import (
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/tg"

	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

func (xd *XdsDump) GetTargetGroupTargetsForUpstream(upstream *core.ResourceRef) []string {
	var out []string
	if xd == nil {
		out = append(out, "use -o wide for instance details")
		return out
	}
	clusterName := translator.UpstreamToClusterName(upstream)
	endpointCount := 0
	for _, clusterEndpoints := range xd.Endpoints {
		if clusterEndpoints.GetClusterName() == clusterName {
			for _, lEp := range clusterEndpoints.GetEndpoints() {
				for _, ep := range lEp.GetLbEndpoints() {
					if k, ok := ep.GetMetadata().GetFilterMetadata()[translator.SoloAnnotations]; ok {
						v, ok := k.GetFields()[tg.AvailabilityZoneAnnotationKey]
						if ok {
							endpointCount++
							out = append(out, v.GetStringValue())
						}
					}
				}
			}
		}
	}
	if endpointCount == 0 {
		out = append(out, "no endpoints")
	}
	return out
}
