// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc1
// source: auth.proto

package AuthRpc

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
	AuthService_Authentication_FullMethodName = "/AuthRpc.AuthService/Authentication"
	AuthService_RefreshSession_FullMethodName = "/AuthRpc.AuthService/RefreshSession"
	AuthService_DeleteSession_FullMethodName  = "/AuthRpc.AuthService/DeleteSession"
	AuthService_CreateVoucher_FullMethodName  = "/AuthRpc.AuthService/CreateVoucher"
	AuthService_IsActive_FullMethodName       = "/AuthRpc.AuthService/IsActive"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authentication(ctx context.Context, in *AuthenticationReq, opts ...grpc.CallOption) (*AuthenticationResp, error)
	RefreshSession(ctx context.Context, in *RefreshSessionReq, opts ...grpc.CallOption) (*RefreshSessionResp, error)
	DeleteSession(ctx context.Context, in *DeleteSessionReq, opts ...grpc.CallOption) (*DeleteSessionResp, error)
	CreateVoucher(ctx context.Context, in *CreateVoucherReq, opts ...grpc.CallOption) (*CreateVoucherResp, error)
	IsActive(ctx context.Context, in *IsActiveReq, opts ...grpc.CallOption) (*IsActiveResp, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authentication(ctx context.Context, in *AuthenticationReq, opts ...grpc.CallOption) (*AuthenticationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticationResp)
	err := c.cc.Invoke(ctx, AuthService_Authentication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshSession(ctx context.Context, in *RefreshSessionReq, opts ...grpc.CallOption) (*RefreshSessionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshSessionResp)
	err := c.cc.Invoke(ctx, AuthService_RefreshSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionReq, opts ...grpc.CallOption) (*DeleteSessionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSessionResp)
	err := c.cc.Invoke(ctx, AuthService_DeleteSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateVoucher(ctx context.Context, in *CreateVoucherReq, opts ...grpc.CallOption) (*CreateVoucherResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVoucherResp)
	err := c.cc.Invoke(ctx, AuthService_CreateVoucher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) IsActive(ctx context.Context, in *IsActiveReq, opts ...grpc.CallOption) (*IsActiveResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsActiveResp)
	err := c.cc.Invoke(ctx, AuthService_IsActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	Authentication(context.Context, *AuthenticationReq) (*AuthenticationResp, error)
	RefreshSession(context.Context, *RefreshSessionReq) (*RefreshSessionResp, error)
	DeleteSession(context.Context, *DeleteSessionReq) (*DeleteSessionResp, error)
	CreateVoucher(context.Context, *CreateVoucherReq) (*CreateVoucherResp, error)
	IsActive(context.Context, *IsActiveReq) (*IsActiveResp, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Authentication(context.Context, *AuthenticationReq) (*AuthenticationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedAuthServiceServer) RefreshSession(context.Context, *RefreshSessionReq) (*RefreshSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedAuthServiceServer) DeleteSession(context.Context, *DeleteSessionReq) (*DeleteSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedAuthServiceServer) CreateVoucher(context.Context, *CreateVoucherReq) (*CreateVoucherResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVoucher not implemented")
}
func (UnimplementedAuthServiceServer) IsActive(context.Context, *IsActiveReq) (*IsActiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsActive not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Authentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authentication(ctx, req.(*AuthenticationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshSession(ctx, req.(*RefreshSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteSession(ctx, req.(*DeleteSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoucherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateVoucher(ctx, req.(*CreateVoucherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_IsActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsActiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_IsActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsActive(ctx, req.(*IsActiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthRpc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authentication",
			Handler:    _AuthService_Authentication_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _AuthService_RefreshSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthService_DeleteSession_Handler,
		},
		{
			MethodName: "CreateVoucher",
			Handler:    _AuthService_CreateVoucher_Handler,
		},
		{
			MethodName: "IsActive",
			Handler:    _AuthService_IsActive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
