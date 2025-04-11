// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: feed.proto

package feedRpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FeedService_PullLatest_FullMethodName = "/feedRpc.feedService/pullLatest"
	FeedService_Pull_FullMethodName       = "/feedRpc.feedService/pull"
	FeedService_PullOutBox_FullMethodName = "/feedRpc.feedService/pullOutBox"
)

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	PullLatest(ctx context.Context, in *PullLatestReq, opts ...grpc.CallOption) (*PullResp, error)
	Pull(ctx context.Context, in *PullReq, opts ...grpc.CallOption) (*PullResp, error)
	PullOutBox(ctx context.Context, in *PullOutBoxReq, opts ...grpc.CallOption) (*PullOutBoxResp, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) PullLatest(ctx context.Context, in *PullLatestReq, opts ...grpc.CallOption) (*PullResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PullResp)
	err := c.cc.Invoke(ctx, FeedService_PullLatest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) Pull(ctx context.Context, in *PullReq, opts ...grpc.CallOption) (*PullResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PullResp)
	err := c.cc.Invoke(ctx, FeedService_Pull_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) PullOutBox(ctx context.Context, in *PullOutBoxReq, opts ...grpc.CallOption) (*PullOutBoxResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PullOutBoxResp)
	err := c.cc.Invoke(ctx, FeedService_PullOutBox_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations must embed UnimplementedFeedServiceServer
// for forward compatibility.
type FeedServiceServer interface {
	PullLatest(context.Context, *PullLatestReq) (*PullResp, error)
	Pull(context.Context, *PullReq) (*PullResp, error)
	PullOutBox(context.Context, *PullOutBoxReq) (*PullOutBoxResp, error)
	mustEmbedUnimplementedFeedServiceServer()
}

// UnimplementedFeedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFeedServiceServer struct{}

func (UnimplementedFeedServiceServer) PullLatest(context.Context, *PullLatestReq) (*PullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullLatest not implemented")
}
func (UnimplementedFeedServiceServer) Pull(context.Context, *PullReq) (*PullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedFeedServiceServer) PullOutBox(context.Context, *PullOutBoxReq) (*PullOutBoxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullOutBox not implemented")
}
func (UnimplementedFeedServiceServer) mustEmbedUnimplementedFeedServiceServer() {}
func (UnimplementedFeedServiceServer) testEmbeddedByValue()                     {}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	// If the following call pancis, it indicates UnimplementedFeedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_PullLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullLatestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).PullLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_PullLatest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).PullLatest(ctx, req.(*PullLatestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).Pull(ctx, req.(*PullReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_PullOutBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullOutBoxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).PullOutBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_PullOutBox_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).PullOutBox(ctx, req.(*PullOutBoxReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feedRpc.feedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pullLatest",
			Handler:    _FeedService_PullLatest_Handler,
		},
		{
			MethodName: "pull",
			Handler:    _FeedService_Pull_Handler,
		},
		{
			MethodName: "pullOutBox",
			Handler:    _FeedService_PullOutBox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed.proto",
}
