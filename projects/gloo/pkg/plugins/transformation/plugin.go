package transformation

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/contextutils"
	"k8s.io/utils/lru"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_type_matcher_v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/solo-io/gloo/pkg/utils/regexutils"
	envoyroutev3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/route/v3"
	envoytransformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
)

var (
	_ plugins.Plugin                    = new(Plugin)
	_ plugins.VirtualHostPlugin         = new(Plugin)
	_ plugins.WeightedDestinationPlugin = new(Plugin)
	_ plugins.RoutePlugin               = new(Plugin)
	_ plugins.HttpFilterPlugin          = new(Plugin)
)

const (
	ExtensionName    = "transformation"
	FilterName       = "io.solo.transformation"
	EarlyStageNumber = 1
<<<<<<< HEAD
=======
	AwsStageNumber   = 2
>>>>>>> master
)

var (
	earlyPluginStage = plugins.AfterStage(plugins.FaultStage)
	pluginStage      = plugins.AfterStage(plugins.AuthZStage)

	UnknownTransformationType = func(transformation interface{}) error {
		return fmt.Errorf("unknown transformation type %T", transformation)
	}
)

type TranslateTransformationFn func(*transformation.Transformation) (*envoytransformation.Transformation, error)

// This Plugin is exported only because it is utilized by the enterprise implementation
// We would prefer if the plugin were not exported and instead the required translation
// methods were exported.
// Other plugins may
type Plugin struct {
	RequireEarlyTransformation bool
	filterNeeded               bool
	TranslateTransformation    TranslateTransformationFn
	settings                   *v1.Settings

<<<<<<< HEAD
=======
	// validationLruCache is a map of: (transformation hash) -> (error state)
	validationLruCache *lru.Cache
}

>>>>>>> master
func NewPlugin() *Plugin {
	return &Plugin{
		validationLruCache: lru.New(1024),
	}
}

func (p *Plugin) Name() string {
	return ExtensionName
}

// Init attempts to set the plugin back to a clean slate state.
func (p *Plugin) Init(params plugins.InitParams) error {
	p.RequireEarlyTransformation = false
	p.filterNeeded = !params.Settings.GetGloo().GetRemoveUnusedFilters().GetValue()
	p.settings = params.Settings
	p.TranslateTransformation = TranslateTransformation
	return nil
}

func mergeFunc(tx *envoytransformation.RouteTransformations) pluginutils.ModifyFunc {
	return func(existing *any.Any) (proto.Message, error) {
		if existing == nil {
			return tx, nil
		}
		var transforms envoytransformation.RouteTransformations
		err := existing.UnmarshalTo(&transforms)
		if err != nil {
			// this should never happen
			return nil, err
		}
		transforms.Transformations = append(transforms.GetTransformations(), tx.GetTransformations()...)
		return &transforms, nil
	}
}

func (p *Plugin) ProcessVirtualHost(
	params plugins.VirtualHostParams,
	in *v1.VirtualHost,
	out *envoy_config_route_v3.VirtualHost,
) error {
	envoyTransformation, err := p.convertTransformation(
		params.Ctx,
		in.GetOptions().GetTransformations(),
		in.GetOptions().GetStagedTransformations(),
	)
	if err != nil {
		return err
	}
	if envoyTransformation == nil {
		return nil
	}
	err = p.validateTransformation(params.Ctx, envoyTransformation)
	if err != nil {
		return err
	}

	p.filterNeeded = true

<<<<<<< HEAD
	return pluginutils.SetVhostPerFilterConfig(out, FilterName, envoyTransformation)
=======
	return pluginutils.ModifyVhostPerFilterConfig(out, FilterName, mergeFunc(envoyTransformation))
>>>>>>> master
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoy_config_route_v3.Route) error {
	envoyTransformation, err := p.convertTransformation(
		params.Ctx,
		in.GetOptions().GetTransformations(),
		in.GetOptions().GetStagedTransformations(),
	)
	if err != nil {
		return err
	}
	if envoyTransformation == nil {
		return nil
	}
	err = p.validateTransformation(params.Ctx, envoyTransformation)
	if err != nil {
		return err
	}

	p.filterNeeded = true
<<<<<<< HEAD
	return pluginutils.SetRoutePerFilterConfig(out, FilterName, envoyTransformation)
=======
	return pluginutils.ModifyRoutePerFilterConfig(out, FilterName, mergeFunc(envoyTransformation))
>>>>>>> master
}

func (p *Plugin) ProcessWeightedDestination(
	params plugins.RouteParams,
	in *v1.WeightedDestination,
	out *envoy_config_route_v3.WeightedCluster_ClusterWeight,
) error {
	envoyTransformation, err := p.convertTransformation(
		params.Ctx,
		in.GetOptions().GetTransformations(),
		in.GetOptions().GetStagedTransformations(),
	)
	if err != nil {
		return err
	}
	if envoyTransformation == nil {
		return nil
	}

	err = p.validateTransformation(params.Ctx, envoyTransformation)
	if err != nil {
		return err
	}
	p.filterNeeded = true
<<<<<<< HEAD
	return pluginutils.SetWeightedClusterPerFilterConfig(out, FilterName, envoyTransformation)
=======
	return pluginutils.ModifyWeightedClusterPerFilterConfig(out, FilterName, mergeFunc(envoyTransformation))
>>>>>>> master
}

// HttpFilters emits the desired set of filters. Either 0, 1 or
// if earlytransformation is needed then 2 staged filters
func (p *Plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	var filters []plugins.StagedHttpFilter

	if !p.filterNeeded {
		return filters, nil
	}

	if p.RequireEarlyTransformation {
		// only add early transformations if we have to, to allow rolling gloo updates;
		// i.e. an older envoy without stages connects to gloo, it shouldn't have 2 filters.
		earlyStageConfig := &envoytransformation.FilterTransformations{
			Stage: EarlyStageNumber,
		}
		earlyFilter, err := plugins.NewStagedFilterWithConfig(FilterName, earlyStageConfig, earlyPluginStage)
		if err != nil {
			return nil, err
		}
		filters = append(filters, earlyFilter)
	}
	filters = append(filters, plugins.NewStagedFilter(FilterName, pluginStage))

	return filters, nil
}

func (p *Plugin) convertTransformation(
	ctx context.Context,
	t *transformation.Transformations,
	stagedTransformations *transformation.TransformationStages,
) (*envoytransformation.RouteTransformations, error) {
	if t == nil && stagedTransformations == nil {
		return nil, nil
	}
	ret := &envoytransformation.RouteTransformations{}
	if t != nil && stagedTransformations.GetRegular() == nil {
		// keep deprecated config until we are sure we don't need it.
		// on newer envoys it will be ignored.
		requestTransform, err := p.TranslateTransformation(t.GetRequestTransformation())
		if err != nil {
			return nil, err
		}
		responseTransform, err := p.TranslateTransformation(t.GetResponseTransformation())
		if err != nil {
			return nil, err
		}
		ret.RequestTransformation = requestTransform
		ret.ClearRouteCache = t.GetClearRouteCache()
		ret.ResponseTransformation = responseTransform
		// new config:
		// we have to have it too, as if any new config is defined the deprecated config is ignored.
		ret.Transformations = append(ret.GetTransformations(),
			&envoytransformation.RouteTransformations_RouteTransformation{
				Match: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch_{
					RequestMatch: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch{
						Match:                  nil,
						RequestTransformation:  requestTransform,
						ClearRouteCache:        t.GetClearRouteCache(),
						ResponseTransformation: responseTransform,
					},
				},
			})
	}

	if early := stagedTransformations.GetEarly(); early != nil {
		p.RequireEarlyTransformation = true
		transformations, err := p.getTransformations(ctx, EarlyStageNumber, early)
		if err != nil {
			return nil, err
		}
		ret.Transformations = append(ret.GetTransformations(), transformations...)
	}
	if regular := stagedTransformations.GetRegular(); regular != nil {
		transformations, err := p.getTransformations(ctx, 0, regular)
		if err != nil {
			return nil, err
		}
		ret.Transformations = append(ret.GetTransformations(), transformations...)
	}
	return ret, nil
}

func (p *Plugin) translateOSSTransformations(
	glooTransform *transformation.Transformation,
) (*envoytransformation.Transformation, error) {
	transform, err := p.TranslateTransformation(glooTransform)
	if err != nil {
		return nil, eris.Wrap(err, "this transformation type is not supported in open source Gloo Edge")
	}
	return transform, nil
}

func TranslateTransformation(glooTransform *transformation.Transformation) (
	*envoytransformation.Transformation,
	error,
) {
	if glooTransform == nil {
		return nil, nil
	}
	out := &envoytransformation.Transformation{}

	switch typedTransformation := glooTransform.GetTransformationType().(type) {
	case *transformation.Transformation_HeaderBodyTransform:
		{
			out.TransformationType = &envoytransformation.Transformation_HeaderBodyTransform{
				HeaderBodyTransform: typedTransformation.HeaderBodyTransform,
			}
		}
	case *transformation.Transformation_TransformationTemplate:
		{
			out.TransformationType = &envoytransformation.Transformation_TransformationTemplate{
				TransformationTemplate: typedTransformation.TransformationTemplate,
			}
		}
	default:
		return nil, UnknownTransformationType(typedTransformation)
	}
	return out, nil
}

func (p *Plugin) validateTransformation(
	ctx context.Context,
	transformations *envoytransformation.RouteTransformations,
) error {

	transformHash, err := transformations.Hash(nil)
	if err != nil {
		contextutils.LoggerFrom(ctx).DPanicf("error hashing transformation, should never happen: %v", err)
		return err
	}

	// This transformation has already been validated, return the result
	if err, ok := p.validationLruCache.Get(transformHash); ok {
		// Error may be nil here since it's just the cached result
		return err.(error)
	}

	err = bootstrap.ValidateBootstrap(ctx, p.settings, FilterName, transformations)
	p.validationLruCache.Add(transformHash, err)
	if err != nil {
		return err
	}
	return nil
}

func (p *Plugin) getTransformations(
	ctx context.Context,
	stage uint32,
	transformations *transformation.RequestResponseTransformations,
) ([]*envoytransformation.RouteTransformations_RouteTransformation, error) {
	var outTransformations []*envoytransformation.RouteTransformations_RouteTransformation
	for _, transformation := range transformations.GetResponseTransforms() {
		responseTransform, err := p.TranslateTransformation(transformation.GetResponseTransformation())
		if err != nil {
			return nil, err
		}
		outTransformations = append(outTransformations, &envoytransformation.RouteTransformations_RouteTransformation{
			Stage: stage,
			Match: &envoytransformation.RouteTransformations_RouteTransformation_ResponseMatch_{
				ResponseMatch: &envoytransformation.RouteTransformations_RouteTransformation_ResponseMatch{
					Match:                  getResponseMatcher(ctx, transformation),
					ResponseTransformation: responseTransform,
				},
			},
		})
	}

	for _, transformation := range transformations.GetRequestTransforms() {
		requestTransform, err := p.TranslateTransformation(transformation.GetRequestTransformation())
		if err != nil {
			return nil, err
		}
		responseTransform, err := p.TranslateTransformation(transformation.GetResponseTransformation())
		if err != nil {
			return nil, err
		}
		outTransformations = append(outTransformations, &envoytransformation.RouteTransformations_RouteTransformation{
			Stage: stage,
			Match: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch_{
				RequestMatch: &envoytransformation.RouteTransformations_RouteTransformation_RequestMatch{
					Match:                  getRequestMatcher(ctx, transformation.GetMatcher()),
					RequestTransformation:  requestTransform,
					ClearRouteCache:        transformation.GetClearRouteCache(),
					ResponseTransformation: responseTransform,
				},
			},
		})
	}
	return outTransformations, nil
}

// Note: these are copied from the translator and adapted to v3 apis. Once the transformer
// is v3 ready, we can remove these.
func getResponseMatcher(ctx context.Context, m *transformation.ResponseMatch) *envoytransformation.ResponseMatcher {
	matcher := &envoytransformation.ResponseMatcher{
		Headers: envoyHeaderMatcher(ctx, m.GetMatchers()),
	}
	if m.GetResponseCodeDetails() != "" {
		matcher.ResponseCodeDetails = &v3.StringMatcher{
			MatchPattern: &v3.StringMatcher_Exact{Exact: m.GetResponseCodeDetails()},
		}
	}
	return matcher
}

func getRequestMatcher(ctx context.Context, matcher *matchers.Matcher) *envoyroutev3.RouteMatch {
	if matcher == nil {
		return nil
	}
	match := &envoyroutev3.RouteMatch{
		Headers:         envoyHeaderMatcher(ctx, matcher.GetHeaders()),
		QueryParameters: envoyQueryMatcher(ctx, matcher.GetQueryParameters()),
	}
	if len(matcher.GetMethods()) > 0 {
		match.Headers = append(match.GetHeaders(), &envoyroutev3.HeaderMatcher{
			Name: ":method",
			HeaderMatchSpecifier: &envoyroutev3.HeaderMatcher_SafeRegexMatch{
				SafeRegexMatch: convertRegex(regexutils.NewRegex(ctx, strings.Join(matcher.GetMethods(), "|"))),
			},
		})
	}
	// need to do this because Go's proto implementation makes oneofs private
	// which genius thought of that?
	setEnvoyPathMatcher(ctx, matcher, match)
	return match
}

func setEnvoyPathMatcher(ctx context.Context, in *matchers.Matcher, out *envoyroutev3.RouteMatch) {
	switch path := in.GetPathSpecifier().(type) {
	case *matchers.Matcher_Exact:
		out.PathSpecifier = &envoyroutev3.RouteMatch_Path{
			Path: path.Exact,
		}
	case *matchers.Matcher_Regex:
		out.PathSpecifier = &envoyroutev3.RouteMatch_SafeRegex{
			SafeRegex: convertRegex(regexutils.NewRegex(ctx, path.Regex)),
		}
	case *matchers.Matcher_Prefix:
		out.PathSpecifier = &envoyroutev3.RouteMatch_Prefix{
			Prefix: path.Prefix,
		}
	}
}

func envoyQueryMatcher(
	ctx context.Context,
	in []*matchers.QueryParameterMatcher,
) []*envoyroutev3.QueryParameterMatcher {
	var out []*envoyroutev3.QueryParameterMatcher
	for _, matcher := range in {
		envoyMatch := &envoyroutev3.QueryParameterMatcher{
			Name: matcher.GetName(),
		}

		if matcher.GetValue() == "" {
			envoyMatch.QueryParameterMatchSpecifier = &envoyroutev3.QueryParameterMatcher_PresentMatch{
				PresentMatch: true,
			}
		} else {
			if matcher.GetRegex() {
				envoyMatch.QueryParameterMatchSpecifier = &envoyroutev3.QueryParameterMatcher_StringMatch{
					StringMatch: &v3.StringMatcher{
						MatchPattern: &v3.StringMatcher_SafeRegex{
							SafeRegex: convertRegex(regexutils.NewRegex(ctx, matcher.GetValue())),
						},
					},
				}
			} else {
				envoyMatch.QueryParameterMatchSpecifier = &envoyroutev3.QueryParameterMatcher_StringMatch{
					StringMatch: &v3.StringMatcher{
						MatchPattern: &v3.StringMatcher_Exact{
							Exact: matcher.GetValue(),
						},
					},
				}
			}
		}
		out = append(out, envoyMatch)
	}
	return out
}

func envoyHeaderMatcher(ctx context.Context, in []*matchers.HeaderMatcher) []*envoyroutev3.HeaderMatcher {
	var out []*envoyroutev3.HeaderMatcher
	for _, matcher := range in {
		envoyMatch := &envoyroutev3.HeaderMatcher{
			Name: matcher.GetName(),
		}
		if matcher.GetValue() == "" {
			envoyMatch.HeaderMatchSpecifier = &envoyroutev3.HeaderMatcher_PresentMatch{
				PresentMatch: true,
			}
		} else {
			if matcher.GetRegex() {
				regex := regexutils.NewRegex(ctx, matcher.GetValue())
				envoyMatch.HeaderMatchSpecifier = &envoyroutev3.HeaderMatcher_SafeRegexMatch{
					SafeRegexMatch: convertRegex(regex),
				}
			} else {
				envoyMatch.HeaderMatchSpecifier = &envoyroutev3.HeaderMatcher_ExactMatch{
					ExactMatch: matcher.GetValue(),
				}
			}
		}

		if matcher.GetInvertMatch() {
			envoyMatch.InvertMatch = true
		}
		out = append(out, envoyMatch)
	}
	return out
}

func convertRegex(regex *envoy_type_matcher_v3.RegexMatcher) *v3.RegexMatcher {
	if regex == nil {
		return nil
	}
	return &v3.RegexMatcher{
		EngineType: &v3.RegexMatcher_GoogleRe2{GoogleRe2: &v3.RegexMatcher_GoogleRE2{MaxProgramSize: regex.GetGoogleRe2().GetMaxProgramSize()}},
		Regex:      regex.GetRegex(),
	}
}
