// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: todo-service.proto

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

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoServiceClient interface {
	// Read all todo tasks
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
// All implementations must embed UnimplementedToDoServiceServer
// for forward compatibility
type ToDoServiceServer interface {
	// Read all todo tasks
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	// Create new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedToDoServiceServer()
}

// UnimplementedToDoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (UnimplementedToDoServiceServer) ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (UnimplementedToDoServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedToDoServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedToDoServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedToDoServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}

// UnsafeToDoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoServiceServer will
// result in compilation errors.
type UnsafeToDoServiceServer interface {
	mustEmbedUnimplementedToDoServiceServer()
}

func RegisterToDoServiceServer(s grpc.ServiceRegistrar, srv ToDoServiceServer) {
	s.RegisterService(&ToDoService_ServiceDesc, srv)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoService_ServiceDesc is the grpc.ServiceDesc for ToDoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-service.proto",
}
