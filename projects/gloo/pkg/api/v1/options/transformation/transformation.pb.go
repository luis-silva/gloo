// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/transformation.proto

package transformation

import (
	_ "github.com/golang/protobuf/ptypes/wrappers"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation"
	xslt "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformers/xslt"
	matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response headers to match on.
	Matchers []*matchers.HeaderMatcher `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	// Response code detail to match on. To see the response code details for your usecase,
	// you can use the envoy access log %RESPONSE_CODE_DETAILS% formatter to log it.
	ResponseCodeDetails string `protobuf:"bytes,2,opt,name=response_code_details,json=responseCodeDetails,proto3" json:"response_code_details,omitempty"`
	// Transformation to apply on the response.
	ResponseTransformation *Transformation `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
}

func (x *ResponseMatch) Reset() {
	*x = ResponseMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMatch) ProtoMessage() {}

func (x *ResponseMatch) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMatch.ProtoReflect.Descriptor instead.
func (*ResponseMatch) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseMatch) GetMatchers() []*matchers.HeaderMatcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

func (x *ResponseMatch) GetResponseCodeDetails() string {
	if x != nil {
		return x.ResponseCodeDetails
	}
	return ""
}

func (x *ResponseMatch) GetResponseTransformation() *Transformation {
	if x != nil {
		return x.ResponseTransformation
	}
	return nil
}

type RequestMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matches on the request properties.
	Matcher *matchers.Matcher `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// Should we clear the route cache if a transformation was matched.
	ClearRouteCache bool `protobuf:"varint,2,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	// Transformation to apply on the request.
	RequestTransformation *Transformation `protobuf:"bytes,3,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// Transformation to apply on the response.
	ResponseTransformation *Transformation `protobuf:"bytes,4,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
}

func (x *RequestMatch) Reset() {
	*x = RequestMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMatch) ProtoMessage() {}

func (x *RequestMatch) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMatch.ProtoReflect.Descriptor instead.
func (*RequestMatch) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{1}
}

func (x *RequestMatch) GetMatcher() *matchers.Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *RequestMatch) GetClearRouteCache() bool {
	if x != nil {
		return x.ClearRouteCache
	}
	return false
}

func (x *RequestMatch) GetRequestTransformation() *Transformation {
	if x != nil {
		return x.RequestTransformation
	}
	return nil
}

func (x *RequestMatch) GetResponseTransformation() *Transformation {
	if x != nil {
		return x.ResponseTransformation
	}
	return nil
}

type Transformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Apply a transformation to requests.
	RequestTransformation *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// Clear the route cache if the request transformation was applied.
	ClearRouteCache bool `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	// Apply a transformation to responses.
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
}

func (x *Transformations) Reset() {
	*x = Transformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transformations) ProtoMessage() {}

func (x *Transformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transformations.ProtoReflect.Descriptor instead.
func (*Transformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{2}
}

func (x *Transformations) GetRequestTransformation() *Transformation {
	if x != nil {
		return x.RequestTransformation
	}
	return nil
}

func (x *Transformations) GetClearRouteCache() bool {
	if x != nil {
		return x.ClearRouteCache
	}
	return false
}

func (x *Transformations) GetResponseTransformation() *Transformation {
	if x != nil {
		return x.ResponseTransformation
	}
	return nil
}

type RequestResponseTransformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transformations to apply on the request. The first request that matches will apply.
	RequestTransforms []*RequestMatch `protobuf:"bytes,1,rep,name=request_transforms,json=requestTransforms,proto3" json:"request_transforms,omitempty"`
	// Transformations to apply on the response. This field is only consulted if there is no
	// response transformation in the matched `request_transforms`. i.e. Only one response transformation
	// will be executed. The first response transformation that matches will
	// apply.
	ResponseTransforms []*ResponseMatch `protobuf:"bytes,2,rep,name=response_transforms,json=responseTransforms,proto3" json:"response_transforms,omitempty"`
}

func (x *RequestResponseTransformations) Reset() {
	*x = RequestResponseTransformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponseTransformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponseTransformations) ProtoMessage() {}

func (x *RequestResponseTransformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponseTransformations.ProtoReflect.Descriptor instead.
func (*RequestResponseTransformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{3}
}

func (x *RequestResponseTransformations) GetRequestTransforms() []*RequestMatch {
	if x != nil {
		return x.RequestTransforms
	}
	return nil
}

func (x *RequestResponseTransformations) GetResponseTransforms() []*ResponseMatch {
	if x != nil {
		return x.ResponseTransforms
	}
	return nil
}

type TransformationStages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Early transformations happen before most other options (Like Auth and Rate Limit).
	Early *RequestResponseTransformations `protobuf:"bytes,1,opt,name=early,proto3" json:"early,omitempty"`
	// Regular transformations happen after Auth and Rate limit decisions has been made.
	Regular *RequestResponseTransformations `protobuf:"bytes,2,opt,name=regular,proto3" json:"regular,omitempty"`
	// Inherit transformation config from parent. This has no affect on VirtualHost level transformations.
	// If a RouteTable or Route wants to inherit transformations from it's parent RouteTable or VirtualHost,
	// this should be set to true, else transformations from parents will not be inherited.
	// Transformations are ordered so the child's transformation gets priority, so in the case where a child
	// and parent's transformation matchers are the same, only the child's transformation will run because
	// only one transformation will run per stage.
	// Defaults to false.
	InheritTransformation bool `protobuf:"varint,3,opt,name=inherit_transformation,json=inheritTransformation,proto3" json:"inherit_transformation,omitempty"`
}

func (x *TransformationStages) Reset() {
	*x = TransformationStages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformationStages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformationStages) ProtoMessage() {}

func (x *TransformationStages) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformationStages.ProtoReflect.Descriptor instead.
func (*TransformationStages) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{4}
}

func (x *TransformationStages) GetEarly() *RequestResponseTransformations {
	if x != nil {
		return x.Early
	}
	return nil
}

func (x *TransformationStages) GetRegular() *RequestResponseTransformations {
	if x != nil {
		return x.Regular
	}
	return nil
}

func (x *TransformationStages) GetInheritTransformation() bool {
	if x != nil {
		return x.InheritTransformation
	}
	return false
}

// User-facing API for transformation.
type Transformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of transformation to apply.
	//
	// Types that are assignable to TransformationType:
	//	*Transformation_TransformationTemplate
	//	*Transformation_HeaderBodyTransform
	//	*Transformation_XsltTransformation
	TransformationType isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
}

func (x *Transformation) Reset() {
	*x = Transformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transformation) ProtoMessage() {}

func (x *Transformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transformation.ProtoReflect.Descriptor instead.
func (*Transformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP(), []int{5}
}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (x *Transformation) GetTransformationTemplate() *transformation.TransformationTemplate {
	if x, ok := x.GetTransformationType().(*Transformation_TransformationTemplate); ok {
		return x.TransformationTemplate
	}
	return nil
}

func (x *Transformation) GetHeaderBodyTransform() *transformation.HeaderBodyTransform {
	if x, ok := x.GetTransformationType().(*Transformation_HeaderBodyTransform); ok {
		return x.HeaderBodyTransform
	}
	return nil
}

func (x *Transformation) GetXsltTransformation() *xslt.XsltTransformation {
	if x, ok := x.GetTransformationType().(*Transformation_XsltTransformation); ok {
		return x.XsltTransformation
	}
	return nil
}

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
}

type Transformation_TransformationTemplate struct {
	// Apply transformation templates.
	TransformationTemplate *transformation.TransformationTemplate `protobuf:"bytes,1,opt,name=transformation_template,json=transformationTemplate,proto3,oneof"`
}

type Transformation_HeaderBodyTransform struct {
	// This type of transformation will make all the headers available in the
	// response body. The resulting JSON body will consist of two attributes:
	// 'headers', containing the headers, and 'body', containing the original
	// body.
	HeaderBodyTransform *transformation.HeaderBodyTransform `protobuf:"bytes,2,opt,name=header_body_transform,json=headerBodyTransform,proto3,oneof"`
}

type Transformation_XsltTransformation struct {
	// (Enterprise Only): Xslt Transformation
	XsltTransformation *xslt.XsltTransformation `protobuf:"bytes,3,opt,name=xslt_transformation,json=xsltTransformation,proto3,oneof"`
}

func (*Transformation_TransformationTemplate) isTransformation_TransformationType() {}

func (*Transformation_HeaderBodyTransform) isTransformation_TransformationType() {}

func (*Transformation_XsltTransformation) isTransformation_TransformationType() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDesc = []byte{
	0x0a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x67, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x78, 0x73, 0x6c, 0x74, 0x2f, 0x78, 0x73, 0x6c, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x6c, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xd3, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x6a, 0x0a, 0x16,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x02, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6a, 0x0a, 0x16, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xe7, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x60, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x63, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x22, 0x87, 0x02, 0x0a, 0x14, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x05, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x12, 0x5d,
	0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x43, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x35, 0x0a,
	0x16, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69,
	0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x16, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x63, 0x0a, 0x15, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x48, 0x00, 0x52, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x67, 0x0a, 0x13, 0x78, 0x73, 0x6c,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x2e, 0x78, 0x73, 0x6c, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x58, 0x73, 0x6c, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12,
	0x78, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x51, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_goTypes = []interface{}{
	(*ResponseMatch)(nil),                         // 0: transformation.options.gloo.solo.io.ResponseMatch
	(*RequestMatch)(nil),                          // 1: transformation.options.gloo.solo.io.RequestMatch
	(*Transformations)(nil),                       // 2: transformation.options.gloo.solo.io.Transformations
	(*RequestResponseTransformations)(nil),        // 3: transformation.options.gloo.solo.io.RequestResponseTransformations
	(*TransformationStages)(nil),                  // 4: transformation.options.gloo.solo.io.TransformationStages
	(*Transformation)(nil),                        // 5: transformation.options.gloo.solo.io.Transformation
	(*matchers.HeaderMatcher)(nil),                // 6: matchers.core.gloo.solo.io.HeaderMatcher
	(*matchers.Matcher)(nil),                      // 7: matchers.core.gloo.solo.io.Matcher
	(*transformation.TransformationTemplate)(nil), // 8: envoy.api.v2.filter.http.TransformationTemplate
	(*transformation.HeaderBodyTransform)(nil),    // 9: envoy.api.v2.filter.http.HeaderBodyTransform
	(*xslt.XsltTransformation)(nil),               // 10: envoy.config.transformer.xslt.v2.XsltTransformation
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_depIdxs = []int32{
	6,  // 0: transformation.options.gloo.solo.io.ResponseMatch.matchers:type_name -> matchers.core.gloo.solo.io.HeaderMatcher
	5,  // 1: transformation.options.gloo.solo.io.ResponseMatch.response_transformation:type_name -> transformation.options.gloo.solo.io.Transformation
	7,  // 2: transformation.options.gloo.solo.io.RequestMatch.matcher:type_name -> matchers.core.gloo.solo.io.Matcher
	5,  // 3: transformation.options.gloo.solo.io.RequestMatch.request_transformation:type_name -> transformation.options.gloo.solo.io.Transformation
	5,  // 4: transformation.options.gloo.solo.io.RequestMatch.response_transformation:type_name -> transformation.options.gloo.solo.io.Transformation
	5,  // 5: transformation.options.gloo.solo.io.Transformations.request_transformation:type_name -> transformation.options.gloo.solo.io.Transformation
	5,  // 6: transformation.options.gloo.solo.io.Transformations.response_transformation:type_name -> transformation.options.gloo.solo.io.Transformation
	1,  // 7: transformation.options.gloo.solo.io.RequestResponseTransformations.request_transforms:type_name -> transformation.options.gloo.solo.io.RequestMatch
	0,  // 8: transformation.options.gloo.solo.io.RequestResponseTransformations.response_transforms:type_name -> transformation.options.gloo.solo.io.ResponseMatch
	3,  // 9: transformation.options.gloo.solo.io.TransformationStages.early:type_name -> transformation.options.gloo.solo.io.RequestResponseTransformations
	3,  // 10: transformation.options.gloo.solo.io.TransformationStages.regular:type_name -> transformation.options.gloo.solo.io.RequestResponseTransformations
	8,  // 11: transformation.options.gloo.solo.io.Transformation.transformation_template:type_name -> envoy.api.v2.filter.http.TransformationTemplate
	9,  // 12: transformation.options.gloo.solo.io.Transformation.header_body_transform:type_name -> envoy.api.v2.filter.http.HeaderBodyTransform
	10, // 13: transformation.options.gloo.solo.io.Transformation.xslt_transformation:type_name -> envoy.config.transformer.xslt.v2.XsltTransformation
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transformations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResponseTransformations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformationStages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Transformation_TransformationTemplate)(nil),
		(*Transformation_HeaderBodyTransform)(nil),
		(*Transformation_XsltTransformation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_transformation_transformation_proto_depIdxs = nil
}
