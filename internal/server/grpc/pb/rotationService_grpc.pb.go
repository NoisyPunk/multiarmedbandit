// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: rotationService.proto

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

// RotatorClient is the client API for Rotator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RotatorClient interface {
	AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*AddBannerResponse, error)
	AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error)
	AddSlot(ctx context.Context, in *AddSlotRequest, opts ...grpc.CallOption) (*AddSlotResponse, error)
	AddRotation(ctx context.Context, in *AddRotationRequest, opts ...grpc.CallOption) (*AddRotationResponse, error)
	RegisterClick(ctx context.Context, in *RegisterClickRequest, opts ...grpc.CallOption) (*RegisterClickResponse, error)
	ShowBanner(ctx context.Context, in *ShowBannerRequest, opts ...grpc.CallOption) (*ShowBannerResponse, error)
}

type rotatorClient struct {
	cc grpc.ClientConnInterface
}

func NewRotatorClient(cc grpc.ClientConnInterface) RotatorClient {
	return &rotatorClient{cc}
}

func (c *rotatorClient) AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*AddBannerResponse, error) {
	out := new(AddBannerResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error) {
	out := new(AddGroupResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) AddSlot(ctx context.Context, in *AddSlotRequest, opts ...grpc.CallOption) (*AddSlotResponse, error) {
	out := new(AddSlotResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) AddRotation(ctx context.Context, in *AddRotationRequest, opts ...grpc.CallOption) (*AddRotationResponse, error) {
	out := new(AddRotationResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/AddRotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) RegisterClick(ctx context.Context, in *RegisterClickRequest, opts ...grpc.CallOption) (*RegisterClickResponse, error) {
	out := new(RegisterClickResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/RegisterClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotatorClient) ShowBanner(ctx context.Context, in *ShowBannerRequest, opts ...grpc.CallOption) (*ShowBannerResponse, error) {
	out := new(ShowBannerResponse)
	err := c.cc.Invoke(ctx, "/rotator.Rotator/ShowBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RotatorServer is the server API for Rotator service.
// All implementations must embed UnimplementedRotatorServer
// for forward compatibility
type RotatorServer interface {
	AddBanner(context.Context, *AddBannerRequest) (*AddBannerResponse, error)
	AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error)
	AddSlot(context.Context, *AddSlotRequest) (*AddSlotResponse, error)
	AddRotation(context.Context, *AddRotationRequest) (*AddRotationResponse, error)
	RegisterClick(context.Context, *RegisterClickRequest) (*RegisterClickResponse, error)
	ShowBanner(context.Context, *ShowBannerRequest) (*ShowBannerResponse, error)
	mustEmbedUnimplementedRotatorServer()
}

// UnimplementedRotatorServer must be embedded to have forward compatible implementations.
type UnimplementedRotatorServer struct {
}

func (UnimplementedRotatorServer) AddBanner(context.Context, *AddBannerRequest) (*AddBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBanner not implemented")
}
func (UnimplementedRotatorServer) AddGroup(context.Context, *AddGroupRequest) (*AddGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedRotatorServer) AddSlot(context.Context, *AddSlotRequest) (*AddSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSlot not implemented")
}
func (UnimplementedRotatorServer) AddRotation(context.Context, *AddRotationRequest) (*AddRotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRotation not implemented")
}
func (UnimplementedRotatorServer) RegisterClick(context.Context, *RegisterClickRequest) (*RegisterClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClick not implemented")
}
func (UnimplementedRotatorServer) ShowBanner(context.Context, *ShowBannerRequest) (*ShowBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowBanner not implemented")
}
func (UnimplementedRotatorServer) mustEmbedUnimplementedRotatorServer() {}

// UnsafeRotatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RotatorServer will
// result in compilation errors.
type UnsafeRotatorServer interface {
	mustEmbedUnimplementedRotatorServer()
}

func RegisterRotatorServer(s grpc.ServiceRegistrar, srv RotatorServer) {
	s.RegisterService(&Rotator_ServiceDesc, srv)
}

func _Rotator_AddBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddBanner(ctx, req.(*AddBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddGroup(ctx, req.(*AddGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_AddSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddSlot(ctx, req.(*AddSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_AddRotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).AddRotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/AddRotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).AddRotation(ctx, req.(*AddRotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_RegisterClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).RegisterClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/RegisterClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).RegisterClick(ctx, req.(*RegisterClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rotator_ShowBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotatorServer).ShowBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotator.Rotator/ShowBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotatorServer).ShowBanner(ctx, req.(*ShowBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rotator_ServiceDesc is the grpc.ServiceDesc for Rotator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rotator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rotator.Rotator",
	HandlerType: (*RotatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBanner",
			Handler:    _Rotator_AddBanner_Handler,
		},
		{
			MethodName: "AddGroup",
			Handler:    _Rotator_AddGroup_Handler,
		},
		{
			MethodName: "AddSlot",
			Handler:    _Rotator_AddSlot_Handler,
		},
		{
			MethodName: "AddRotation",
			Handler:    _Rotator_AddRotation_Handler,
		},
		{
			MethodName: "RegisterClick",
			Handler:    _Rotator_RegisterClick_Handler,
		},
		{
			MethodName: "ShowBanner",
			Handler:    _Rotator_ShowBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rotationService.proto",
}