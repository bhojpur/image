// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ImageUIClient is the client API for ImageUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageUIClient interface {
	// ListEngineSpecs returns a list of Image Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (ImageUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type imageUIClient struct {
	cc grpc.ClientConnInterface
}

func NewImageUIClient(cc grpc.ClientConnInterface) ImageUIClient {
	return &imageUIClient{cc}
}

func (c *imageUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (ImageUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageUI_ServiceDesc.Streams[0], "/v1.ImageUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type imageUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *imageUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.ImageUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageUIServer is the server API for ImageUI service.
// All implementations must embed UnimplementedImageUIServer
// for forward compatibility
type ImageUIServer interface {
	// ListEngineSpecs returns a list of Image Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, ImageUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedImageUIServer()
}

// UnimplementedImageUIServer must be embedded to have forward compatible implementations.
type UnimplementedImageUIServer struct {
}

func (UnimplementedImageUIServer) ListEngineSpecs(*ListEngineSpecsRequest, ImageUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedImageUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedImageUIServer) mustEmbedUnimplementedImageUIServer() {}

// UnsafeImageUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageUIServer will
// result in compilation errors.
type UnsafeImageUIServer interface {
	mustEmbedUnimplementedImageUIServer()
}

func RegisterImageUIServer(s grpc.ServiceRegistrar, srv ImageUIServer) {
	s.RegisterService(&ImageUI_ServiceDesc, srv)
}

func _ImageUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageUIServer).ListEngineSpecs(m, &imageUIListEngineSpecsServer{stream})
}

type ImageUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type imageUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *imageUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ImageUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ImageUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageUI_ServiceDesc is the grpc.ServiceDesc for ImageUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ImageUI",
	HandlerType: (*ImageUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _ImageUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _ImageUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "image-ui.proto",
}
