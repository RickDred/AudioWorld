// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: messenger.proto

package messenger

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

const (
	Messenger_SendMessage_FullMethodName = "/messenger.Messenger/SendMessage"
)

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Messenger_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	SendMessage(context.Context, *Message) (*SendMessageResponse, error)
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) SendMessage(context.Context, *Message) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Messenger_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messenger.proto",
}