// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"os"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewAuthConfig(namespace, name string) *AuthConfig {
	authconfig := &AuthConfig{}
	authconfig.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return authconfig
}

func (r *AuthConfig) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

func (r *AuthConfig) SetStatus(status *core.Status) {
	r.Status = status
}

func (r *AuthConfig) SetReporterStatus(status *core.ReporterStatus) {
	r.ReporterStatus = status
}

func (r *AuthConfig) AddToReporterStatus(status *core.Status) {
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace != "" {
		if r.ReporterStatus == nil {
			r.ReporterStatus = &core.ReporterStatus{}
		}
		if r.ReporterStatus.Statuses == nil {
			r.ReporterStatus.Statuses = make(map[string]*core.Status)
		}
		key := podNamespace + ":" + status.GetReportedBy()
		r.ReporterStatus.Statuses[key] = status
	}
}

func (r *AuthConfig) GetStatusForReporter(reportedBy string) *core.Status {
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace != "" {
		key := podNamespace + ":" + reportedBy
		if r.ReporterStatus == nil {
			return nil
		}
		if r.ReporterStatus.Statuses == nil {
			return nil
		}
		return r.ReporterStatus.Statuses[key]
	}
	return nil
}

func (r *AuthConfig) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *AuthConfig) GroupVersionKind() schema.GroupVersionKind {
	return AuthConfigGVK
}

type AuthConfigList []*AuthConfig

func (list AuthConfigList) Find(namespace, name string) (*AuthConfig, error) {
	for _, authConfig := range list {
		if authConfig.GetMetadata().Name == name && authConfig.GetMetadata().Namespace == namespace {
			return authConfig, nil
		}
	}
	return nil, errors.Errorf("list did not find authConfig %v.%v", namespace, name)
}

func (list AuthConfigList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, authConfig := range list {
		ress = append(ress, authConfig)
	}
	return ress
}

func (list AuthConfigList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, authConfig := range list {
		ress = append(ress, authConfig)
	}
	return ress
}

func (list AuthConfigList) Names() []string {
	var names []string
	for _, authConfig := range list {
		names = append(names, authConfig.GetMetadata().Name)
	}
	return names
}

func (list AuthConfigList) NamespacesDotNames() []string {
	var names []string
	for _, authConfig := range list {
		names = append(names, authConfig.GetMetadata().Namespace+"."+authConfig.GetMetadata().Name)
	}
	return names
}

func (list AuthConfigList) Sort() AuthConfigList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list AuthConfigList) Clone() AuthConfigList {
	var authConfigList AuthConfigList
	for _, authConfig := range list {
		authConfigList = append(authConfigList, resources.Clone(authConfig).(*AuthConfig))
	}
	return authConfigList
}

func (list AuthConfigList) Each(f func(element *AuthConfig)) {
	for _, authConfig := range list {
		f(authConfig)
	}
}

func (list AuthConfigList) EachResource(f func(element resources.Resource)) {
	for _, authConfig := range list {
		f(authConfig)
	}
}

func (list AuthConfigList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *AuthConfig) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for AuthConfig

func (o *AuthConfig) GetObjectKind() schema.ObjectKind {
	t := AuthConfigCrd.TypeMeta()
	return &t
}

func (o *AuthConfig) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*AuthConfig)
}

func (o *AuthConfig) DeepCopyInto(out *AuthConfig) {
	clone := resources.Clone(o).(*AuthConfig)
	*out = *clone
}

var (
	AuthConfigCrd = crd.NewCrd(
		"authconfigs",
		AuthConfigGVK.Group,
		AuthConfigGVK.Version,
		AuthConfigGVK.Kind,
		"ac",
		false,
		&AuthConfig{})
)

var (
	AuthConfigGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "enterprise.gloo.solo.io",
		Kind:    "AuthConfig",
	}
)
