// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: fana.proto

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

// FanaClient is the client API for Fana service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FanaClient interface {
	GetFlag(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FlagFullResp, error)
}

type fanaClient struct {
	cc grpc.ClientConnInterface
}

func NewFanaClient(cc grpc.ClientConnInterface) FanaClient {
	return &fanaClient{cc}
}

func (c *fanaClient) GetFlag(ctx context.Context, in *ID, opts ...grpc.CallOption) (*FlagFullResp, error) {
	out := new(FlagFullResp)
	err := c.cc.Invoke(ctx, "/Fana/GetFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FanaServer is the server API for Fana service.
// All implementations must embed UnimplementedFanaServer
// for forward compatibility
type FanaServer interface {
	GetFlag(context.Context, *ID) (*FlagFullResp, error)
	mustEmbedUnimplementedFanaServer()
}

// UnimplementedFanaServer must be embedded to have forward compatible implementations.
type UnimplementedFanaServer struct {
}

func (UnimplementedFanaServer) GetFlag(context.Context, *ID) (*FlagFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}
func (UnimplementedFanaServer) mustEmbedUnimplementedFanaServer() {}

// UnsafeFanaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FanaServer will
// result in compilation errors.
type UnsafeFanaServer interface {
	mustEmbedUnimplementedFanaServer()
}

func RegisterFanaServer(s grpc.ServiceRegistrar, srv FanaServer) {
	s.RegisterService(&Fana_ServiceDesc, srv)
}

func _Fana_GetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetFlag(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Fana_ServiceDesc is the grpc.ServiceDesc for Fana service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fana_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Fana",
	HandlerType: (*FanaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlag",
			Handler:    _Fana_GetFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fana.proto",
}
