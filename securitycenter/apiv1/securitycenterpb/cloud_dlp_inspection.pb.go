// Copyright 2023 Google LLC
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
// source: google/cloud/securitycenter/v1/cloud_dlp_inspection.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Details about the Cloud Data Loss Prevention (Cloud DLP) [inspection
// job](https://cloud.google.com/dlp/docs/concepts-job-triggers) that produced
// the finding.
type CloudDlpInspection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the inspection job, for example,
	// `projects/123/locations/europe/dlpJobs/i-8383929`.
	InspectJob string `protobuf:"bytes,1,opt,name=inspect_job,json=inspectJob,proto3" json:"inspect_job,omitempty"`
	// The [type of
	// information](https://cloud.google.com/dlp/docs/infotypes-reference) found,
	// for example, `EMAIL_ADDRESS` or `STREET_ADDRESS`.
	InfoType string `protobuf:"bytes,2,opt,name=info_type,json=infoType,proto3" json:"info_type,omitempty"`
	// The number of times Cloud DLP found this infoType within this job
	// and resource.
	InfoTypeCount int64 `protobuf:"varint,3,opt,name=info_type_count,json=infoTypeCount,proto3" json:"info_type_count,omitempty"`
	// Whether Cloud DLP scanned the complete resource or a sampled subset.
	FullScan bool `protobuf:"varint,4,opt,name=full_scan,json=fullScan,proto3" json:"full_scan,omitempty"`
}

func (x *CloudDlpInspection) Reset() {
	*x = CloudDlpInspection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudDlpInspection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudDlpInspection) ProtoMessage() {}

func (x *CloudDlpInspection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudDlpInspection.ProtoReflect.Descriptor instead.
func (*CloudDlpInspection) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescGZIP(), []int{0}
}

func (x *CloudDlpInspection) GetInspectJob() string {
	if x != nil {
		return x.InspectJob
	}
	return ""
}

func (x *CloudDlpInspection) GetInfoType() string {
	if x != nil {
		return x.InfoType
	}
	return ""
}

func (x *CloudDlpInspection) GetInfoTypeCount() int64 {
	if x != nil {
		return x.InfoTypeCount
	}
	return 0
}

func (x *CloudDlpInspection) GetFullScan() bool {
	if x != nil {
		return x.FullScan
	}
	return false
}

var File_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x64, 0x6c, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x44, 0x6c, 0x70, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1e, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x64, 0x6c, 0x70, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6c, 0x70, 0x4a,
	0x6f, 0x62, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x66, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x53, 0x63, 0x61, 0x6e,
	0x42, 0xf0, 0x02, 0xea, 0x41, 0x7c, 0x0a, 0x19, 0x64, 0x6c, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6c, 0x70, 0x4a, 0x6f,
	0x62, 0x12, 0x24, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x6c, 0x70, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x64,
	0x6c, 0x70, 0x5f, 0x6a, 0x6f, 0x62, 0x7d, 0x12, 0x39, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x64, 0x6c, 0x70, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x64, 0x6c, 0x70, 0x5f, 0x6a, 0x6f,
	0x62, 0x7d, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x44, 0x6c, 0x70,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescData = file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_goTypes = []interface{}{
	(*CloudDlpInspection)(nil), // 0: google.cloud.securitycenter.v1.CloudDlpInspection
}
var file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_init() }
func file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_init() {
	if File_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudDlpInspection); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto = out.File
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_cloud_dlp_inspection_proto_depIdxs = nil
}
