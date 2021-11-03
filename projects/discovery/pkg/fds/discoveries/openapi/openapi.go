package openApi

import (
	"context"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1"
	"net/url"
	"time"

	"github.com/solo-io/gloo/projects/discovery/pkg/fds"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
)

var _ fds.Upgradable = new(OpenApiFunctionDiscoveryFactory)
var _ fds.FunctionDiscoveryFactory = new(OpenApiFunctionDiscoveryFactory)

type OpenApiFunctionDiscoveryFactory struct {
	DetectionTimeout time.Duration
	FunctionPollTime time.Duration
	OpenApiUrisToTry []string

	GraphqlClient v1alpha1.GraphQLSchemaClient
}

func (f *OpenApiFunctionDiscoveryFactory) FunctionDiscoveryFactoryName() string {
	return "OpenApiFunctionDiscoveryFactory"
}

func (f *OpenApiFunctionDiscoveryFactory) IsUpgrade() bool {
	return false
}

func (f *OpenApiFunctionDiscoveryFactory) NewFunctionDiscovery(u *v1.Upstream, clients fds.AdditionalClients) fds.UpstreamFunctionDiscovery {
	panic("not implemented")
}

type OpenApiFunctionDiscovery struct {
	detectionTimeout time.Duration
	functionPollTime time.Duration
	upstream         *v1.Upstream
	openApiUrisToTry []string

	graphqlClient v1alpha1.GraphQLSchemaClient
}

func (f *OpenApiFunctionDiscovery) IsFunctional() bool {
	panic("not implemented")
}

func (f *OpenApiFunctionDiscovery) DetectType(ctx context.Context, baseurl *url.URL) (*plugins.ServiceSpec, error) {
	panic("not implemented")
}

func (f *OpenApiFunctionDiscovery) detectUpstreamTypeOnce(ctx context.Context, baseUrl *url.URL) (*plugins.ServiceSpec, error) {
	panic("not implemented")
}

func (f *OpenApiFunctionDiscovery) DetectFunctions(ctx context.Context, url *url.URL, _ func() fds.Dependencies, updatecb func(fds.UpstreamMutator) error) error {
	panic("not implemented")
}
