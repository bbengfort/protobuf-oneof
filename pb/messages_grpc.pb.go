// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: messages.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamerClient interface {
	Onef(ctx context.Context, opts ...grpc.CallOption) (Streamer_OnefClient, error)
	Anyf(ctx context.Context, opts ...grpc.CallOption) (Streamer_AnyfClient, error)
	Enum(ctx context.Context, opts ...grpc.CallOption) (Streamer_EnumClient, error)
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) Onef(ctx context.Context, opts ...grpc.CallOption) (Streamer_OnefClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[0], "/oneof.Streamer/Onef", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerOnefClient{stream}
	return x, nil
}

type Streamer_OnefClient interface {
	Send(*OnefMessage) error
	CloseAndRecv() (*Done, error)
	grpc.ClientStream
}

type streamerOnefClient struct {
	grpc.ClientStream
}

func (x *streamerOnefClient) Send(m *OnefMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamerOnefClient) CloseAndRecv() (*Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerClient) Anyf(ctx context.Context, opts ...grpc.CallOption) (Streamer_AnyfClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[1], "/oneof.Streamer/Anyf", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerAnyfClient{stream}
	return x, nil
}

type Streamer_AnyfClient interface {
	Send(*AnyfMessage) error
	CloseAndRecv() (*Done, error)
	grpc.ClientStream
}

type streamerAnyfClient struct {
	grpc.ClientStream
}

func (x *streamerAnyfClient) Send(m *AnyfMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamerAnyfClient) CloseAndRecv() (*Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamerClient) Enum(ctx context.Context, opts ...grpc.CallOption) (Streamer_EnumClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[2], "/oneof.Streamer/Enum", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerEnumClient{stream}
	return x, nil
}

type Streamer_EnumClient interface {
	Send(*EnumMessage) error
	CloseAndRecv() (*Done, error)
	grpc.ClientStream
}

type streamerEnumClient struct {
	grpc.ClientStream
}

func (x *streamerEnumClient) Send(m *EnumMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamerEnumClient) CloseAndRecv() (*Done, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Done)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamerServer is the server API for Streamer service.
// All implementations must embed UnimplementedStreamerServer
// for forward compatibility
type StreamerServer interface {
	Onef(Streamer_OnefServer) error
	Anyf(Streamer_AnyfServer) error
	Enum(Streamer_EnumServer) error
	mustEmbedUnimplementedStreamerServer()
}

// UnimplementedStreamerServer must be embedded to have forward compatible implementations.
type UnimplementedStreamerServer struct {
}

func (UnimplementedStreamerServer) Onef(Streamer_OnefServer) error {
	return status.Errorf(codes.Unimplemented, "method Onef not implemented")
}
func (UnimplementedStreamerServer) Anyf(Streamer_AnyfServer) error {
	return status.Errorf(codes.Unimplemented, "method Anyf not implemented")
}
func (UnimplementedStreamerServer) Enum(Streamer_EnumServer) error {
	return status.Errorf(codes.Unimplemented, "method Enum not implemented")
}
func (UnimplementedStreamerServer) mustEmbedUnimplementedStreamerServer() {}

// UnsafeStreamerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerServer will
// result in compilation errors.
type UnsafeStreamerServer interface {
	mustEmbedUnimplementedStreamerServer()
}

func RegisterStreamerServer(s grpc.ServiceRegistrar, srv StreamerServer) {
	s.RegisterService(&Streamer_ServiceDesc, srv)
}

func _Streamer_Onef_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).Onef(&streamerOnefServer{stream})
}

type Streamer_OnefServer interface {
	SendAndClose(*Done) error
	Recv() (*OnefMessage, error)
	grpc.ServerStream
}

type streamerOnefServer struct {
	grpc.ServerStream
}

func (x *streamerOnefServer) SendAndClose(m *Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamerOnefServer) Recv() (*OnefMessage, error) {
	m := new(OnefMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streamer_Anyf_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).Anyf(&streamerAnyfServer{stream})
}

type Streamer_AnyfServer interface {
	SendAndClose(*Done) error
	Recv() (*AnyfMessage, error)
	grpc.ServerStream
}

type streamerAnyfServer struct {
	grpc.ServerStream
}

func (x *streamerAnyfServer) SendAndClose(m *Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamerAnyfServer) Recv() (*AnyfMessage, error) {
	m := new(AnyfMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Streamer_Enum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).Enum(&streamerEnumServer{stream})
}

type Streamer_EnumServer interface {
	SendAndClose(*Done) error
	Recv() (*EnumMessage, error)
	grpc.ServerStream
}

type streamerEnumServer struct {
	grpc.ServerStream
}

func (x *streamerEnumServer) SendAndClose(m *Done) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamerEnumServer) Recv() (*EnumMessage, error) {
	m := new(EnumMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Streamer_ServiceDesc is the grpc.ServiceDesc for Streamer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streamer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oneof.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Onef",
			Handler:       _Streamer_Onef_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Anyf",
			Handler:       _Streamer_Anyf_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Enum",
			Handler:       _Streamer_Enum_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}