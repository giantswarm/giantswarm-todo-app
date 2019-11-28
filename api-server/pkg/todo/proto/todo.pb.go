// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todomgrpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Todo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Todo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type TodoIdReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoIdReq) Reset()         { *m = TodoIdReq{} }
func (m *TodoIdReq) String() string { return proto.CompactTextString(m) }
func (*TodoIdReq) ProtoMessage()    {}
func (*TodoIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *TodoIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoIdReq.Unmarshal(m, b)
}
func (m *TodoIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoIdReq.Marshal(b, m, deterministic)
}
func (m *TodoIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoIdReq.Merge(m, src)
}
func (m *TodoIdReq) XXX_Size() int {
	return xxx_messageInfo_TodoIdReq.Size(m)
}
func (m *TodoIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_TodoIdReq proto.InternalMessageInfo

func (m *TodoIdReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TodoIdReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type ListTodosReq struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosReq) Reset()         { *m = ListTodosReq{} }
func (m *ListTodosReq) String() string { return proto.CompactTextString(m) }
func (*ListTodosReq) ProtoMessage()    {}
func (*ListTodosReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *ListTodosReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosReq.Unmarshal(m, b)
}
func (m *ListTodosReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosReq.Marshal(b, m, deterministic)
}
func (m *ListTodosReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosReq.Merge(m, src)
}
func (m *ListTodosReq) XXX_Size() int {
	return xxx_messageInfo_ListTodosReq.Size(m)
}
func (m *ListTodosReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosReq proto.InternalMessageInfo

func (m *ListTodosReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type DeleteTodoRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRes) Reset()         { *m = DeleteTodoRes{} }
func (m *DeleteTodoRes) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRes) ProtoMessage()    {}
func (*DeleteTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *DeleteTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRes.Unmarshal(m, b)
}
func (m *DeleteTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRes.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRes.Merge(m, src)
}
func (m *DeleteTodoRes) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRes.Size(m)
}
func (m *DeleteTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRes proto.InternalMessageInfo

func (m *DeleteTodoRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo_mgr.Todo")
	proto.RegisterType((*TodoIdReq)(nil), "todo_mgr.TodoIdReq")
	proto.RegisterType((*ListTodosReq)(nil), "todo_mgr.ListTodosReq")
	proto.RegisterType((*DeleteTodoRes)(nil), "todo_mgr.DeleteTodoRes")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xd9, 0x7d, 0xf3, 0xda, 0x64, 0xaa, 0x3d, 0x8c, 0xa2, 0x8b, 0xa7, 0x50, 0x3c, 0x44,
	0x28, 0xc1, 0x3f, 0x78, 0xf1, 0xa8, 0x82, 0x08, 0x7a, 0x59, 0xea, 0xc5, 0x8b, 0xa4, 0xdd, 0x21,
	0x04, 0x6c, 0x36, 0xee, 0xae, 0xe8, 0x87, 0xf0, 0x43, 0xcb, 0x6e, 0x69, 0x92, 0xb6, 0x1e, 0xbc,
	0xcd, 0x33, 0xf3, 0x7b, 0x66, 0x66, 0x87, 0x05, 0x70, 0x5a, 0xe9, 0xbc, 0x31, 0xda, 0x69, 0x8c,
	0x7d, 0xfc, 0xba, 0x28, 0xcd, 0x78, 0x0a, 0xd1, 0x54, 0x2b, 0x8d, 0x23, 0xe0, 0x95, 0x12, 0x2c,
	0x65, 0x59, 0x24, 0x79, 0xa5, 0x10, 0x21, 0x72, 0xf4, 0xe5, 0x04, 0x4f, 0x59, 0x96, 0xc8, 0x10,
	0xfb, 0x9c, 0xd2, 0x35, 0x89, 0x7f, 0x29, 0xcb, 0x62, 0x19, 0x62, 0x3c, 0x80, 0xff, 0xfa, 0xb3,
	0x26, 0x23, 0xa2, 0x00, 0x2e, 0xc5, 0xf8, 0x1c, 0x12, 0xdf, 0xf5, 0x41, 0x49, 0x7a, 0xdf, 0x6a,
	0xdd, 0x5a, 0x78, 0xdf, 0x72, 0x02, 0xbb, 0x8f, 0x95, 0x75, 0xde, 0x66, 0xbd, 0xab, 0xa5, 0x58,
	0x9f, 0x3a, 0x85, 0xbd, 0x3b, 0x7a, 0x23, 0x47, 0x9e, 0x93, 0x64, 0x51, 0xc0, 0xc0, 0x7e, 0xcc,
	0xe7, 0x64, 0x6d, 0x00, 0x63, 0xb9, 0x92, 0x17, 0xdf, 0x1c, 0x86, 0x9e, 0x7a, 0x2a, 0xea, 0xa2,
	0x24, 0x83, 0x13, 0x80, 0x5b, 0x43, 0xc5, 0xd2, 0x8a, 0xa3, 0x7c, 0x75, 0x82, 0xdc, 0xeb, 0xe3,
	0x0d, 0x8d, 0x57, 0x90, 0xb4, 0xeb, 0xe0, 0x61, 0x57, 0xec, 0xef, 0xb8, 0x69, 0x3a, 0x63, 0x98,
	0xc3, 0xe0, 0x9e, 0x02, 0x80, 0xfb, 0xeb, 0xc5, 0x70, 0x8b, 0xad, 0x31, 0x13, 0x80, 0xe7, 0x46,
	0xfd, 0x75, 0xa9, 0x6b, 0x80, 0xee, 0xf5, 0xbf, 0x0f, 0x38, 0xea, 0x92, 0x6b, 0x87, 0xba, 0x19,
	0xbe, 0x24, 0xbe, 0xb2, 0x28, 0x4d, 0x33, 0x9b, 0xed, 0x84, 0x6f, 0x70, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x93, 0x4f, 0x16, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoManagerClient is the client API for TodoManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoManagerClient interface {
	CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	ListTodos(ctx context.Context, in *ListTodosReq, opts ...grpc.CallOption) (TodoManager_ListTodosClient, error)
	GetTodo(ctx context.Context, in *TodoIdReq, opts ...grpc.CallOption) (*Todo, error)
	UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	DeleteTodo(ctx context.Context, in *TodoIdReq, opts ...grpc.CallOption) (*DeleteTodoRes, error)
}

type todoManagerClient struct {
	cc *grpc.ClientConn
}

func NewTodoManagerClient(cc *grpc.ClientConn) TodoManagerClient {
	return &todoManagerClient{cc}
}

func (c *todoManagerClient) CreateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todo_mgr.TodoManager/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) ListTodos(ctx context.Context, in *ListTodosReq, opts ...grpc.CallOption) (TodoManager_ListTodosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoManager_serviceDesc.Streams[0], "/todo_mgr.TodoManager/ListTodos", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoManagerListTodosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoManager_ListTodosClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type todoManagerListTodosClient struct {
	grpc.ClientStream
}

func (x *todoManagerListTodosClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoManagerClient) GetTodo(ctx context.Context, in *TodoIdReq, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todo_mgr.TodoManager/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) UpdateTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todo_mgr.TodoManager/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoManagerClient) DeleteTodo(ctx context.Context, in *TodoIdReq, opts ...grpc.CallOption) (*DeleteTodoRes, error) {
	out := new(DeleteTodoRes)
	err := c.cc.Invoke(ctx, "/todo_mgr.TodoManager/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoManagerServer is the server API for TodoManager service.
type TodoManagerServer interface {
	CreateTodo(context.Context, *Todo) (*Todo, error)
	ListTodos(*ListTodosReq, TodoManager_ListTodosServer) error
	GetTodo(context.Context, *TodoIdReq) (*Todo, error)
	UpdateTodo(context.Context, *Todo) (*Todo, error)
	DeleteTodo(context.Context, *TodoIdReq) (*DeleteTodoRes, error)
}

// UnimplementedTodoManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTodoManagerServer struct {
}

func (*UnimplementedTodoManagerServer) CreateTodo(ctx context.Context, req *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoManagerServer) ListTodos(req *ListTodosReq, srv TodoManager_ListTodosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
func (*UnimplementedTodoManagerServer) GetTodo(ctx context.Context, req *TodoIdReq) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (*UnimplementedTodoManagerServer) UpdateTodo(ctx context.Context, req *Todo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoManagerServer) DeleteTodo(ctx context.Context, req *TodoIdReq) (*DeleteTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func RegisterTodoManagerServer(s *grpc.Server, srv TodoManagerServer) {
	s.RegisterService(&_TodoManager_serviceDesc, srv)
}

func _TodoManager_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo_mgr.TodoManager/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).CreateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_ListTodos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTodosReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoManagerServer).ListTodos(m, &todoManagerListTodosServer{stream})
}

type TodoManager_ListTodosServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type todoManagerListTodosServer struct {
	grpc.ServerStream
}

func (x *todoManagerListTodosServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

func _TodoManager_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo_mgr.TodoManager/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).GetTodo(ctx, req.(*TodoIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo_mgr.TodoManager/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).UpdateTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoManager_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoManagerServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo_mgr.TodoManager/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoManagerServer).DeleteTodo(ctx, req.(*TodoIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo_mgr.TodoManager",
	HandlerType: (*TodoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoManager_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoManager_GetTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoManager_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoManager_DeleteTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTodos",
			Handler:       _TodoManager_ListTodos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}