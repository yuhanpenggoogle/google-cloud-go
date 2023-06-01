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
// source: google/cloud/gaming/v1/realms.proto

package gamingpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for RealmsService.ListRealms.
type ListRealmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name, in the following form:
	// `projects/{project}/locations/{location}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return.  If unspecified, server
	// will pick an appropriate default. Server may return fewer items than
	// requested. A caller should only rely on response's
	// [next_page_token][google.cloud.gaming.v1.ListRealmsResponse.next_page_token] to
	// determine if there are more realms left to be queried.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous List request,
	// if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The filter to apply to list results.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specifies the ordering of results following syntax at
	// https://cloud.google.com/apis/design/design_patterns#sorting_order.
	OrderBy string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListRealmsRequest) Reset() {
	*x = ListRealmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRealmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRealmsRequest) ProtoMessage() {}

func (x *ListRealmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRealmsRequest.ProtoReflect.Descriptor instead.
func (*ListRealmsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{0}
}

func (x *ListRealmsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListRealmsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRealmsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListRealmsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListRealmsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response message for RealmsService.ListRealms.
type ListRealmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of realms.
	Realms []*Realm `protobuf:"bytes,1,rep,name=realms,proto3" json:"realms,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListRealmsResponse) Reset() {
	*x = ListRealmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRealmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRealmsResponse) ProtoMessage() {}

func (x *ListRealmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRealmsResponse.ProtoReflect.Descriptor instead.
func (*ListRealmsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{1}
}

func (x *ListRealmsResponse) GetRealms() []*Realm {
	if x != nil {
		return x.Realms
	}
	return nil
}

func (x *ListRealmsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListRealmsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// Request message for RealmsService.GetRealm.
type GetRealmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the realm to retrieve, in the following form:
	// `projects/{project}/locations/{location}/realms/{realm}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRealmRequest) Reset() {
	*x = GetRealmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRealmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealmRequest) ProtoMessage() {}

func (x *GetRealmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealmRequest.ProtoReflect.Descriptor instead.
func (*GetRealmRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{2}
}

func (x *GetRealmRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for RealmsService.CreateRealm.
type CreateRealmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name, in the following form:
	// `projects/{project}/locations/{location}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The ID of the realm resource to be created.
	RealmId string `protobuf:"bytes,2,opt,name=realm_id,json=realmId,proto3" json:"realm_id,omitempty"`
	// Required. The realm resource to be created.
	Realm *Realm `protobuf:"bytes,3,opt,name=realm,proto3" json:"realm,omitempty"`
}

func (x *CreateRealmRequest) Reset() {
	*x = CreateRealmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRealmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRealmRequest) ProtoMessage() {}

func (x *CreateRealmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRealmRequest.ProtoReflect.Descriptor instead.
func (*CreateRealmRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRealmRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateRealmRequest) GetRealmId() string {
	if x != nil {
		return x.RealmId
	}
	return ""
}

func (x *CreateRealmRequest) GetRealm() *Realm {
	if x != nil {
		return x.Realm
	}
	return nil
}

// Request message for RealmsService.DeleteRealm.
type DeleteRealmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the realm to delete, in the following form:
	// `projects/{project}/locations/{location}/realms/{realm}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteRealmRequest) Reset() {
	*x = DeleteRealmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRealmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRealmRequest) ProtoMessage() {}

func (x *DeleteRealmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRealmRequest.ProtoReflect.Descriptor instead.
func (*DeleteRealmRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRealmRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for RealmsService.UpdateRealm.
type UpdateRealmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The realm to be updated.
	// Only fields specified in update_mask are updated.
	Realm *Realm `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Required. The update mask applies to the resource. For the `FieldMask`
	// definition, see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRealmRequest) Reset() {
	*x = UpdateRealmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRealmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRealmRequest) ProtoMessage() {}

func (x *UpdateRealmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRealmRequest.ProtoReflect.Descriptor instead.
func (*UpdateRealmRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRealmRequest) GetRealm() *Realm {
	if x != nil {
		return x.Realm
	}
	return nil
}

func (x *UpdateRealmRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for RealmsService.PreviewRealmUpdate.
type PreviewRealmUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The realm to be updated.
	// Only fields specified in update_mask are updated.
	Realm *Realm `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Required. The update mask applies to the resource. For the `FieldMask`
	// definition, see
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Optional. The target timestamp to compute the preview.
	PreviewTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=preview_time,json=previewTime,proto3" json:"preview_time,omitempty"`
}

func (x *PreviewRealmUpdateRequest) Reset() {
	*x = PreviewRealmUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewRealmUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewRealmUpdateRequest) ProtoMessage() {}

func (x *PreviewRealmUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewRealmUpdateRequest.ProtoReflect.Descriptor instead.
func (*PreviewRealmUpdateRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{6}
}

func (x *PreviewRealmUpdateRequest) GetRealm() *Realm {
	if x != nil {
		return x.Realm
	}
	return nil
}

func (x *PreviewRealmUpdateRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *PreviewRealmUpdateRequest) GetPreviewTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PreviewTime
	}
	return nil
}

// Response message for RealmsService.PreviewRealmUpdate.
type PreviewRealmUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ETag of the realm.
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	// The target state.
	TargetState *TargetState `protobuf:"bytes,3,opt,name=target_state,json=targetState,proto3" json:"target_state,omitempty"`
}

func (x *PreviewRealmUpdateResponse) Reset() {
	*x = PreviewRealmUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviewRealmUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviewRealmUpdateResponse) ProtoMessage() {}

func (x *PreviewRealmUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviewRealmUpdateResponse.ProtoReflect.Descriptor instead.
func (*PreviewRealmUpdateResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{7}
}

func (x *PreviewRealmUpdateResponse) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *PreviewRealmUpdateResponse) GetTargetState() *TargetState {
	if x != nil {
		return x.TargetState
	}
	return nil
}

// A realm resource.
type Realm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the realm, in the following form:
	// `projects/{project}/locations/{location}/realms/{realm}`. For
	// example, `projects/my-project/locations/{location}/realms/my-realm`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The creation time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last-modified time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The labels associated with this realm. Each label is a key-value pair.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Time zone where all policies targeting this realm are evaluated. The value
	// of this field must be from the IANA time zone database:
	// https://www.iana.org/time-zones.
	TimeZone string `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// ETag of the resource.
	Etag string `protobuf:"bytes,7,opt,name=etag,proto3" json:"etag,omitempty"`
	// Human readable description of the realm.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Realm) Reset() {
	*x = Realm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Realm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Realm) ProtoMessage() {}

func (x *Realm) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_realms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Realm.ProtoReflect.Descriptor instead.
func (*Realm) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_realms_proto_rawDescGZIP(), []int{8}
}

func (x *Realm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Realm) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Realm) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Realm) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Realm) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Realm) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Realm) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_gaming_v1_realms_proto protoreflect.FileDescriptor

var file_google_cloud_gaming_v1_realms_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x12,
	0x21, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x95, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x61, 0x6c, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61, 0x6c,
	0x6d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x12, 0x21, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x22, 0x53, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x6c, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x22, 0xdb, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x61, 0x6c, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x40, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x42, 0x0a,
	0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x78, 0x0a, 0x1a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x61, 0x6c,
	0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x12, 0x46, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xd5, 0x03, 0x0a, 0x05,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x20, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x21, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x36, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x61,
	0x6c, 0x6d, 0x7d, 0x42, 0x52, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gaming_v1_realms_proto_rawDescOnce sync.Once
	file_google_cloud_gaming_v1_realms_proto_rawDescData = file_google_cloud_gaming_v1_realms_proto_rawDesc
)

func file_google_cloud_gaming_v1_realms_proto_rawDescGZIP() []byte {
	file_google_cloud_gaming_v1_realms_proto_rawDescOnce.Do(func() {
		file_google_cloud_gaming_v1_realms_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gaming_v1_realms_proto_rawDescData)
	})
	return file_google_cloud_gaming_v1_realms_proto_rawDescData
}

var file_google_cloud_gaming_v1_realms_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_google_cloud_gaming_v1_realms_proto_goTypes = []interface{}{
	(*ListRealmsRequest)(nil),          // 0: google.cloud.gaming.v1.ListRealmsRequest
	(*ListRealmsResponse)(nil),         // 1: google.cloud.gaming.v1.ListRealmsResponse
	(*GetRealmRequest)(nil),            // 2: google.cloud.gaming.v1.GetRealmRequest
	(*CreateRealmRequest)(nil),         // 3: google.cloud.gaming.v1.CreateRealmRequest
	(*DeleteRealmRequest)(nil),         // 4: google.cloud.gaming.v1.DeleteRealmRequest
	(*UpdateRealmRequest)(nil),         // 5: google.cloud.gaming.v1.UpdateRealmRequest
	(*PreviewRealmUpdateRequest)(nil),  // 6: google.cloud.gaming.v1.PreviewRealmUpdateRequest
	(*PreviewRealmUpdateResponse)(nil), // 7: google.cloud.gaming.v1.PreviewRealmUpdateResponse
	(*Realm)(nil),                      // 8: google.cloud.gaming.v1.Realm
	nil,                                // 9: google.cloud.gaming.v1.Realm.LabelsEntry
	(*fieldmaskpb.FieldMask)(nil),      // 10: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
	(*TargetState)(nil),                // 12: google.cloud.gaming.v1.TargetState
}
var file_google_cloud_gaming_v1_realms_proto_depIdxs = []int32{
	8,  // 0: google.cloud.gaming.v1.ListRealmsResponse.realms:type_name -> google.cloud.gaming.v1.Realm
	8,  // 1: google.cloud.gaming.v1.CreateRealmRequest.realm:type_name -> google.cloud.gaming.v1.Realm
	8,  // 2: google.cloud.gaming.v1.UpdateRealmRequest.realm:type_name -> google.cloud.gaming.v1.Realm
	10, // 3: google.cloud.gaming.v1.UpdateRealmRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 4: google.cloud.gaming.v1.PreviewRealmUpdateRequest.realm:type_name -> google.cloud.gaming.v1.Realm
	10, // 5: google.cloud.gaming.v1.PreviewRealmUpdateRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 6: google.cloud.gaming.v1.PreviewRealmUpdateRequest.preview_time:type_name -> google.protobuf.Timestamp
	12, // 7: google.cloud.gaming.v1.PreviewRealmUpdateResponse.target_state:type_name -> google.cloud.gaming.v1.TargetState
	11, // 8: google.cloud.gaming.v1.Realm.create_time:type_name -> google.protobuf.Timestamp
	11, // 9: google.cloud.gaming.v1.Realm.update_time:type_name -> google.protobuf.Timestamp
	9,  // 10: google.cloud.gaming.v1.Realm.labels:type_name -> google.cloud.gaming.v1.Realm.LabelsEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_cloud_gaming_v1_realms_proto_init() }
func file_google_cloud_gaming_v1_realms_proto_init() {
	if File_google_cloud_gaming_v1_realms_proto != nil {
		return
	}
	file_google_cloud_gaming_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_gaming_v1_realms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRealmsRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRealmsResponse); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRealmRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRealmRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRealmRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRealmRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewRealmUpdateRequest); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviewRealmUpdateResponse); i {
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
		file_google_cloud_gaming_v1_realms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Realm); i {
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
			RawDescriptor: file_google_cloud_gaming_v1_realms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gaming_v1_realms_proto_goTypes,
		DependencyIndexes: file_google_cloud_gaming_v1_realms_proto_depIdxs,
		MessageInfos:      file_google_cloud_gaming_v1_realms_proto_msgTypes,
	}.Build()
	File_google_cloud_gaming_v1_realms_proto = out.File
	file_google_cloud_gaming_v1_realms_proto_rawDesc = nil
	file_google_cloud_gaming_v1_realms_proto_goTypes = nil
	file_google_cloud_gaming_v1_realms_proto_depIdxs = nil
}
