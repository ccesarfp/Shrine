// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/token.proto

package protobuf

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

// TokenClient is the client API for Token service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenClient interface {
	// create token using user data and return jwt
	CreateToken(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	// receive token and return all user data
	GetClaimsByToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*UserResponse, error)
	// receive token ID and return all user data
	GetClaimsByKey(ctx context.Context, in *TokenRequestWithId, opts ...grpc.CallOption) (*UserResponseWithToken, error)
	// receive token and return if is valid
	CheckTokenValidity(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenStatus, error)
}

type tokenClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenClient(cc grpc.ClientConnInterface) TokenClient {
	return &tokenClient{cc}
}

func (c *tokenClient) CreateToken(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/Token/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) GetClaimsByToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/Token/GetClaimsByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) GetClaimsByKey(ctx context.Context, in *TokenRequestWithId, opts ...grpc.CallOption) (*UserResponseWithToken, error) {
	out := new(UserResponseWithToken)
	err := c.cc.Invoke(ctx, "/Token/GetClaimsByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) CheckTokenValidity(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenStatus, error) {
	out := new(TokenStatus)
	err := c.cc.Invoke(ctx, "/Token/CheckTokenValidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServer is the server API for Token service.
// All implementations must embed UnimplementedTokenServer
// for forward compatibility
type TokenServer interface {
	// create token using user data and return jwt
	CreateToken(context.Context, *UserRequest) (*TokenResponse, error)
	// receive token and return all user data
	GetClaimsByToken(context.Context, *TokenRequest) (*UserResponse, error)
	// receive token ID and return all user data
	GetClaimsByKey(context.Context, *TokenRequestWithId) (*UserResponseWithToken, error)
	// receive token and return if is valid
	CheckTokenValidity(context.Context, *TokenRequest) (*TokenStatus, error)
	mustEmbedUnimplementedTokenServer()
}

// UnimplementedTokenServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServer struct {
}

func (UnimplementedTokenServer) CreateToken(context.Context, *UserRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedTokenServer) GetClaimsByToken(context.Context, *TokenRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimsByToken not implemented")
}
func (UnimplementedTokenServer) GetClaimsByKey(context.Context, *TokenRequestWithId) (*UserResponseWithToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimsByKey not implemented")
}
func (UnimplementedTokenServer) CheckTokenValidity(context.Context, *TokenRequest) (*TokenStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTokenValidity not implemented")
}
func (UnimplementedTokenServer) mustEmbedUnimplementedTokenServer() {}

// UnsafeTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServer will
// result in compilation errors.
type UnsafeTokenServer interface {
	mustEmbedUnimplementedTokenServer()
}

func RegisterTokenServer(s grpc.ServiceRegistrar, srv TokenServer) {
	s.RegisterService(&Token_ServiceDesc, srv)
}

func _Token_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Token/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CreateToken(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_GetClaimsByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).GetClaimsByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Token/GetClaimsByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).GetClaimsByToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_GetClaimsByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequestWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).GetClaimsByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Token/GetClaimsByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).GetClaimsByKey(ctx, req.(*TokenRequestWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_CheckTokenValidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).CheckTokenValidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Token/CheckTokenValidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).CheckTokenValidity(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Token_ServiceDesc is the grpc.ServiceDesc for Token service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Token_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Token",
	HandlerType: (*TokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Token_CreateToken_Handler,
		},
		{
			MethodName: "GetClaimsByToken",
			Handler:    _Token_GetClaimsByToken_Handler,
		},
		{
			MethodName: "GetClaimsByKey",
			Handler:    _Token_GetClaimsByKey_Handler,
		},
		{
			MethodName: "CheckTokenValidity",
			Handler:    _Token_CheckTokenValidity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/token.proto",
}
