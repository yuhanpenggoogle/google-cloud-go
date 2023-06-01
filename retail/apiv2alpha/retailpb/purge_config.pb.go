// Copyright 2021 Google LLC
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
// source: google/cloud/retail/v2alpha/purge_config.proto

package retailpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Metadata related to the progress of the Purge operation.
// This will be returned by the google.longrunning.Operation.metadata field.
type PurgeMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PurgeMetadata) Reset() {
	*x = PurgeMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeMetadata) ProtoMessage() {}

func (x *PurgeMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeMetadata.ProtoReflect.Descriptor instead.
func (*PurgeMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{0}
}

// Metadata related to the progress of the PurgeProducts operation.
// This will be returned by the google.longrunning.Operation.metadata field.
type PurgeProductsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operation last update time. If the operation is done, this is also the
	// finish time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Count of entries that were deleted successfully.
	SuccessCount int64 `protobuf:"varint,3,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	// Count of entries that encountered errors while processing.
	FailureCount int64 `protobuf:"varint,4,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`
}

func (x *PurgeProductsMetadata) Reset() {
	*x = PurgeProductsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeProductsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeProductsMetadata) ProtoMessage() {}

func (x *PurgeProductsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeProductsMetadata.ProtoReflect.Descriptor instead.
func (*PurgeProductsMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{1}
}

func (x *PurgeProductsMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PurgeProductsMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *PurgeProductsMetadata) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *PurgeProductsMetadata) GetFailureCount() int64 {
	if x != nil {
		return x.FailureCount
	}
	return 0
}

// Request message for PurgeProducts method.
type PurgeProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the branch under which the products are
	// created. The format is
	// `projects/${projectId}/locations/global/catalogs/${catalogId}/branches/${branchId}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The filter string to specify the products to be deleted with a
	// length limit of 5,000 characters.
	//
	// Empty string filter is not allowed. "*" implies delete all items in a
	// branch.
	//
	// The eligible fields for filtering are:
	//
	// * `availability`: Double quoted
	// [Product.availability][google.cloud.retail.v2alpha.Product.availability]
	// string.
	// * `create_time` : in ISO 8601 "zulu" format.
	//
	// Supported syntax:
	//
	// * Comparators (">", "<", ">=", "<=", "=").
	//   Examples:
	//   * create_time <= "2015-02-13T17:05:46Z"
	//   * availability = "IN_STOCK"
	//
	// * Conjunctions ("AND")
	//   Examples:
	//   * create_time <= "2015-02-13T17:05:46Z" AND availability = "PREORDER"
	//
	// * Disjunctions ("OR")
	//   Examples:
	//   * create_time <= "2015-02-13T17:05:46Z" OR availability = "IN_STOCK"
	//
	// * Can support nested queries.
	//   Examples:
	//   * (create_time <= "2015-02-13T17:05:46Z" AND availability = "PREORDER")
	//   OR (create_time >= "2015-02-14T13:03:32Z" AND availability = "IN_STOCK")
	//
	// * Filter Limits:
	//   * Filter should not contain more than 6 conditions.
	//   * Max nesting depth should not exceed 2 levels.
	//
	// Examples queries:
	// * Delete back order products created before a timestamp.
	//   create_time <= "2015-02-13T17:05:46Z" OR availability = "BACKORDER"
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Actually perform the purge.
	// If `force` is set to false, the method will return the expected purge count
	// without deleting any products.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *PurgeProductsRequest) Reset() {
	*x = PurgeProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeProductsRequest) ProtoMessage() {}

func (x *PurgeProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeProductsRequest.ProtoReflect.Descriptor instead.
func (*PurgeProductsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{2}
}

func (x *PurgeProductsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PurgeProductsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PurgeProductsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Response of the PurgeProductsRequest. If the long running operation is
// successfully done, then this message is returned by the
// google.longrunning.Operations.response field.
type PurgeProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total count of products purged as a result of the operation.
	PurgeCount int64 `protobuf:"varint,1,opt,name=purge_count,json=purgeCount,proto3" json:"purge_count,omitempty"`
	// A sample of the product names that will be deleted.
	// Only populated if `force` is set to false. A max of 100 names will be
	// returned and the names are chosen at random.
	PurgeSample []string `protobuf:"bytes,2,rep,name=purge_sample,json=purgeSample,proto3" json:"purge_sample,omitempty"`
}

func (x *PurgeProductsResponse) Reset() {
	*x = PurgeProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeProductsResponse) ProtoMessage() {}

func (x *PurgeProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeProductsResponse.ProtoReflect.Descriptor instead.
func (*PurgeProductsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{3}
}

func (x *PurgeProductsResponse) GetPurgeCount() int64 {
	if x != nil {
		return x.PurgeCount
	}
	return 0
}

func (x *PurgeProductsResponse) GetPurgeSample() []string {
	if x != nil {
		return x.PurgeSample
	}
	return nil
}

// Request message for PurgeUserEvents method.
type PurgeUserEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the catalog under which the events are
	// created. The format is
	// `projects/${projectId}/locations/global/catalogs/${catalogId}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The filter string to specify the events to be deleted with a
	// length limit of 5,000 characters. Empty string filter is not allowed. The
	// eligible fields for filtering are:
	//
	// * `eventType`: Double quoted
	// [UserEvent.event_type][google.cloud.retail.v2alpha.UserEvent.event_type]
	// string.
	// * `eventTime`: in ISO 8601 "zulu" format.
	// * `visitorId`: Double quoted string. Specifying this will delete all
	//   events associated with a visitor.
	// * `userId`: Double quoted string. Specifying this will delete all events
	//   associated with a user.
	//
	// Examples:
	//
	// * Deleting all events in a time range:
	//   `eventTime > "2012-04-23T18:25:43.511Z"
	//   eventTime < "2012-04-23T18:30:43.511Z"`
	// * Deleting specific eventType in time range:
	//   `eventTime > "2012-04-23T18:25:43.511Z" eventType = "detail-page-view"`
	// * Deleting all events for a specific visitor:
	//   `visitorId = "visitor1024"`
	//
	// The filtering fields are assumed to have an implicit AND.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Actually perform the purge.
	// If `force` is set to false, the method will return the expected purge count
	// without deleting any user events.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *PurgeUserEventsRequest) Reset() {
	*x = PurgeUserEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeUserEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeUserEventsRequest) ProtoMessage() {}

func (x *PurgeUserEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeUserEventsRequest.ProtoReflect.Descriptor instead.
func (*PurgeUserEventsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{4}
}

func (x *PurgeUserEventsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PurgeUserEventsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PurgeUserEventsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Response of the PurgeUserEventsRequest. If the long running operation is
// successfully done, then this message is returned by the
// google.longrunning.Operations.response field.
type PurgeUserEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total count of events purged as a result of the operation.
	PurgedEventsCount int64 `protobuf:"varint,1,opt,name=purged_events_count,json=purgedEventsCount,proto3" json:"purged_events_count,omitempty"`
}

func (x *PurgeUserEventsResponse) Reset() {
	*x = PurgeUserEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeUserEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeUserEventsResponse) ProtoMessage() {}

func (x *PurgeUserEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeUserEventsResponse.ProtoReflect.Descriptor instead.
func (*PurgeUserEventsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP(), []int{5}
}

func (x *PurgeUserEventsResponse) GetPurgedEventsCount() int64 {
	if x != nil {
		return x.PurgedEventsCount
	}
	return 0
}

var File_google_cloud_retail_v2alpha_purge_config_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_purge_config_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x75,
	0x72, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xdb, 0x01, 0x0a, 0x15,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x14, 0x50, 0x75,
	0x72, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x24, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1e, 0x0a, 0x1c, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x22, 0x7f, 0x0a, 0x15, 0x50, 0x75, 0x72, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a,
	0x0c, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x22, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x70, 0x75, 0x72, 0x67, 0x65, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x50, 0x75, 0x72, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x25, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x22, 0x49, 0x0a, 0x17, 0x50, 0x75, 0x72, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x70, 0x75, 0x72, 0x67, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xd4, 0x01, 0x0a,
	0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x10, 0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06,
	0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_purge_config_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_purge_config_proto_rawDescData = file_google_cloud_retail_v2alpha_purge_config_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_purge_config_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_purge_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_purge_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_purge_config_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_purge_config_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_retail_v2alpha_purge_config_proto_goTypes = []interface{}{
	(*PurgeMetadata)(nil),           // 0: google.cloud.retail.v2alpha.PurgeMetadata
	(*PurgeProductsMetadata)(nil),   // 1: google.cloud.retail.v2alpha.PurgeProductsMetadata
	(*PurgeProductsRequest)(nil),    // 2: google.cloud.retail.v2alpha.PurgeProductsRequest
	(*PurgeProductsResponse)(nil),   // 3: google.cloud.retail.v2alpha.PurgeProductsResponse
	(*PurgeUserEventsRequest)(nil),  // 4: google.cloud.retail.v2alpha.PurgeUserEventsRequest
	(*PurgeUserEventsResponse)(nil), // 5: google.cloud.retail.v2alpha.PurgeUserEventsResponse
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_google_cloud_retail_v2alpha_purge_config_proto_depIdxs = []int32{
	6, // 0: google.cloud.retail.v2alpha.PurgeProductsMetadata.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: google.cloud.retail.v2alpha.PurgeProductsMetadata.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_purge_config_proto_init() }
func file_google_cloud_retail_v2alpha_purge_config_proto_init() {
	if File_google_cloud_retail_v2alpha_purge_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeMetadata); i {
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
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeProductsMetadata); i {
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
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeProductsRequest); i {
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
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeProductsResponse); i {
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
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeUserEventsRequest); i {
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
		file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeUserEventsResponse); i {
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
			RawDescriptor: file_google_cloud_retail_v2alpha_purge_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_purge_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_purge_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_purge_config_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_purge_config_proto = out.File
	file_google_cloud_retail_v2alpha_purge_config_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_purge_config_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_purge_config_proto_depIdxs = nil
}
