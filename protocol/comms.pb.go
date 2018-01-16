// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comms.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	comms.proto

It has these top-level messages:
	Volume
	Status
	DockerVolume
	VolumeAndHost
	Files
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Volume struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Volume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Status struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DockerVolume struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size    string `protobuf:"bytes,2,opt,name=size" json:"size,omitempty"`
	Updated string `protobuf:"bytes,3,opt,name=updated" json:"updated,omitempty"`
}

func (m *DockerVolume) Reset()                    { *m = DockerVolume{} }
func (m *DockerVolume) String() string            { return proto.CompactTextString(m) }
func (*DockerVolume) ProtoMessage()               {}
func (*DockerVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DockerVolume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerVolume) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *DockerVolume) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

type VolumeAndHost struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	HostFQDM string `protobuf:"bytes,2,opt,name=hostFQDM" json:"hostFQDM,omitempty"`
}

func (m *VolumeAndHost) Reset()                    { *m = VolumeAndHost{} }
func (m *VolumeAndHost) String() string            { return proto.CompactTextString(m) }
func (*VolumeAndHost) ProtoMessage()               {}
func (*VolumeAndHost) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VolumeAndHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VolumeAndHost) GetHostFQDM() string {
	if m != nil {
		return m.HostFQDM
	}
	return ""
}

type Files struct {
	Path string `protobuf:"bytes,64,opt,name=path" json:"path,omitempty"`
}

func (m *Files) Reset()                    { *m = Files{} }
func (m *Files) String() string            { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()               {}
func (*Files) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Files) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Volume)(nil), "protocol.Volume")
	proto.RegisterType((*Status)(nil), "protocol.Status")
	proto.RegisterType((*DockerVolume)(nil), "protocol.DockerVolume")
	proto.RegisterType((*VolumeAndHost)(nil), "protocol.VolumeAndHost")
	proto.RegisterType((*Files)(nil), "protocol.Files")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommsProto service

type CommsProtoClient interface {
	RMVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*Status, error)
	LSVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (CommsProto_LSVolumeClient, error)
	AddVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*Status, error)
	GetVolume(ctx context.Context, in *VolumeAndHost, opts ...grpc.CallOption) (*Status, error)
	VolumeFiles(ctx context.Context, in *Volume, opts ...grpc.CallOption) (CommsProto_VolumeFilesClient, error)
}

type commsProtoClient struct {
	cc *grpc.ClientConn
}

func NewCommsProtoClient(cc *grpc.ClientConn) CommsProtoClient {
	return &commsProtoClient{cc}
}

func (c *commsProtoClient) RMVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/protocol.CommsProto/RMVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsProtoClient) LSVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (CommsProto_LSVolumeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CommsProto_serviceDesc.Streams[0], c.cc, "/protocol.CommsProto/LSVolume", opts...)
	if err != nil {
		return nil, err
	}
	x := &commsProtoLSVolumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommsProto_LSVolumeClient interface {
	Recv() (*DockerVolume, error)
	grpc.ClientStream
}

type commsProtoLSVolumeClient struct {
	grpc.ClientStream
}

func (x *commsProtoLSVolumeClient) Recv() (*DockerVolume, error) {
	m := new(DockerVolume)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commsProtoClient) AddVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/protocol.CommsProto/AddVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsProtoClient) GetVolume(ctx context.Context, in *VolumeAndHost, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/protocol.CommsProto/GetVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commsProtoClient) VolumeFiles(ctx context.Context, in *Volume, opts ...grpc.CallOption) (CommsProto_VolumeFilesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CommsProto_serviceDesc.Streams[1], c.cc, "/protocol.CommsProto/VolumeFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &commsProtoVolumeFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommsProto_VolumeFilesClient interface {
	Recv() (*Files, error)
	grpc.ClientStream
}

type commsProtoVolumeFilesClient struct {
	grpc.ClientStream
}

func (x *commsProtoVolumeFilesClient) Recv() (*Files, error) {
	m := new(Files)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for CommsProto service

type CommsProtoServer interface {
	RMVolume(context.Context, *Volume) (*Status, error)
	LSVolume(*Volume, CommsProto_LSVolumeServer) error
	AddVolume(context.Context, *Volume) (*Status, error)
	GetVolume(context.Context, *VolumeAndHost) (*Status, error)
	VolumeFiles(*Volume, CommsProto_VolumeFilesServer) error
}

func RegisterCommsProtoServer(s *grpc.Server, srv CommsProtoServer) {
	s.RegisterService(&_CommsProto_serviceDesc, srv)
}

func _CommsProto_RMVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsProtoServer).RMVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.CommsProto/RMVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsProtoServer).RMVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommsProto_LSVolume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Volume)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommsProtoServer).LSVolume(m, &commsProtoLSVolumeServer{stream})
}

type CommsProto_LSVolumeServer interface {
	Send(*DockerVolume) error
	grpc.ServerStream
}

type commsProtoLSVolumeServer struct {
	grpc.ServerStream
}

func (x *commsProtoLSVolumeServer) Send(m *DockerVolume) error {
	return x.ServerStream.SendMsg(m)
}

func _CommsProto_AddVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsProtoServer).AddVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.CommsProto/AddVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsProtoServer).AddVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommsProto_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAndHost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommsProtoServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.CommsProto/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommsProtoServer).GetVolume(ctx, req.(*VolumeAndHost))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommsProto_VolumeFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Volume)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommsProtoServer).VolumeFiles(m, &commsProtoVolumeFilesServer{stream})
}

type CommsProto_VolumeFilesServer interface {
	Send(*Files) error
	grpc.ServerStream
}

type commsProtoVolumeFilesServer struct {
	grpc.ServerStream
}

func (x *commsProtoVolumeFilesServer) Send(m *Files) error {
	return x.ServerStream.SendMsg(m)
}

var _CommsProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.CommsProto",
	HandlerType: (*CommsProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RMVolume",
			Handler:    _CommsProto_RMVolume_Handler,
		},
		{
			MethodName: "AddVolume",
			Handler:    _CommsProto_AddVolume_Handler,
		},
		{
			MethodName: "GetVolume",
			Handler:    _CommsProto_GetVolume_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LSVolume",
			Handler:       _CommsProto_LSVolume_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "VolumeFiles",
			Handler:       _CommsProto_VolumeFiles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "comms.proto",
}

func init() { proto.RegisterFile("comms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x14, 0xc4, 0x29, 0x6a, 0x6d, 0x07, 0x8d, 0x66, 0x0f, 0xb8, 0x41, 0x0f, 0x64, 0x4f, 0x9e, 0x1a,
	0xfc, 0x73, 0xf0, 0xa6, 0x44, 0x82, 0x1e, 0x24, 0x41, 0x48, 0xbc, 0xd7, 0xee, 0x26, 0x10, 0xbb,
	0x6c, 0xc3, 0x6e, 0x2f, 0x7e, 0x15, 0xbf, 0xac, 0xd9, 0x3f, 0x08, 0x51, 0x38, 0x78, 0xea, 0xbc,
	0x5f, 0x77, 0xe6, 0xbd, 0x0c, 0x5a, 0x85, 0x92, 0x52, 0x67, 0xd5, 0x52, 0x19, 0x45, 0x12, 0xf7,
	0x29, 0x54, 0xc9, 0x2e, 0x10, 0xbf, 0xa9, 0xb2, 0x96, 0x82, 0x10, 0xec, 0x2f, 0x72, 0x29, 0x68,
	0xd4, 0x8d, 0x2e, 0xd3, 0x89, 0xd3, 0xac, 0x8b, 0x78, 0x6a, 0x72, 0x53, 0x6b, 0xd2, 0x46, 0xac,
	0x9d, 0x0a, 0xff, 0xc3, 0xc4, 0xc6, 0x38, 0x1a, 0xa8, 0xe2, 0x43, 0x2c, 0x77, 0xa7, 0x58, 0xa6,
	0xe7, 0x9f, 0x82, 0x36, 0x3d, 0xb3, 0x9a, 0x50, 0x1c, 0xd6, 0x15, 0xcf, 0x8d, 0xe0, 0x74, 0xcf,
	0xe1, 0xd5, 0xc8, 0xee, 0x71, 0xec, 0xb3, 0xfa, 0x0b, 0xfe, 0xac, 0xb4, 0xd9, 0x1a, 0xd9, 0x41,
	0x32, 0x53, 0xda, 0x0c, 0x5f, 0x07, 0xa3, 0x10, 0xfb, 0x33, 0xb3, 0x73, 0x1c, 0x0c, 0xe7, 0xa5,
	0xd0, 0xd6, 0x58, 0xe5, 0x66, 0x46, 0x1f, 0xbc, 0xd1, 0xea, 0xeb, 0xaf, 0x26, 0xf0, 0x68, 0x9b,
	0x18, 0xbb, 0x22, 0x7a, 0x48, 0x26, 0xa3, 0x70, 0xfa, 0x69, 0xb6, 0x6a, 0x25, 0xf3, 0xa4, 0xb3,
	0x41, 0x7c, 0x0d, 0xac, 0x41, 0xee, 0x90, 0xbc, 0x4c, 0x77, 0x3a, 0xda, 0x6b, 0xb2, 0x59, 0x0b,
	0x6b, 0xf4, 0x22, 0x72, 0x85, 0xb4, 0xcf, 0xf9, 0x3f, 0x97, 0xa5, 0x4f, 0xc2, 0x04, 0xcb, 0xd9,
	0x6f, 0x4b, 0x28, 0x68, 0xab, 0xf3, 0x16, 0x2d, 0xff, 0xc8, 0x57, 0xf1, 0x77, 0xdd, 0xc9, 0x9a,
	0xb8, 0x27, 0xf6, 0xc4, 0xf7, 0xd8, 0xb1, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xa9,
	0xdc, 0xb0, 0x2d, 0x02, 0x00, 0x00,
}
