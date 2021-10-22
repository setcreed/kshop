// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/user/service/v1/user.proto

package v1

import (
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

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  uint32 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *PageInfo) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageInfo) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Birthday string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     int32  `protobuf:"varint,6,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfoResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfoResponse) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UserInfoResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfoResponse) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type UserResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserResponseList) Reset() {
	*x = UserResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponseList) ProtoMessage() {}

func (x *UserResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponseList.ProtoReflect.Descriptor instead.
func (*UserResponseList) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserResponseList) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserResponseList) GetData() []*UserInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_user_service_v1_user_proto protoreflect.FileDescriptor

var file_api_user_service_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x42, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x5f, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x55, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x4d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_user_service_v1_user_proto_rawDescOnce sync.Once
	file_api_user_service_v1_user_proto_rawDescData = file_api_user_service_v1_user_proto_rawDesc
)

func file_api_user_service_v1_user_proto_rawDescGZIP() []byte {
	file_api_user_service_v1_user_proto_rawDescOnce.Do(func() {
		file_api_user_service_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_service_v1_user_proto_rawDescData)
	})
	return file_api_user_service_v1_user_proto_rawDescData
}

var file_api_user_service_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_user_service_v1_user_proto_goTypes = []interface{}{
	(*PageInfo)(nil),         // 0: user.service.v1.PageInfo
	(*UserInfoResponse)(nil), // 1: user.service.v1.UserInfoResponse
	(*UserResponseList)(nil), // 2: user.service.v1.UserResponseList
}
var file_api_user_service_v1_user_proto_depIdxs = []int32{
	1, // 0: user.service.v1.UserResponseList.data:type_name -> user.service.v1.UserInfoResponse
	0, // 1: user.service.v1.User.GetUserList:input_type -> user.service.v1.PageInfo
	2, // 2: user.service.v1.User.GetUserList:output_type -> user.service.v1.UserResponseList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_user_service_v1_user_proto_init() }
func file_api_user_service_v1_user_proto_init() {
	if File_api_user_service_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_service_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_api_user_service_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_api_user_service_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponseList); i {
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
			RawDescriptor: file_api_user_service_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_service_v1_user_proto_goTypes,
		DependencyIndexes: file_api_user_service_v1_user_proto_depIdxs,
		MessageInfos:      file_api_user_service_v1_user_proto_msgTypes,
	}.Build()
	File_api_user_service_v1_user_proto = out.File
	file_api_user_service_v1_user_proto_rawDesc = nil
	file_api_user_service_v1_user_proto_goTypes = nil
	file_api_user_service_v1_user_proto_depIdxs = nil
}
