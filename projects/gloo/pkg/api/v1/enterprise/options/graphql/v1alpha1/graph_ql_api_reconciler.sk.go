// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionGraphQLApiFunc func(original, desired *GraphQLApi) (bool, error)

type GraphQLApiReconciler interface {
	Reconcile(namespace string, desiredResources GraphQLApiList, transition TransitionGraphQLApiFunc, opts clients.ListOpts) error
}

func graphQLApisToResources(list GraphQLApiList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, graphQLApi := range list {
		resourceList = append(resourceList, graphQLApi)
	}
	return resourceList
}

func NewGraphQLApiReconciler(client GraphQLApiClient, statusSetter resources.StatusSetter) GraphQLApiReconciler {
	return &graphQLApiReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type graphQLApiReconciler struct {
	base reconcile.Reconciler
}

func (r *graphQLApiReconciler) Reconcile(namespace string, desiredResources GraphQLApiList, transition TransitionGraphQLApiFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "graphQLApi_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*GraphQLApi), desired.(*GraphQLApi))
		}
	}
	return r.base.Reconcile(namespace, graphQLApisToResources(desiredResources), transitionResources, opts)
}
