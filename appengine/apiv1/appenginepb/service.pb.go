// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/appengine/v1/service.proto

package appenginepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Available sharding mechanisms.
type TrafficSplit_ShardBy int32

const (
	// Diversion method unspecified.
	TrafficSplit_UNSPECIFIED TrafficSplit_ShardBy = 0
	// Diversion based on a specially named cookie, "GOOGAPPUID." The cookie
	// must be set by the application itself or no diversion will occur.
	TrafficSplit_COOKIE TrafficSplit_ShardBy = 1
	// Diversion based on applying the modulus operation to a fingerprint
	// of the IP address.
	TrafficSplit_IP TrafficSplit_ShardBy = 2
	// Diversion based on weighted random assignment. An incoming request is
	// randomly routed to a version in the traffic split, with probability
	// proportional to the version's traffic share.
	TrafficSplit_RANDOM TrafficSplit_ShardBy = 3
)

// Enum value maps for TrafficSplit_ShardBy.
var (
	TrafficSplit_ShardBy_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "COOKIE",
		2: "IP",
		3: "RANDOM",
	}
	TrafficSplit_ShardBy_value = map[string]int32{
		"UNSPECIFIED": 0,
		"COOKIE":      1,
		"IP":          2,
		"RANDOM":      3,
	}
)

func (x TrafficSplit_ShardBy) Enum() *TrafficSplit_ShardBy {
	p := new(TrafficSplit_ShardBy)
	*p = x
	return p
}

func (x TrafficSplit_ShardBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrafficSplit_ShardBy) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_service_proto_enumTypes[0].Descriptor()
}

func (TrafficSplit_ShardBy) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_service_proto_enumTypes[0]
}

func (x TrafficSplit_ShardBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrafficSplit_ShardBy.Descriptor instead.
func (TrafficSplit_ShardBy) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

// A Service resource is a logical component of an application that can share
// state and communicate in a secure fashion with other services.
// For example, an application that handles customer requests might
// include separate services to handle tasks such as backend data
// analysis or API requests from mobile devices. Each service has a
// collection of versions that define a specific set of code used to
// implement the functionality of that service.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the Service resource in the API.
	// Example: `apps/myapp/services/default`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the service within the application.
	// Example: `default`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Mapping that defines fractional HTTP traffic diversion to
	// different versions within the service.
	Split *TrafficSplit `protobuf:"bytes,3,opt,name=split,proto3" json:"split,omitempty"`
	// A set of labels to apply to this service. Labels are key/value pairs that
	// describe the service and all resources that belong to it (e.g.,
	// versions). The labels can be used to search and group resources, and are
	// propagated to the usage and billing reports, enabling fine-grain analysis
	// of costs. An example of using labels is to tag resources belonging to
	// different environments (e.g., "env=prod", "env=qa").
	//
	// <p>Label keys and values can be no longer than 63 characters and can only
	// contain lowercase letters, numeric characters, underscores, dashes, and
	// international characters. Label keys must start with a lowercase letter
	// or an international character. Each service can have at most 32 labels.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Ingress settings for this service. Will apply to all versions.
	NetworkSettings *NetworkSettings `protobuf:"bytes,6,opt,name=network_settings,json=networkSettings,proto3" json:"network_settings,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetSplit() *TrafficSplit {
	if x != nil {
		return x.Split
	}
	return nil
}

func (x *Service) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Service) GetNetworkSettings() *NetworkSettings {
	if x != nil {
		return x.NetworkSettings
	}
	return nil
}

// Traffic routing configuration for versions within a single service. Traffic
// splits define how traffic directed to the service is assigned to versions.
type TrafficSplit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mechanism used to determine which version a request is sent to.
	// The traffic selection algorithm will
	// be stable for either type until allocations are changed.
	ShardBy TrafficSplit_ShardBy `protobuf:"varint,1,opt,name=shard_by,json=shardBy,proto3,enum=google.appengine.v1.TrafficSplit_ShardBy" json:"shard_by,omitempty"`
	// Mapping from version IDs within the service to fractional
	// (0.000, 1] allocations of traffic for that version. Each version can
	// be specified only once, but some versions in the service may not
	// have any traffic allocation. Services that have traffic allocated
	// cannot be deleted until either the service is deleted or
	// their traffic allocation is removed. Allocations must sum to 1.
	// Up to two decimal place precision is supported for IP-based splits and
	// up to three decimal places is supported for cookie-based splits.
	Allocations map[string]float64 `protobuf:"bytes,2,rep,name=allocations,proto3" json:"allocations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *TrafficSplit) Reset() {
	*x = TrafficSplit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficSplit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficSplit) ProtoMessage() {}

func (x *TrafficSplit) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficSplit.ProtoReflect.Descriptor instead.
func (*TrafficSplit) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *TrafficSplit) GetShardBy() TrafficSplit_ShardBy {
	if x != nil {
		return x.ShardBy
	}
	return TrafficSplit_UNSPECIFIED
}

func (x *TrafficSplit) GetAllocations() map[string]float64 {
	if x != nil {
		return x.Allocations
	}
	return nil
}

var File_google_appengine_v1_service_proto protoreflect.FileDescriptor

var file_google_appengine_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x52, 0x05, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x40, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x4f, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa6, 0x02, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x44, 0x0a, 0x08,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42, 0x79, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x42, 0x79, 0x12, 0x54, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x07, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x42, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x10, 0x01,
	0x12, 0x06, 0x0a, 0x02, 0x49, 0x50, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x41, 0x4e, 0x44,
	0x4f, 0x4d, 0x10, 0x03, 0x42, 0xbd, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0x3b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x19,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_service_proto_rawDescOnce sync.Once
	file_google_appengine_v1_service_proto_rawDescData = file_google_appengine_v1_service_proto_rawDesc
)

func file_google_appengine_v1_service_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_service_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_service_proto_rawDescData)
	})
	return file_google_appengine_v1_service_proto_rawDescData
}

var file_google_appengine_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_appengine_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_appengine_v1_service_proto_goTypes = []interface{}{
	(TrafficSplit_ShardBy)(0), // 0: google.appengine.v1.TrafficSplit.ShardBy
	(*Service)(nil),           // 1: google.appengine.v1.Service
	(*TrafficSplit)(nil),      // 2: google.appengine.v1.TrafficSplit
	nil,                       // 3: google.appengine.v1.Service.LabelsEntry
	nil,                       // 4: google.appengine.v1.TrafficSplit.AllocationsEntry
	(*NetworkSettings)(nil),   // 5: google.appengine.v1.NetworkSettings
}
var file_google_appengine_v1_service_proto_depIdxs = []int32{
	2, // 0: google.appengine.v1.Service.split:type_name -> google.appengine.v1.TrafficSplit
	3, // 1: google.appengine.v1.Service.labels:type_name -> google.appengine.v1.Service.LabelsEntry
	5, // 2: google.appengine.v1.Service.network_settings:type_name -> google.appengine.v1.NetworkSettings
	0, // 3: google.appengine.v1.TrafficSplit.shard_by:type_name -> google.appengine.v1.TrafficSplit.ShardBy
	4, // 4: google.appengine.v1.TrafficSplit.allocations:type_name -> google.appengine.v1.TrafficSplit.AllocationsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_service_proto_init() }
func file_google_appengine_v1_service_proto_init() {
	if File_google_appengine_v1_service_proto != nil {
		return
	}
	file_google_appengine_v1_network_settings_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_google_appengine_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficSplit); i {
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
			RawDescriptor: file_google_appengine_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_service_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_service_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1_service_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1_service_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_service_proto = out.File
	file_google_appengine_v1_service_proto_rawDesc = nil
	file_google_appengine_v1_service_proto_goTypes = nil
	file_google_appengine_v1_service_proto_depIdxs = nil
}
