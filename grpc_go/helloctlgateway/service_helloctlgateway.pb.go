// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: protos/helloctlgateway/service_helloctlgateway.proto

package helloctlgateway

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_helloctlgateway_service_helloctlgateway_proto protoreflect.FileDescriptor

var file_protos_helloctlgateway_service_helloctlgateway_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74,
	0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x63, 0x74, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x79, 0x0a, 0x08, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x43, 0x74, 0x6c, 0x12, 0x6d, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74, 0x6c, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x63, 0x74,
	0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x61, 0x79, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x63, 0x74, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_protos_helloctlgateway_service_helloctlgateway_proto_goTypes = []interface{}{
	(*SayHelloReq)(nil),  // 0: helloctlgateway.SayHelloReq
	(*SayHelloResp)(nil), // 1: helloctlgateway.SayHelloResp
}
var file_protos_helloctlgateway_service_helloctlgateway_proto_depIdxs = []int32{
	0, // 0: helloctlgateway.HelloCtl.SayHello:input_type -> helloctlgateway.SayHelloReq
	1, // 1: helloctlgateway.HelloCtl.SayHello:output_type -> helloctlgateway.SayHelloResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_helloctlgateway_service_helloctlgateway_proto_init() }
func file_protos_helloctlgateway_service_helloctlgateway_proto_init() {
	if File_protos_helloctlgateway_service_helloctlgateway_proto != nil {
		return
	}
	file_protos_helloctlgateway_hello_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_helloctlgateway_service_helloctlgateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_helloctlgateway_service_helloctlgateway_proto_goTypes,
		DependencyIndexes: file_protos_helloctlgateway_service_helloctlgateway_proto_depIdxs,
	}.Build()
	File_protos_helloctlgateway_service_helloctlgateway_proto = out.File
	file_protos_helloctlgateway_service_helloctlgateway_proto_rawDesc = nil
	file_protos_helloctlgateway_service_helloctlgateway_proto_goTypes = nil
	file_protos_helloctlgateway_service_helloctlgateway_proto_depIdxs = nil
}
