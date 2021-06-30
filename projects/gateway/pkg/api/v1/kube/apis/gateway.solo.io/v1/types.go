// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"encoding/json"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"

	api "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type metaOnly struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=gateways
// +genclient
type Gateway struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec           api.Gateway         `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status         core.Status         `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	ReporterStatus core.ReporterStatus `json:"reporter_status,omitempty" protobuf:"bytes,4,opt,name=reporter_status"`
}

func (o *Gateway) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	delete(spec, "reporter_status")
	asMap := map[string]interface{}{
		"metadata":        o.ObjectMeta,
		"apiVersion":      o.TypeMeta.APIVersion,
		"kind":            o.TypeMeta.Kind,
		"status":          o.Status,
		"reporter_status": o.ReporterStatus,
		"spec":            spec,
	}
	return json.Marshal(asMap)
}

func (o *Gateway) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.Gateway
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = Gateway{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}
	if spec.ReporterStatus != nil {
		o.ReporterStatus = *spec.ReporterStatus
		o.Spec.ReporterStatus = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GatewayList is a collection of Gateways.
type GatewayList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Gateway `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=routeoptions
// +genclient
type RouteOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec           api.RouteOption     `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status         core.Status         `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	ReporterStatus core.ReporterStatus `json:"reporter_status,omitempty" protobuf:"bytes,4,opt,name=reporter_status"`
}

func (o *RouteOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	delete(spec, "reporter_status")
	asMap := map[string]interface{}{
		"metadata":        o.ObjectMeta,
		"apiVersion":      o.TypeMeta.APIVersion,
		"kind":            o.TypeMeta.Kind,
		"status":          o.Status,
		"reporter_status": o.ReporterStatus,
		"spec":            spec,
	}
	return json.Marshal(asMap)
}

func (o *RouteOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.RouteOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = RouteOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}
	if spec.ReporterStatus != nil {
		o.ReporterStatus = *spec.ReporterStatus
		o.Spec.ReporterStatus = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RouteOptionList is a collection of RouteOptions.
type RouteOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RouteOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=routetables
// +genclient
type RouteTable struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec           api.RouteTable      `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status         core.Status         `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	ReporterStatus core.ReporterStatus `json:"reporter_status,omitempty" protobuf:"bytes,4,opt,name=reporter_status"`
}

func (o *RouteTable) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	delete(spec, "reporter_status")
	asMap := map[string]interface{}{
		"metadata":        o.ObjectMeta,
		"apiVersion":      o.TypeMeta.APIVersion,
		"kind":            o.TypeMeta.Kind,
		"status":          o.Status,
		"reporter_status": o.ReporterStatus,
		"spec":            spec,
	}
	return json.Marshal(asMap)
}

func (o *RouteTable) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.RouteTable
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = RouteTable{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}
	if spec.ReporterStatus != nil {
		o.ReporterStatus = *spec.ReporterStatus
		o.Spec.ReporterStatus = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RouteTableList is a collection of RouteTables.
type RouteTableList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []RouteTable `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=virtualhostoptions
// +genclient
type VirtualHostOption struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec           api.VirtualHostOption `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status         core.Status           `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	ReporterStatus core.ReporterStatus   `json:"reporter_status,omitempty" protobuf:"bytes,4,opt,name=reporter_status"`
}

func (o *VirtualHostOption) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	delete(spec, "reporter_status")
	asMap := map[string]interface{}{
		"metadata":        o.ObjectMeta,
		"apiVersion":      o.TypeMeta.APIVersion,
		"kind":            o.TypeMeta.Kind,
		"status":          o.Status,
		"reporter_status": o.ReporterStatus,
		"spec":            spec,
	}
	return json.Marshal(asMap)
}

func (o *VirtualHostOption) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.VirtualHostOption
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = VirtualHostOption{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}
	if spec.ReporterStatus != nil {
		o.ReporterStatus = *spec.ReporterStatus
		o.Spec.ReporterStatus = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VirtualHostOptionList is a collection of VirtualHostOptions.
type VirtualHostOptionList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []VirtualHostOption `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=virtualservices
// +genclient
type VirtualService struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec           api.VirtualService  `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status         core.Status         `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	ReporterStatus core.ReporterStatus `json:"reporter_status,omitempty" protobuf:"bytes,4,opt,name=reporter_status"`
}

func (o *VirtualService) MarshalJSON() ([]byte, error) {
	spec, err := protoutils.MarshalMap(&o.Spec)
	if err != nil {
		return nil, err
	}
	delete(spec, "metadata")
	delete(spec, "status")
	delete(spec, "reporter_status")
	asMap := map[string]interface{}{
		"metadata":        o.ObjectMeta,
		"apiVersion":      o.TypeMeta.APIVersion,
		"kind":            o.TypeMeta.Kind,
		"status":          o.Status,
		"reporter_status": o.ReporterStatus,
		"spec":            spec,
	}
	return json.Marshal(asMap)
}

func (o *VirtualService) UnmarshalJSON(data []byte) error {
	var metaOnly metaOnly
	if err := json.Unmarshal(data, &metaOnly); err != nil {
		return err
	}
	var spec api.VirtualService
	if err := protoutils.UnmarshalResource(data, &spec); err != nil {
		return err
	}
	spec.Metadata = nil
	*o = VirtualService{
		ObjectMeta: metaOnly.ObjectMeta,
		TypeMeta:   metaOnly.TypeMeta,
		Spec:       spec,
	}
	if spec.Status != nil {
		o.Status = *spec.Status
		o.Spec.Status = nil
	}
	if spec.ReporterStatus != nil {
		o.ReporterStatus = *spec.ReporterStatus
		o.Spec.ReporterStatus = nil
	}

	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VirtualServiceList is a collection of VirtualServices.
type VirtualServiceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []VirtualService `json:"items" protobuf:"bytes,2,rep,name=items"`
}
