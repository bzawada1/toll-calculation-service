// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: types/ptypes.proto

package types

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

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorClient interface {
	Aggregator(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*None, error)
}

type aggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorClient(cc grpc.ClientConnInterface) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) Aggregator(ctx context.Context, in *AggregateRequest, opts ...grpc.CallOption) (*None, error) {
	out := new(None)
	err := c.cc.Invoke(ctx, "/Aggregator/Aggregator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
// All implementations must embed UnimplementedAggregatorServer
// for forward compatibility
type AggregatorServer interface {
	Aggregator(context.Context, *AggregateRequest) (*None, error)
	mustEmbedUnimplementedAggregatorServer()
}

// UnimplementedAggregatorServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServer struct {
}

func (UnimplementedAggregatorServer) Aggregator(context.Context, *AggregateRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregator not implemented")
}
func (UnimplementedAggregatorServer) mustEmbedUnimplementedAggregatorServer() {}

// UnsafeAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServer will
// result in compilation errors.
type UnsafeAggregatorServer interface {
	mustEmbedUnimplementedAggregatorServer()
}

func RegisterAggregatorServer(s grpc.ServiceRegistrar, srv AggregatorServer) {
	s.RegisterService(&Aggregator_ServiceDesc, srv)
}

func _Aggregator_Aggregator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).Aggregator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aggregator/Aggregator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).Aggregator(ctx, req.(*AggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aggregator_ServiceDesc is the grpc.ServiceDesc for Aggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Aggregator",
			Handler:    _Aggregator_Aggregator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/ptypes.proto",
}
