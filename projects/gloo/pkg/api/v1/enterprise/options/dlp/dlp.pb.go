// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/dlp/dlp.proto

package dlp

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	transformation_ee "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/transformation_ee"
	matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
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

type FilterConfig_EnableFor int32

const (
	// Only enable DLP masking of response bodies. Defaults to this value.
	FilterConfig_RESPONSE_BODY FilterConfig_EnableFor = 0
	// Only enable DLP masking of access logs.
	FilterConfig_ACCESS_LOGS FilterConfig_EnableFor = 1
	// Enable DLP masking for both responses and access logs.
	FilterConfig_ALL FilterConfig_EnableFor = 2
)

// Enum value maps for FilterConfig_EnableFor.
var (
	FilterConfig_EnableFor_name = map[int32]string{
		0: "RESPONSE_BODY",
		1: "ACCESS_LOGS",
		2: "ALL",
	}
	FilterConfig_EnableFor_value = map[string]int32{
		"RESPONSE_BODY": 0,
		"ACCESS_LOGS":   1,
		"ALL":           2,
	}
)

func (x FilterConfig_EnableFor) Enum() *FilterConfig_EnableFor {
	p := new(FilterConfig_EnableFor)
	*p = x
	return p
}

func (x FilterConfig_EnableFor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterConfig_EnableFor) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[0].Descriptor()
}

func (FilterConfig_EnableFor) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[0]
}

func (x FilterConfig_EnableFor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterConfig_EnableFor.Descriptor instead.
func (FilterConfig_EnableFor) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{0, 0}
}

type Config_EnableFor int32

const (
	// Only enable DLP masking of response bodies. Defaults to this value.
	Config_RESPONSE_BODY Config_EnableFor = 0
	// Only enable DLP masking of access logs.
	Config_ACCESS_LOGS Config_EnableFor = 1
	// Enable DLP masking for both responses and access logs.
	Config_ALL Config_EnableFor = 2
)

// Enum value maps for Config_EnableFor.
var (
	Config_EnableFor_name = map[int32]string{
		0: "RESPONSE_BODY",
		1: "ACCESS_LOGS",
		2: "ALL",
	}
	Config_EnableFor_value = map[string]int32{
		"RESPONSE_BODY": 0,
		"ACCESS_LOGS":   1,
		"ALL":           2,
	}
)

func (x Config_EnableFor) Enum() *Config_EnableFor {
	p := new(Config_EnableFor)
	*p = x
	return p
}

func (x Config_EnableFor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_EnableFor) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[1].Descriptor()
}

func (Config_EnableFor) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[1]
}

func (x Config_EnableFor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_EnableFor.Descriptor instead.
func (Config_EnableFor) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{2, 0}
}

//
//The following pre-made action types map to the following regex matchers:
//
//SSN:
//- '(?!\D)[0-9]{9}(?=\D|$)'
//- '(?!\D)[0-9]{3}\-[0-9]{2}\-[0-9]{4}(?=\D|$)'
//- '(?!\D)[0-9]{3}\ [0-9]{2}\ [0-9]{4}(?=\D|$)'
//
//MASTERCARD:
//- '(?!\D)5[1-5][0-9]{2}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//VISA:
//- '(?!\D)4[0-9]{3}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//AMEX:
//- '(?!\D)(34|37)[0-9]{2}(\ |\-|)[0-9]{6}(\ |\-|)[0-9]{5}(?=\D|$)'
//
//DISCOVER:
//- '(?!\D)6011(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//JCB:
//- '(?!\D)3[0-9]{3}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//- '(?!\D)(2131|1800)[0-9]{11}(?=\D|$)'
//
//DINERS_CLUB:
//- '(?!\D)30[0-5][0-9](\ |\-|)[0-9]{6}(\ |\-|)[0-9]{4}(?=\D|$)'
//- '(?!\D)(36|38)[0-9]{2}(\ |\-|)[0-9]{6}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//CREDIT_CARD_TRACKERS:
//- '[1-9][0-9]{2}\-[0-9]{2}\-[0-9]{4}\^\d'
//- '(?!\D)\%?[Bb]\d{13,19}\^[\-\/\.\w\s]{2,26}\^[0-9][0-9][01][0-9][0-9]{3}'
//- '(?!\D)\;\d{13,19}\=(\d{3}|)(\d{4}|\=)'
//
//ALL_CREDIT_CARDS:
//- (All credit card related regexes from above)
//
type Action_ActionType int32

const (
	Action_CUSTOM               Action_ActionType = 0
	Action_SSN                  Action_ActionType = 1
	Action_MASTERCARD           Action_ActionType = 2
	Action_VISA                 Action_ActionType = 3
	Action_AMEX                 Action_ActionType = 4
	Action_DISCOVER             Action_ActionType = 5
	Action_JCB                  Action_ActionType = 6
	Action_DINERS_CLUB          Action_ActionType = 7
	Action_CREDIT_CARD_TRACKERS Action_ActionType = 8
	Action_ALL_CREDIT_CARDS     Action_ActionType = 9
)

// Enum value maps for Action_ActionType.
var (
	Action_ActionType_name = map[int32]string{
		0: "CUSTOM",
		1: "SSN",
		2: "MASTERCARD",
		3: "VISA",
		4: "AMEX",
		5: "DISCOVER",
		6: "JCB",
		7: "DINERS_CLUB",
		8: "CREDIT_CARD_TRACKERS",
		9: "ALL_CREDIT_CARDS",
	}
	Action_ActionType_value = map[string]int32{
		"CUSTOM":               0,
		"SSN":                  1,
		"MASTERCARD":           2,
		"VISA":                 3,
		"AMEX":                 4,
		"DISCOVER":             5,
		"JCB":                  6,
		"DINERS_CLUB":          7,
		"CREDIT_CARD_TRACKERS": 8,
		"ALL_CREDIT_CARDS":     9,
	}
)

func (x Action_ActionType) Enum() *Action_ActionType {
	p := new(Action_ActionType)
	*p = x
	return p
}

func (x Action_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[2].Descriptor()
}

func (Action_ActionType) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes[2]
}

func (x Action_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action_ActionType.Descriptor instead.
func (Action_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{3, 0}
}

// Listener level config for dlp filter
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of transformation, matcher pairs.
	// The first rule which matches will be applied.
	DlpRules []*DlpRule `protobuf:"bytes,1,rep,name=dlp_rules,json=dlpRules,proto3" json:"dlp_rules,omitempty"`
	// Whether responses, access logs, or both should be masked by the applied actions.
	// If not defined, masking will only be enabled for responses bodies.
	EnabledFor FilterConfig_EnableFor `protobuf:"varint,2,opt,name=enabled_for,json=enabledFor,proto3,enum=dlp.options.gloo.solo.io.FilterConfig_EnableFor" json:"enabled_for,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetDlpRules() []*DlpRule {
	if x != nil {
		return x.DlpRules
	}
	return nil
}

func (x *FilterConfig) GetEnabledFor() FilterConfig_EnableFor {
	if x != nil {
		return x.EnabledFor
	}
	return FilterConfig_RESPONSE_BODY
}

// Rule which applies a given set of actions to a matching route.
// The route matching functions exactly the same as the envoy routes in the virtual host.
type DlpRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matcher by which to determine if the given transformation should be applied
	// if omitted, will it match all (i.e., default to / prefix matcher)
	Matcher *matchers.Matcher `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// List of data loss prevention actions to be applied.
	// These actions will be applied in order, one at a time.
	Actions []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *DlpRule) Reset() {
	*x = DlpRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DlpRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DlpRule) ProtoMessage() {}

func (x *DlpRule) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DlpRule.ProtoReflect.Descriptor instead.
func (*DlpRule) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{1}
}

func (x *DlpRule) GetMatcher() *matchers.Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *DlpRule) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

//
//Route/Vhost level config for dlp filter
//
//If a config is present on the route or vhost level it will completely overwrite the
//listener level config.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of data loss prevention actions to be applied.
	// These actions will be applied in order, one at a time.
	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	// Whether responses, access logs, or both should be masked by the applied actions.
	// If not defined, masking will only be enabled for responses bodies.
	EnabledFor Config_EnableFor `protobuf:"varint,2,opt,name=enabled_for,json=enabledFor,proto3,enum=dlp.options.gloo.solo.io.Config_EnableFor" json:"enabled_for,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Config) GetEnabledFor() Config_EnableFor {
	if x != nil {
		return x.EnabledFor
	}
	return Config_RESPONSE_BODY
}

//
//A single action meant to mask sensitive data.
//The action type represents a set of pre configured actions,
//as well as the ability to create custom actions.
//These actions can also be shadowed, a shadowed action will be recorded
//in the statistics, and debug logs, but not actually committed in the response body.
//
//To use a pre-made action simply set the action type to anything other than `CUSTOM`
//
//``` yaml
//actionType: VISA
//```
//
//To create a custom action set the custom action field. The default enum value
//is custom, so that can be left empty.
//
//``` yaml
//customAction:
//name: test
//regex:
//- "hello"
//- "world"
//maskChar: Y
//percent: 60
//```
//
//
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The action type to implement.
	ActionType Action_ActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=dlp.options.gloo.solo.io.Action_ActionType" json:"action_type,omitempty"`
	// The custom user action to be applied.
	// This field will only be used if the custom action type is specified above.
	CustomAction *CustomAction `protobuf:"bytes,2,opt,name=custom_action,json=customAction,proto3" json:"custom_action,omitempty"`
	// Shadow represents whether the action should be taken, or just recorded.
	Shadow bool `protobuf:"varint,3,opt,name=shadow,proto3" json:"shadow,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{3}
}

func (x *Action) GetActionType() Action_ActionType {
	if x != nil {
		return x.ActionType
	}
	return Action_CUSTOM
}

func (x *Action) GetCustomAction() *CustomAction {
	if x != nil {
		return x.CustomAction
	}
	return nil
}

func (x *Action) GetShadow() bool {
	if x != nil {
		return x.Shadow
	}
	return false
}

//
//A user defined custom action to carry out on the response body.
//
//The list of regex strings are applied in order. So for instance, if there is a response body with the content:
//`hello world`
//
//And there is a custom action
//``` yaml
//customAction:
//name: test
//regex:
//- "hello"
//- "world"
//maskChar: Y
//percent: 60
//```
//
//the result would be:
//`YYYlo YYYld`
//
//If the mask_char, and percent were left to default, the result would be:
//`XXXXo XXXXd`
//
type CustomAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the custom action.
	// This name is used for logging and debugging purposes.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of regex strings which will be applied in order.
	//
	// Deprecated: Do not use.
	Regex []string `protobuf:"bytes,2,rep,name=regex,proto3" json:"regex,omitempty"`
	// The masking character for the sensitive data.
	// default value: X
	MaskChar string `protobuf:"bytes,3,opt,name=mask_char,json=maskChar,proto3" json:"mask_char,omitempty"`
	// The percent of the string which will be masked by the mask_char
	// default value: 75%
	// rounds ratio (percent/100) by std::round http://www.cplusplus.com/reference/cmath/round/
	Percent *_type.Percent `protobuf:"bytes,4,opt,name=percent,proto3" json:"percent,omitempty"`
	// List of regexes to apply to the response body to match data which should be
	// masked. They will be applied iteratively in the order which they are
	// specified. If this field and `regex` are both provided, all the regexes will
	// be applied iteratively in the order provided, starting with the ones from `regex`
	RegexActions []*transformation_ee.RegexAction `protobuf:"bytes,5,rep,name=regex_actions,json=regexActions,proto3" json:"regex_actions,omitempty"`
}

func (x *CustomAction) Reset() {
	*x = CustomAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomAction) ProtoMessage() {}

func (x *CustomAction) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomAction.ProtoReflect.Descriptor instead.
func (*CustomAction) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP(), []int{4}
}

func (x *CustomAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Do not use.
func (x *CustomAction) GetRegex() []string {
	if x != nil {
		return x.Regex
	}
	return nil
}

func (x *CustomAction) GetMaskChar() string {
	if x != nil {
		return x.MaskChar
	}
	return ""
}

func (x *CustomAction) GetPercent() *_type.Percent {
	if x != nil {
		return x.Percent
	}
	return nil
}

func (x *CustomAction) GetRegexActions() []*transformation_ee.RegexAction {
	if x != nil {
		return x.RegexActions
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x64, 0x6c, 0x70, 0x2f, 0x64, 0x6c, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdb, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x3e, 0x0a, 0x09, 0x64, 0x6c, 0x70, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x44, 0x6c, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x64, 0x6c, 0x70, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x51, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x46, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x09, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f,
	0x72, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x42, 0x4f,
	0x44, 0x59, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c,
	0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x22, 0x84,
	0x01, 0x0a, 0x07, 0x44, 0x6c, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6c, 0x70,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x0b,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x52, 0x0a, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x22, 0x38, 0x0a, 0x09, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x42, 0x4f, 0x44, 0x59, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c,
	0x4c, 0x10, 0x02, 0x22, 0xdb, 0x02, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x6c, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x22, 0x9d, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x53, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x43,
	0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x49, 0x53, 0x41, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x41, 0x4d, 0x45, 0x58, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x43, 0x42, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x49, 0x4e, 0x45, 0x52, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x42, 0x10,
	0x07, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44,
	0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x52, 0x53, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x41,
	0x4c, 0x4c, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x53, 0x10,
	0x09, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x6b, 0x43, 0x68, 0x61, 0x72, 0x12, 0x35, 0x0a,
	0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x65, 0x78, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x51, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6c, 0x70,
	0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_goTypes = []interface{}{
	(FilterConfig_EnableFor)(0),           // 0: dlp.options.gloo.solo.io.FilterConfig.EnableFor
	(Config_EnableFor)(0),                 // 1: dlp.options.gloo.solo.io.Config.EnableFor
	(Action_ActionType)(0),                // 2: dlp.options.gloo.solo.io.Action.ActionType
	(*FilterConfig)(nil),                  // 3: dlp.options.gloo.solo.io.FilterConfig
	(*DlpRule)(nil),                       // 4: dlp.options.gloo.solo.io.DlpRule
	(*Config)(nil),                        // 5: dlp.options.gloo.solo.io.Config
	(*Action)(nil),                        // 6: dlp.options.gloo.solo.io.Action
	(*CustomAction)(nil),                  // 7: dlp.options.gloo.solo.io.CustomAction
	(*matchers.Matcher)(nil),              // 8: matchers.core.gloo.solo.io.Matcher
	(*_type.Percent)(nil),                 // 9: solo.io.envoy.type.Percent
	(*transformation_ee.RegexAction)(nil), // 10: envoy.config.filter.http.transformation_ee.v2.RegexAction
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_depIdxs = []int32{
	4,  // 0: dlp.options.gloo.solo.io.FilterConfig.dlp_rules:type_name -> dlp.options.gloo.solo.io.DlpRule
	0,  // 1: dlp.options.gloo.solo.io.FilterConfig.enabled_for:type_name -> dlp.options.gloo.solo.io.FilterConfig.EnableFor
	8,  // 2: dlp.options.gloo.solo.io.DlpRule.matcher:type_name -> matchers.core.gloo.solo.io.Matcher
	6,  // 3: dlp.options.gloo.solo.io.DlpRule.actions:type_name -> dlp.options.gloo.solo.io.Action
	6,  // 4: dlp.options.gloo.solo.io.Config.actions:type_name -> dlp.options.gloo.solo.io.Action
	1,  // 5: dlp.options.gloo.solo.io.Config.enabled_for:type_name -> dlp.options.gloo.solo.io.Config.EnableFor
	2,  // 6: dlp.options.gloo.solo.io.Action.action_type:type_name -> dlp.options.gloo.solo.io.Action.ActionType
	7,  // 7: dlp.options.gloo.solo.io.Action.custom_action:type_name -> dlp.options.gloo.solo.io.CustomAction
	9,  // 8: dlp.options.gloo.solo.io.CustomAction.percent:type_name -> solo.io.envoy.type.Percent
	10, // 9: dlp.options.gloo.solo.io.CustomAction.regex_actions:type_name -> envoy.config.filter.http.transformation_ee.v2.RegexAction
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DlpRule); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomAction); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_dlp_dlp_proto_depIdxs = nil
}
