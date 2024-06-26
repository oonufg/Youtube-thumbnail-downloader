// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api.proto

package YtThumbGRPC

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
	YtThumbGRPC_DownloadThumbnails_FullMethodName      = "/YtThumbGRPC.YtThumbGRPC/DownloadThumbnails"
	YtThumbGRPC_DownloadThumbnailsAsync_FullMethodName = "/YtThumbGRPC.YtThumbGRPC/DownloadThumbnailsAsync"
)

// YtThumbGRPCClient is the client API for YtThumbGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YtThumbGRPCClient interface {
	DownloadThumbnails(ctx context.Context, in *DownloadThumbnailsRequest, opts ...grpc.CallOption) (*Empty, error)
	DownloadThumbnailsAsync(ctx context.Context, in *DownloadThumbnailsRequest, opts ...grpc.CallOption) (*Empty, error)
}

type ytThumbGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewYtThumbGRPCClient(cc grpc.ClientConnInterface) YtThumbGRPCClient {
	return &ytThumbGRPCClient{cc}
}

func (c *ytThumbGRPCClient) DownloadThumbnails(ctx context.Context, in *DownloadThumbnailsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YtThumbGRPC_DownloadThumbnails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ytThumbGRPCClient) DownloadThumbnailsAsync(ctx context.Context, in *DownloadThumbnailsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YtThumbGRPC_DownloadThumbnailsAsync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YtThumbGRPCServer is the server API for YtThumbGRPC service.
// All implementations must embed UnimplementedYtThumbGRPCServer
// for forward compatibility
type YtThumbGRPCServer interface {
	DownloadThumbnails(context.Context, *DownloadThumbnailsRequest) (*Empty, error)
	DownloadThumbnailsAsync(context.Context, *DownloadThumbnailsRequest) (*Empty, error)
	mustEmbedUnimplementedYtThumbGRPCServer()
}

// UnimplementedYtThumbGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedYtThumbGRPCServer struct {
}

func (UnimplementedYtThumbGRPCServer) DownloadThumbnails(context.Context, *DownloadThumbnailsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadThumbnails not implemented")
}
func (UnimplementedYtThumbGRPCServer) DownloadThumbnailsAsync(context.Context, *DownloadThumbnailsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadThumbnailsAsync not implemented")
}
func (UnimplementedYtThumbGRPCServer) mustEmbedUnimplementedYtThumbGRPCServer() {}

// UnsafeYtThumbGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YtThumbGRPCServer will
// result in compilation errors.
type UnsafeYtThumbGRPCServer interface {
	mustEmbedUnimplementedYtThumbGRPCServer()
}

func RegisterYtThumbGRPCServer(s grpc.ServiceRegistrar, srv YtThumbGRPCServer) {
	s.RegisterService(&YtThumbGRPC_ServiceDesc, srv)
}

func _YtThumbGRPC_DownloadThumbnails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadThumbnailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YtThumbGRPCServer).DownloadThumbnails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YtThumbGRPC_DownloadThumbnails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YtThumbGRPCServer).DownloadThumbnails(ctx, req.(*DownloadThumbnailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YtThumbGRPC_DownloadThumbnailsAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadThumbnailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YtThumbGRPCServer).DownloadThumbnailsAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YtThumbGRPC_DownloadThumbnailsAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YtThumbGRPCServer).DownloadThumbnailsAsync(ctx, req.(*DownloadThumbnailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YtThumbGRPC_ServiceDesc is the grpc.ServiceDesc for YtThumbGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YtThumbGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YtThumbGRPC.YtThumbGRPC",
	HandlerType: (*YtThumbGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadThumbnails",
			Handler:    _YtThumbGRPC_DownloadThumbnails_Handler,
		},
		{
			MethodName: "DownloadThumbnailsAsync",
			Handler:    _YtThumbGRPC_DownloadThumbnailsAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
