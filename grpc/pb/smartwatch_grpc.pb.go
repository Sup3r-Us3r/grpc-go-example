// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protofile/smartwatch.proto

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

// SmartwatchServiceClient is the client API for SmartwatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartwatchServiceClient interface {
	BeatsPerMinute(ctx context.Context, in *BeatsPerMinuteRequest, opts ...grpc.CallOption) (SmartwatchService_BeatsPerMinuteClient, error)
}

type smartwatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartwatchServiceClient(cc grpc.ClientConnInterface) SmartwatchServiceClient {
	return &smartwatchServiceClient{cc}
}

func (c *smartwatchServiceClient) BeatsPerMinute(ctx context.Context, in *BeatsPerMinuteRequest, opts ...grpc.CallOption) (SmartwatchService_BeatsPerMinuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &SmartwatchService_ServiceDesc.Streams[0], "/gogrpc.SmartwatchService/BeatsPerMinute", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartwatchServiceBeatsPerMinuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartwatchService_BeatsPerMinuteClient interface {
	Recv() (*BeatsPerMinuteResponse, error)
	grpc.ClientStream
}

type smartwatchServiceBeatsPerMinuteClient struct {
	grpc.ClientStream
}

func (x *smartwatchServiceBeatsPerMinuteClient) Recv() (*BeatsPerMinuteResponse, error) {
	m := new(BeatsPerMinuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SmartwatchServiceServer is the server API for SmartwatchService service.
// All implementations must embed UnimplementedSmartwatchServiceServer
// for forward compatibility
type SmartwatchServiceServer interface {
	BeatsPerMinute(*BeatsPerMinuteRequest, SmartwatchService_BeatsPerMinuteServer) error
	mustEmbedUnimplementedSmartwatchServiceServer()
}

// UnimplementedSmartwatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmartwatchServiceServer struct {
}

func (UnimplementedSmartwatchServiceServer) BeatsPerMinute(*BeatsPerMinuteRequest, SmartwatchService_BeatsPerMinuteServer) error {
	return status.Errorf(codes.Unimplemented, "method BeatsPerMinute not implemented")
}
func (UnimplementedSmartwatchServiceServer) mustEmbedUnimplementedSmartwatchServiceServer() {}

// UnsafeSmartwatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartwatchServiceServer will
// result in compilation errors.
type UnsafeSmartwatchServiceServer interface {
	mustEmbedUnimplementedSmartwatchServiceServer()
}

func RegisterSmartwatchServiceServer(s grpc.ServiceRegistrar, srv SmartwatchServiceServer) {
	s.RegisterService(&SmartwatchService_ServiceDesc, srv)
}

func _SmartwatchService_BeatsPerMinute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BeatsPerMinuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartwatchServiceServer).BeatsPerMinute(m, &smartwatchServiceBeatsPerMinuteServer{stream})
}

type SmartwatchService_BeatsPerMinuteServer interface {
	Send(*BeatsPerMinuteResponse) error
	grpc.ServerStream
}

type smartwatchServiceBeatsPerMinuteServer struct {
	grpc.ServerStream
}

func (x *smartwatchServiceBeatsPerMinuteServer) Send(m *BeatsPerMinuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SmartwatchService_ServiceDesc is the grpc.ServiceDesc for SmartwatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartwatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpc.SmartwatchService",
	HandlerType: (*SmartwatchServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BeatsPerMinute",
			Handler:       _SmartwatchService_BeatsPerMinute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protofile/smartwatch.proto",
}