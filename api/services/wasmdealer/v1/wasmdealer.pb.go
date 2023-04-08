// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/services/wasmdealer/v1/wasmdealer.proto

package wasmdealer

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the user-specified identifier.
	WasmId string `protobuf:"bytes,1,opt,name=wasm_id,json=wasmId,proto3" json:"wasm_id,omitempty"`
	// the directory path where of the wasm file
	ImagePath string `protobuf:"bytes,2,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	// Spec to be used when creating the container. This is runtime specific.
	Spec   *any1.Any `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	Stdin  string    `protobuf:"bytes,4,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout string    `protobuf:"bytes,5,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string    `protobuf:"bytes,6,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// runtime options
	RuntimeOptions *any1.Any `protobuf:"bytes,7,opt,name=runtime_options,json=runtimeOptions,proto3" json:"runtime_options,omitempty"`
	// task options
	TaskOptions *any1.Any `protobuf:"bytes,8,opt,name=task_options,json=taskOptions,proto3" json:"task_options,omitempty"`
	// optional and prioritized runtime, should be a constant value for wasmedge
	Runtime string `protobuf:"bytes,9,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetWasmId() string {
	if x != nil {
		return x.WasmId
	}
	return ""
}

func (x *CreateTaskRequest) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *CreateTaskRequest) GetSpec() *any1.Any {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CreateTaskRequest) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *CreateTaskRequest) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *CreateTaskRequest) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *CreateTaskRequest) GetRuntimeOptions() *any1.Any {
	if x != nil {
		return x.RuntimeOptions
	}
	return nil
}

func (x *CreateTaskRequest) GetTaskOptions() *any1.Any {
	if x != nil {
		return x.TaskOptions
	}
	return nil
}

func (x *CreateTaskRequest) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WasmId string `protobuf:"bytes,1,opt,name=wasm_id,json=wasmId,proto3" json:"wasm_id,omitempty"`
	Pid    uint32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetWasmId() string {
	if x != nil {
		return x.WasmId
	}
	return ""
}

func (x *CreateTaskResponse) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

var File_api_services_wasmdealer_v1_wasmdealer_proto protoreflect.FileDescriptor

var file_api_services_wasmdealer_v1_wasmdealer_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x77,
	0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x73,
	0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x73, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x3d, 0x0a, 0x0f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x73, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x32, 0x83, 0x01, 0x0a,
	0x0a, 0x57, 0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x61, 0x73, 0x6d, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescOnce sync.Once
	file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescData = file_api_services_wasmdealer_v1_wasmdealer_proto_rawDesc
)

func file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescGZIP() []byte {
	file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescOnce.Do(func() {
		file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescData)
	})
	return file_api_services_wasmdealer_v1_wasmdealer_proto_rawDescData
}

var file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_services_wasmdealer_v1_wasmdealer_proto_goTypes = []interface{}{
	(*CreateTaskRequest)(nil),  // 0: containerd.services.wasmdealer.v1.CreateTaskRequest
	(*CreateTaskResponse)(nil), // 1: containerd.services.wasmdealer.v1.CreateTaskResponse
	(*any1.Any)(nil),           // 2: google.protobuf.Any
}
var file_api_services_wasmdealer_v1_wasmdealer_proto_depIdxs = []int32{
	2, // 0: containerd.services.wasmdealer.v1.CreateTaskRequest.spec:type_name -> google.protobuf.Any
	2, // 1: containerd.services.wasmdealer.v1.CreateTaskRequest.runtime_options:type_name -> google.protobuf.Any
	2, // 2: containerd.services.wasmdealer.v1.CreateTaskRequest.task_options:type_name -> google.protobuf.Any
	0, // 3: containerd.services.wasmdealer.v1.Wasmdealer.Create:input_type -> containerd.services.wasmdealer.v1.CreateTaskRequest
	1, // 4: containerd.services.wasmdealer.v1.Wasmdealer.Create:output_type -> containerd.services.wasmdealer.v1.CreateTaskResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_services_wasmdealer_v1_wasmdealer_proto_init() }
func file_api_services_wasmdealer_v1_wasmdealer_proto_init() {
	if File_api_services_wasmdealer_v1_wasmdealer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
			RawDescriptor: file_api_services_wasmdealer_v1_wasmdealer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_services_wasmdealer_v1_wasmdealer_proto_goTypes,
		DependencyIndexes: file_api_services_wasmdealer_v1_wasmdealer_proto_depIdxs,
		MessageInfos:      file_api_services_wasmdealer_v1_wasmdealer_proto_msgTypes,
	}.Build()
	File_api_services_wasmdealer_v1_wasmdealer_proto = out.File
	file_api_services_wasmdealer_v1_wasmdealer_proto_rawDesc = nil
	file_api_services_wasmdealer_v1_wasmdealer_proto_goTypes = nil
	file_api_services_wasmdealer_v1_wasmdealer_proto_depIdxs = nil
}
