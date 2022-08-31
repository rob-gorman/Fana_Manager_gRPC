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
	GetFlags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Flags, error)
	GetAudience(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AudienceFullResp, error)
	GetAudiences(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Audiences, error)
	GetAttribute(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AttributeResp, error)
	GetAttributes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Attributes, error)
	GetSDKKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SDKKeys, error)
	GetAuditLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuditLogResp, error)
	CreateFlag(ctx context.Context, in *FlagSubmit, opts ...grpc.CallOption) (*FlagFullResp, error)
	CreateAudience(ctx context.Context, in *AudSubmit, opts ...grpc.CallOption) (*AudienceFullResp, error)
	CreateAttribute(ctx context.Context, in *AttrSubmit, opts ...grpc.CallOption) (*AttributeResp, error)
	UpdateFlag(ctx context.Context, in *FlagUpdate, opts ...grpc.CallOption) (*FlagFullResp, error)
	ToggleFlag(ctx context.Context, in *FlagToggle, opts ...grpc.CallOption) (*Empty, error)
	UpdateAudience(ctx context.Context, in *AudUpdate, opts ...grpc.CallOption) (*AudienceFullResp, error)
	RegenerateSDK(ctx context.Context, in *ID, opts ...grpc.CallOption) (*SDKKey, error)
	DeleteFlag(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	DeleteAudience(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	DeleteAttribute(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
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

func (c *fanaClient) GetFlags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Flags, error) {
	out := new(Flags)
	err := c.cc.Invoke(ctx, "/Fana/GetFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetAudience(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AudienceFullResp, error) {
	out := new(AudienceFullResp)
	err := c.cc.Invoke(ctx, "/Fana/GetAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetAudiences(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Audiences, error) {
	out := new(Audiences)
	err := c.cc.Invoke(ctx, "/Fana/GetAudiences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetAttribute(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AttributeResp, error) {
	out := new(AttributeResp)
	err := c.cc.Invoke(ctx, "/Fana/GetAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetAttributes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Attributes, error) {
	out := new(Attributes)
	err := c.cc.Invoke(ctx, "/Fana/GetAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetSDKKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SDKKeys, error) {
	out := new(SDKKeys)
	err := c.cc.Invoke(ctx, "/Fana/GetSDKKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) GetAuditLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuditLogResp, error) {
	out := new(AuditLogResp)
	err := c.cc.Invoke(ctx, "/Fana/GetAuditLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) CreateFlag(ctx context.Context, in *FlagSubmit, opts ...grpc.CallOption) (*FlagFullResp, error) {
	out := new(FlagFullResp)
	err := c.cc.Invoke(ctx, "/Fana/CreateFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) CreateAudience(ctx context.Context, in *AudSubmit, opts ...grpc.CallOption) (*AudienceFullResp, error) {
	out := new(AudienceFullResp)
	err := c.cc.Invoke(ctx, "/Fana/CreateAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) CreateAttribute(ctx context.Context, in *AttrSubmit, opts ...grpc.CallOption) (*AttributeResp, error) {
	out := new(AttributeResp)
	err := c.cc.Invoke(ctx, "/Fana/CreateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) UpdateFlag(ctx context.Context, in *FlagUpdate, opts ...grpc.CallOption) (*FlagFullResp, error) {
	out := new(FlagFullResp)
	err := c.cc.Invoke(ctx, "/Fana/UpdateFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) ToggleFlag(ctx context.Context, in *FlagToggle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Fana/ToggleFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) UpdateAudience(ctx context.Context, in *AudUpdate, opts ...grpc.CallOption) (*AudienceFullResp, error) {
	out := new(AudienceFullResp)
	err := c.cc.Invoke(ctx, "/Fana/UpdateAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) RegenerateSDK(ctx context.Context, in *ID, opts ...grpc.CallOption) (*SDKKey, error) {
	out := new(SDKKey)
	err := c.cc.Invoke(ctx, "/Fana/RegenerateSDK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) DeleteFlag(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Fana/DeleteFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) DeleteAudience(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Fana/DeleteAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fanaClient) DeleteAttribute(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Fana/DeleteAttribute", in, out, opts...)
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
	GetFlags(context.Context, *Empty) (*Flags, error)
	GetAudience(context.Context, *ID) (*AudienceFullResp, error)
	GetAudiences(context.Context, *Empty) (*Audiences, error)
	GetAttribute(context.Context, *ID) (*AttributeResp, error)
	GetAttributes(context.Context, *Empty) (*Attributes, error)
	GetSDKKeys(context.Context, *Empty) (*SDKKeys, error)
	GetAuditLogs(context.Context, *Empty) (*AuditLogResp, error)
	CreateFlag(context.Context, *FlagSubmit) (*FlagFullResp, error)
	CreateAudience(context.Context, *AudSubmit) (*AudienceFullResp, error)
	CreateAttribute(context.Context, *AttrSubmit) (*AttributeResp, error)
	UpdateFlag(context.Context, *FlagUpdate) (*FlagFullResp, error)
	ToggleFlag(context.Context, *FlagToggle) (*Empty, error)
	UpdateAudience(context.Context, *AudUpdate) (*AudienceFullResp, error)
	RegenerateSDK(context.Context, *ID) (*SDKKey, error)
	DeleteFlag(context.Context, *ID) (*Empty, error)
	DeleteAudience(context.Context, *ID) (*Empty, error)
	DeleteAttribute(context.Context, *ID) (*Empty, error)
	mustEmbedUnimplementedFanaServer()
}

// UnimplementedFanaServer must be embedded to have forward compatible implementations.
type UnimplementedFanaServer struct {
}

func (UnimplementedFanaServer) GetFlag(context.Context, *ID) (*FlagFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}
func (UnimplementedFanaServer) GetFlags(context.Context, *Empty) (*Flags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlags not implemented")
}
func (UnimplementedFanaServer) GetAudience(context.Context, *ID) (*AudienceFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAudience not implemented")
}
func (UnimplementedFanaServer) GetAudiences(context.Context, *Empty) (*Audiences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAudiences not implemented")
}
func (UnimplementedFanaServer) GetAttribute(context.Context, *ID) (*AttributeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttribute not implemented")
}
func (UnimplementedFanaServer) GetAttributes(context.Context, *Empty) (*Attributes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttributes not implemented")
}
func (UnimplementedFanaServer) GetSDKKeys(context.Context, *Empty) (*SDKKeys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSDKKeys not implemented")
}
func (UnimplementedFanaServer) GetAuditLogs(context.Context, *Empty) (*AuditLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditLogs not implemented")
}
func (UnimplementedFanaServer) CreateFlag(context.Context, *FlagSubmit) (*FlagFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlag not implemented")
}
func (UnimplementedFanaServer) CreateAudience(context.Context, *AudSubmit) (*AudienceFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAudience not implemented")
}
func (UnimplementedFanaServer) CreateAttribute(context.Context, *AttrSubmit) (*AttributeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttribute not implemented")
}
func (UnimplementedFanaServer) UpdateFlag(context.Context, *FlagUpdate) (*FlagFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlag not implemented")
}
func (UnimplementedFanaServer) ToggleFlag(context.Context, *FlagToggle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleFlag not implemented")
}
func (UnimplementedFanaServer) UpdateAudience(context.Context, *AudUpdate) (*AudienceFullResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAudience not implemented")
}
func (UnimplementedFanaServer) RegenerateSDK(context.Context, *ID) (*SDKKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateSDK not implemented")
}
func (UnimplementedFanaServer) DeleteFlag(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFlag not implemented")
}
func (UnimplementedFanaServer) DeleteAudience(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAudience not implemented")
}
func (UnimplementedFanaServer) DeleteAttribute(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttribute not implemented")
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

func _Fana_GetFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetFlags(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetAudience(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetAudiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetAudiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetAudiences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetAudiences(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetAttribute(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetAttributes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetSDKKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetSDKKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetSDKKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetSDKKeys(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_GetAuditLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).GetAuditLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/GetAuditLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).GetAuditLogs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_CreateFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagSubmit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).CreateFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/CreateFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).CreateFlag(ctx, req.(*FlagSubmit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_CreateAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AudSubmit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).CreateAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/CreateAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).CreateAudience(ctx, req.(*AudSubmit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_CreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttrSubmit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).CreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/CreateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).CreateAttribute(ctx, req.(*AttrSubmit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_UpdateFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).UpdateFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/UpdateFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).UpdateFlag(ctx, req.(*FlagUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_ToggleFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagToggle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).ToggleFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/ToggleFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).ToggleFlag(ctx, req.(*FlagToggle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_UpdateAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AudUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).UpdateAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/UpdateAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).UpdateAudience(ctx, req.(*AudUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_RegenerateSDK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).RegenerateSDK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/RegenerateSDK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).RegenerateSDK(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_DeleteFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).DeleteFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/DeleteFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).DeleteFlag(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_DeleteAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).DeleteAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/DeleteAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).DeleteAudience(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fana_DeleteAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FanaServer).DeleteAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fana/DeleteAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FanaServer).DeleteAttribute(ctx, req.(*ID))
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
		{
			MethodName: "GetFlags",
			Handler:    _Fana_GetFlags_Handler,
		},
		{
			MethodName: "GetAudience",
			Handler:    _Fana_GetAudience_Handler,
		},
		{
			MethodName: "GetAudiences",
			Handler:    _Fana_GetAudiences_Handler,
		},
		{
			MethodName: "GetAttribute",
			Handler:    _Fana_GetAttribute_Handler,
		},
		{
			MethodName: "GetAttributes",
			Handler:    _Fana_GetAttributes_Handler,
		},
		{
			MethodName: "GetSDKKeys",
			Handler:    _Fana_GetSDKKeys_Handler,
		},
		{
			MethodName: "GetAuditLogs",
			Handler:    _Fana_GetAuditLogs_Handler,
		},
		{
			MethodName: "CreateFlag",
			Handler:    _Fana_CreateFlag_Handler,
		},
		{
			MethodName: "CreateAudience",
			Handler:    _Fana_CreateAudience_Handler,
		},
		{
			MethodName: "CreateAttribute",
			Handler:    _Fana_CreateAttribute_Handler,
		},
		{
			MethodName: "UpdateFlag",
			Handler:    _Fana_UpdateFlag_Handler,
		},
		{
			MethodName: "ToggleFlag",
			Handler:    _Fana_ToggleFlag_Handler,
		},
		{
			MethodName: "UpdateAudience",
			Handler:    _Fana_UpdateAudience_Handler,
		},
		{
			MethodName: "RegenerateSDK",
			Handler:    _Fana_RegenerateSDK_Handler,
		},
		{
			MethodName: "DeleteFlag",
			Handler:    _Fana_DeleteFlag_Handler,
		},
		{
			MethodName: "DeleteAudience",
			Handler:    _Fana_DeleteAudience_Handler,
		},
		{
			MethodName: "DeleteAttribute",
			Handler:    _Fana_DeleteAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fana.proto",
}
