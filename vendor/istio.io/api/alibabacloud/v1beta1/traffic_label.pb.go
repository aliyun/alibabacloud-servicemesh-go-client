// Copyright 2023 Alibaba Cloud Service Mesh
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: alibabacloud/v1beta1/traffic_label.proto

// In the Alibaba Cloud Service Mesh, traffic labeling refers to marking traffic for more detailed
// traffic control and management. By labeling traffic requests between application services
// with specific labels, they can be divided into different services or versions, and operations
// such as traffic control, circuit breaking, and rate limiting can be performed based on the labels.
//
// traffic labeling can be set for namespaces and workloads.
// `TrafficLabel` describes a way of traffic labeling. By defining specific traffic labeling custom resource,
// traffic labeling can be configured for specific workloads or by namespace.

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1alpha3 "istio.io/api/networking/v1alpha3"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `TrafficLabel` describes the configuration to apply the traffic label for specific workloads or by namespace.
//
// <!-- crd generation tags
// +cue-gen:TrafficLabel:groupName:istio.alibabacloud.com
// +cue-gen:TrafficLabel:version:v1beta1
// +cue-gen:TrafficLabel:storageVersion
// +cue-gen:TrafficLabel:annotations:helm.sh/resource-policy=keep
// +cue-gen:TrafficLabel:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:TrafficLabel:subresource:status
// +cue-gen:TrafficLabel:scope:Namespaced
// +cue-gen:TrafficLabel:resource:categories=alibabacloud-com,istio-alibabacloud-com,shortNames=tl
// +cue-gen:TrafficLabel:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=istio.alibabacloud.com/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type TrafficLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This refers to the scope of the workload to which the traffic labels will be applied.
	// The selector determines where the traffic label should be applied.
	// If not set, the selector will match all workloads in the current namespace.
	WorkloadSelector *v1alpha3.WorkloadSelector `protobuf:"bytes,1,opt,name=workload_selector,json=workloadSelector,proto3" json:"workload_selector,omitempty"`
	// The traffic rules for configuring the tags.
	Rules []*TrafficLabelRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	// The scope of the destination host which the traffic label rules should be applied.
	// "*" represents all services within the same namespace which the traffic label belongs to.
	// When short names are used (e.g. "reviews" instead of "reviews.default.svc.cluster.local"),
	// Istio will interpret the short name based on the namespace of the traffic label rule.
	// A rule in the "default" namespace containing a host "reviews" will be
	// interpreted as "reviews.default.svc.cluster.local". _To avoid
	// potential misconfigurations, it is recommended to always use fully
	// qualified domain names over short names._
	Hosts []string `protobuf:"bytes,5,rep,name=hosts,proto3" json:"hosts,omitempty"`
}

func (x *TrafficLabel) Reset() {
	*x = TrafficLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficLabel) ProtoMessage() {}

func (x *TrafficLabel) ProtoReflect() protoreflect.Message {
	mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficLabel.ProtoReflect.Descriptor instead.
func (*TrafficLabel) Descriptor() ([]byte, []int) {
	return file_alibabacloud_v1beta1_traffic_label_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficLabel) GetWorkloadSelector() *v1alpha3.WorkloadSelector {
	if x != nil {
		return x.WorkloadSelector
	}
	return nil
}

func (x *TrafficLabel) GetRules() []*TrafficLabelRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *TrafficLabel) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

// TrafficLabelRule describes the rule configuration of the traffic label.
type TrafficLabelRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The labels to be attached
	Labels []*LabelVar `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// attach to all protocols
	AttachTo []string `protobuf:"bytes,3,rep,name=attachTo,proto3" json:"attachTo,omitempty"`
	// protocols on which this rule is applied
	// '*' stands for all protocols
	// ” stands for none
	Protocols string `protobuf:"bytes,4,opt,name=protocols,proto3" json:"protocols,omitempty"`
	// Configuration of specific rules for HTTP request
	Http []*HTTPLabelRule `protobuf:"bytes,11,rep,name=http,proto3" json:"http,omitempty"`
}

func (x *TrafficLabelRule) Reset() {
	*x = TrafficLabelRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficLabelRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficLabelRule) ProtoMessage() {}

func (x *TrafficLabelRule) ProtoReflect() protoreflect.Message {
	mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficLabelRule.ProtoReflect.Descriptor instead.
func (*TrafficLabelRule) Descriptor() ([]byte, []int) {
	return file_alibabacloud_v1beta1_traffic_label_proto_rawDescGZIP(), []int{1}
}

func (x *TrafficLabelRule) GetLabels() []*LabelVar {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TrafficLabelRule) GetAttachTo() []string {
	if x != nil {
		return x.AttachTo
	}
	return nil
}

func (x *TrafficLabelRule) GetProtocols() string {
	if x != nil {
		return x.Protocols
	}
	return ""
}

func (x *TrafficLabelRule) GetHttp() []*HTTPLabelRule {
	if x != nil {
		return x.Http
	}
	return nil
}

// LabelVar describes the properties of label, e.g. name and value.
type LabelVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of label
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value of label. The parameter value of valueFrom follows a natural order of priority,
	// with the tag value first obtained from the first row and only obtained from the second row
	// when it is not available, and so on.
	ValueFrom []string `protobuf:"bytes,2,rep,name=valueFrom,proto3" json:"valueFrom,omitempty"`
	// Overwrite specifies whether the corresponding values of the original request header should be overwritten
	Overwrite bool `protobuf:"varint,5,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
}

func (x *LabelVar) Reset() {
	*x = LabelVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelVar) ProtoMessage() {}

func (x *LabelVar) ProtoReflect() protoreflect.Message {
	mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelVar.ProtoReflect.Descriptor instead.
func (*LabelVar) Descriptor() ([]byte, []int) {
	return file_alibabacloud_v1beta1_traffic_label_proto_rawDescGZIP(), []int{2}
}

func (x *LabelVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelVar) GetValueFrom() []string {
	if x != nil {
		return x.ValueFrom
	}
	return nil
}

func (x *LabelVar) GetOverwrite() bool {
	if x != nil {
		return x.Overwrite
	}
	return false
}

// HTTPLabelRule describes the rule configuration for http protocol
type HTTPLabelRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Match conditions to be satisfied for the rule to be
	// activated. All conditions inside a single match block have AND
	// semantics, while the list of match blocks have OR semantics. The rule
	// is matched if any one of the match blocks succeed.
	Match []*v1alpha3.HTTPMatchRequest `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	// labels for http request
	Labels []*LabelVar `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	// attach to http request
	AttachTo []string `protobuf:"bytes,6,rep,name=attachTo,proto3" json:"attachTo,omitempty"`
}

func (x *HTTPLabelRule) Reset() {
	*x = HTTPLabelRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPLabelRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPLabelRule) ProtoMessage() {}

func (x *HTTPLabelRule) ProtoReflect() protoreflect.Message {
	mi := &file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPLabelRule.ProtoReflect.Descriptor instead.
func (*HTTPLabelRule) Descriptor() ([]byte, []int) {
	return file_alibabacloud_v1beta1_traffic_label_proto_rawDescGZIP(), []int{3}
}

func (x *HTTPLabelRule) GetMatch() []*v1alpha3.HTTPMatchRequest {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *HTTPLabelRule) GetLabels() []*LabelVar {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *HTTPLabelRule) GetAttachTo() []string {
	if x != nil {
		return x.AttachTo
	}
	return nil
}

var File_alibabacloud_v1beta1_traffic_label_proto protoreflect.FileDescriptor

var file_alibabacloud_v1beta1_traffic_label_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x61, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x29, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x58, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x42, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x3c,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x56, 0x61, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61, 0x6c, 0x69,
	0x62, 0x61, 0x62, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x22, 0x5a, 0x0a, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x22, 0xac, 0x01, 0x0a, 0x0d, 0x48, 0x54, 0x54, 0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61,
	0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x72, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x6f,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x54, 0x6f,
	0x42, 0x23, 0x5a, 0x21, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alibabacloud_v1beta1_traffic_label_proto_rawDescOnce sync.Once
	file_alibabacloud_v1beta1_traffic_label_proto_rawDescData = file_alibabacloud_v1beta1_traffic_label_proto_rawDesc
)

func file_alibabacloud_v1beta1_traffic_label_proto_rawDescGZIP() []byte {
	file_alibabacloud_v1beta1_traffic_label_proto_rawDescOnce.Do(func() {
		file_alibabacloud_v1beta1_traffic_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_alibabacloud_v1beta1_traffic_label_proto_rawDescData)
	})
	return file_alibabacloud_v1beta1_traffic_label_proto_rawDescData
}

var file_alibabacloud_v1beta1_traffic_label_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_alibabacloud_v1beta1_traffic_label_proto_goTypes = []interface{}{
	(*TrafficLabel)(nil),              // 0: istio.alibabacloud.v1beta1.TrafficLabel
	(*TrafficLabelRule)(nil),          // 1: istio.alibabacloud.v1beta1.TrafficLabelRule
	(*LabelVar)(nil),                  // 2: istio.alibabacloud.v1beta1.LabelVar
	(*HTTPLabelRule)(nil),             // 3: istio.alibabacloud.v1beta1.HTTPLabelRule
	(*v1alpha3.WorkloadSelector)(nil), // 4: istio.networking.v1alpha3.WorkloadSelector
	(*v1alpha3.HTTPMatchRequest)(nil), // 5: istio.networking.v1alpha3.HTTPMatchRequest
}
var file_alibabacloud_v1beta1_traffic_label_proto_depIdxs = []int32{
	4, // 0: istio.alibabacloud.v1beta1.TrafficLabel.workload_selector:type_name -> istio.networking.v1alpha3.WorkloadSelector
	1, // 1: istio.alibabacloud.v1beta1.TrafficLabel.rules:type_name -> istio.alibabacloud.v1beta1.TrafficLabelRule
	2, // 2: istio.alibabacloud.v1beta1.TrafficLabelRule.labels:type_name -> istio.alibabacloud.v1beta1.LabelVar
	3, // 3: istio.alibabacloud.v1beta1.TrafficLabelRule.http:type_name -> istio.alibabacloud.v1beta1.HTTPLabelRule
	5, // 4: istio.alibabacloud.v1beta1.HTTPLabelRule.match:type_name -> istio.networking.v1alpha3.HTTPMatchRequest
	2, // 5: istio.alibabacloud.v1beta1.HTTPLabelRule.labels:type_name -> istio.alibabacloud.v1beta1.LabelVar
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_alibabacloud_v1beta1_traffic_label_proto_init() }
func file_alibabacloud_v1beta1_traffic_label_proto_init() {
	if File_alibabacloud_v1beta1_traffic_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficLabel); i {
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
		file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficLabelRule); i {
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
		file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelVar); i {
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
		file_alibabacloud_v1beta1_traffic_label_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPLabelRule); i {
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
			RawDescriptor: file_alibabacloud_v1beta1_traffic_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alibabacloud_v1beta1_traffic_label_proto_goTypes,
		DependencyIndexes: file_alibabacloud_v1beta1_traffic_label_proto_depIdxs,
		MessageInfos:      file_alibabacloud_v1beta1_traffic_label_proto_msgTypes,
	}.Build()
	File_alibabacloud_v1beta1_traffic_label_proto = out.File
	file_alibabacloud_v1beta1_traffic_label_proto_rawDesc = nil
	file_alibabacloud_v1beta1_traffic_label_proto_goTypes = nil
	file_alibabacloud_v1beta1_traffic_label_proto_depIdxs = nil
}
