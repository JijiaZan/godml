// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pyserver

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

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushReply, error)
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullReply, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushReply, error) {
	out := new(PushReply)
	err := c.cc.Invoke(ctx, "/pyserver.Worker/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullReply, error) {
	out := new(PullReply)
	err := c.cc.Invoke(ctx, "/pyserver.Worker/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	Push(context.Context, *PushRequest) (*PushReply, error)
	Pull(context.Context, *PullRequest) (*PullReply, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) Push(context.Context, *PushRequest) (*PushReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedWorkerServer) Pull(context.Context, *PullRequest) (*PullReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pyserver.Worker/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pyserver.Worker/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pyserver.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Worker_Push_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _Worker_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pyserver/pyserver.proto",
}
