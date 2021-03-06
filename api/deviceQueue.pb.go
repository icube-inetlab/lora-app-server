// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deviceQueue.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnqueueDeviceQueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0)
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data (or use the jsonObject when an application codec has been configured).
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// String containing a JSON object (to be enqueued by the application codec).
	JsonObject           string   `protobuf:"bytes,6,opt,name=jsonObject" json:"jsonObject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueDeviceQueueItemRequest) Reset()         { *m = EnqueueDeviceQueueItemRequest{} }
func (m *EnqueueDeviceQueueItemRequest) String() string { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemRequest) ProtoMessage()    {}
func (*EnqueueDeviceQueueItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{0}
}
func (m *EnqueueDeviceQueueItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueDeviceQueueItemRequest.Unmarshal(m, b)
}
func (m *EnqueueDeviceQueueItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueDeviceQueueItemRequest.Marshal(b, m, deterministic)
}
func (dst *EnqueueDeviceQueueItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueDeviceQueueItemRequest.Merge(dst, src)
}
func (m *EnqueueDeviceQueueItemRequest) XXX_Size() int {
	return xxx_messageInfo_EnqueueDeviceQueueItemRequest.Size(m)
}
func (m *EnqueueDeviceQueueItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueDeviceQueueItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueDeviceQueueItemRequest proto.InternalMessageInfo

func (m *EnqueueDeviceQueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *EnqueueDeviceQueueItemRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *EnqueueDeviceQueueItemRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EnqueueDeviceQueueItemRequest) GetJsonObject() string {
	if m != nil {
		return m.JsonObject
	}
	return ""
}

type EnqueueDeviceQueueItemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnqueueDeviceQueueItemResponse) Reset()         { *m = EnqueueDeviceQueueItemResponse{} }
func (m *EnqueueDeviceQueueItemResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemResponse) ProtoMessage()    {}
func (*EnqueueDeviceQueueItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{1}
}
func (m *EnqueueDeviceQueueItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnqueueDeviceQueueItemResponse.Unmarshal(m, b)
}
func (m *EnqueueDeviceQueueItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnqueueDeviceQueueItemResponse.Marshal(b, m, deterministic)
}
func (dst *EnqueueDeviceQueueItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnqueueDeviceQueueItemResponse.Merge(dst, src)
}
func (m *EnqueueDeviceQueueItemResponse) XXX_Size() int {
	return xxx_messageInfo_EnqueueDeviceQueueItemResponse.Size(m)
}
func (m *EnqueueDeviceQueueItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnqueueDeviceQueueItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnqueueDeviceQueueItemResponse proto.InternalMessageInfo

type FlushDeviceQueueRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI               string   `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlushDeviceQueueRequest) Reset()         { *m = FlushDeviceQueueRequest{} }
func (m *FlushDeviceQueueRequest) String() string { return proto.CompactTextString(m) }
func (*FlushDeviceQueueRequest) ProtoMessage()    {}
func (*FlushDeviceQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{2}
}
func (m *FlushDeviceQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlushDeviceQueueRequest.Unmarshal(m, b)
}
func (m *FlushDeviceQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlushDeviceQueueRequest.Marshal(b, m, deterministic)
}
func (dst *FlushDeviceQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlushDeviceQueueRequest.Merge(dst, src)
}
func (m *FlushDeviceQueueRequest) XXX_Size() int {
	return xxx_messageInfo_FlushDeviceQueueRequest.Size(m)
}
func (m *FlushDeviceQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlushDeviceQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlushDeviceQueueRequest proto.InternalMessageInfo

func (m *FlushDeviceQueueRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type FlushDeviceQueueResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlushDeviceQueueResponse) Reset()         { *m = FlushDeviceQueueResponse{} }
func (m *FlushDeviceQueueResponse) String() string { return proto.CompactTextString(m) }
func (*FlushDeviceQueueResponse) ProtoMessage()    {}
func (*FlushDeviceQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{3}
}
func (m *FlushDeviceQueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlushDeviceQueueResponse.Unmarshal(m, b)
}
func (m *FlushDeviceQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlushDeviceQueueResponse.Marshal(b, m, deterministic)
}
func (dst *FlushDeviceQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlushDeviceQueueResponse.Merge(dst, src)
}
func (m *FlushDeviceQueueResponse) XXX_Size() int {
	return xxx_messageInfo_FlushDeviceQueueResponse.Size(m)
}
func (m *FlushDeviceQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlushDeviceQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlushDeviceQueueResponse proto.InternalMessageInfo

type DeviceQueueItem struct {
	// Hex encoded DevEUI of the device.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,3,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the device.
	Confirmed bool `protobuf:"varint,4,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0).
	FPort uint32 `protobuf:"varint,6,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	// FCnt of the queue item.
	FCnt                 uint32   `protobuf:"varint,8,opt,name=fCnt" json:"fCnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceQueueItem) Reset()         { *m = DeviceQueueItem{} }
func (m *DeviceQueueItem) String() string { return proto.CompactTextString(m) }
func (*DeviceQueueItem) ProtoMessage()    {}
func (*DeviceQueueItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{4}
}
func (m *DeviceQueueItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceQueueItem.Unmarshal(m, b)
}
func (m *DeviceQueueItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceQueueItem.Marshal(b, m, deterministic)
}
func (dst *DeviceQueueItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceQueueItem.Merge(dst, src)
}
func (m *DeviceQueueItem) XXX_Size() int {
	return xxx_messageInfo_DeviceQueueItem.Size(m)
}
func (m *DeviceQueueItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceQueueItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceQueueItem proto.InternalMessageInfo

func (m *DeviceQueueItem) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DeviceQueueItem) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DeviceQueueItem) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *DeviceQueueItem) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *DeviceQueueItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DeviceQueueItem) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type ListDeviceQueueItemsRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI               string   `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeviceQueueItemsRequest) Reset()         { *m = ListDeviceQueueItemsRequest{} }
func (m *ListDeviceQueueItemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsRequest) ProtoMessage()    {}
func (*ListDeviceQueueItemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{5}
}
func (m *ListDeviceQueueItemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceQueueItemsRequest.Unmarshal(m, b)
}
func (m *ListDeviceQueueItemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceQueueItemsRequest.Marshal(b, m, deterministic)
}
func (dst *ListDeviceQueueItemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceQueueItemsRequest.Merge(dst, src)
}
func (m *ListDeviceQueueItemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceQueueItemsRequest.Size(m)
}
func (m *ListDeviceQueueItemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceQueueItemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceQueueItemsRequest proto.InternalMessageInfo

func (m *ListDeviceQueueItemsRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type ListDeviceQueueItemsResponse struct {
	Items                []*DeviceQueueItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListDeviceQueueItemsResponse) Reset()         { *m = ListDeviceQueueItemsResponse{} }
func (m *ListDeviceQueueItemsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsResponse) ProtoMessage()    {}
func (*ListDeviceQueueItemsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deviceQueue_7a2ba650b6273f85, []int{6}
}
func (m *ListDeviceQueueItemsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceQueueItemsResponse.Unmarshal(m, b)
}
func (m *ListDeviceQueueItemsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceQueueItemsResponse.Marshal(b, m, deterministic)
}
func (dst *ListDeviceQueueItemsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceQueueItemsResponse.Merge(dst, src)
}
func (m *ListDeviceQueueItemsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeviceQueueItemsResponse.Size(m)
}
func (m *ListDeviceQueueItemsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceQueueItemsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceQueueItemsResponse proto.InternalMessageInfo

func (m *ListDeviceQueueItemsResponse) GetItems() []*DeviceQueueItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EnqueueDeviceQueueItemRequest)(nil), "api.EnqueueDeviceQueueItemRequest")
	proto.RegisterType((*EnqueueDeviceQueueItemResponse)(nil), "api.EnqueueDeviceQueueItemResponse")
	proto.RegisterType((*FlushDeviceQueueRequest)(nil), "api.FlushDeviceQueueRequest")
	proto.RegisterType((*FlushDeviceQueueResponse)(nil), "api.FlushDeviceQueueResponse")
	proto.RegisterType((*DeviceQueueItem)(nil), "api.DeviceQueueItem")
	proto.RegisterType((*ListDeviceQueueItemsRequest)(nil), "api.ListDeviceQueueItemsRequest")
	proto.RegisterType((*ListDeviceQueueItemsResponse)(nil), "api.ListDeviceQueueItemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceQueueClient is the client API for DeviceQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceQueueClient interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*FlushDeviceQueueResponse, error)
	// List lists the items in the device-queue.
	List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error)
}

type deviceQueueClient struct {
	cc *grpc.ClientConn
}

func NewDeviceQueueClient(cc *grpc.ClientConn) DeviceQueueClient {
	return &deviceQueueClient{cc}
}

func (c *deviceQueueClient) Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error) {
	out := new(EnqueueDeviceQueueItemResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceQueue/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) Flush(ctx context.Context, in *FlushDeviceQueueRequest, opts ...grpc.CallOption) (*FlushDeviceQueueResponse, error) {
	out := new(FlushDeviceQueueResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceQueue/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error) {
	out := new(ListDeviceQueueItemsResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceQueue/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceQueue service

type DeviceQueueServer interface {
	// Enqueue adds the given item to the device-queue.
	Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error)
	// Flush flushes the downlink device-queue.
	Flush(context.Context, *FlushDeviceQueueRequest) (*FlushDeviceQueueResponse, error)
	// List lists the items in the device-queue.
	List(context.Context, *ListDeviceQueueItemsRequest) (*ListDeviceQueueItemsResponse, error)
}

func RegisterDeviceQueueServer(s *grpc.Server, srv DeviceQueueServer) {
	s.RegisterService(&_DeviceQueue_serviceDesc, srv)
}

func _DeviceQueue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Enqueue(ctx, req.(*EnqueueDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushDeviceQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Flush(ctx, req.(*FlushDeviceQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).List(ctx, req.(*ListDeviceQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceQueue",
	HandlerType: (*DeviceQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DeviceQueue_Enqueue_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _DeviceQueue_Flush_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceQueue_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceQueue.proto",
}

func init() { proto.RegisterFile("deviceQueue.proto", fileDescriptor_deviceQueue_7a2ba650b6273f85) }

var fileDescriptor_deviceQueue_7a2ba650b6273f85 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x6e, 0xd4, 0x30,
	0x10, 0x95, 0x77, 0xb3, 0xdb, 0x76, 0x0a, 0x42, 0x8c, 0x2a, 0xb0, 0xd2, 0x6c, 0x65, 0x52, 0x09,
	0x45, 0xfb, 0xb1, 0x11, 0x45, 0xfc, 0xf0, 0x0b, 0x45, 0x2a, 0x42, 0x02, 0x22, 0x71, 0x80, 0x34,
	0x99, 0x14, 0x57, 0xad, 0x9d, 0xc6, 0x4e, 0x3f, 0x40, 0xfc, 0x70, 0x05, 0x8e, 0xc0, 0x39, 0x38,
	0x05, 0x57, 0x40, 0x9c, 0x03, 0xc5, 0x89, 0xb4, 0x61, 0xdb, 0xa4, 0x7f, 0xf6, 0xcc, 0x7b, 0xf3,
	0x3c, 0x6f, 0xc6, 0xf0, 0x30, 0xa7, 0x6b, 0x99, 0xd1, 0xc7, 0x9a, 0x6a, 0x5a, 0x95, 0x95, 0xb6,
	0x1a, 0xa7, 0x69, 0x29, 0xfd, 0xe0, 0x4c, 0xeb, 0xb3, 0x0b, 0x8a, 0xd3, 0x52, 0xc6, 0xa9, 0x52,
	0xda, 0xa6, 0x56, 0x6a, 0x65, 0x5a, 0x48, 0xf8, 0x8b, 0xc1, 0xe2, 0x58, 0x5d, 0x35, 0xa4, 0xd7,
	0x6b, 0xfe, 0x89, 0xa5, 0xcb, 0x84, 0xae, 0x6a, 0x32, 0x16, 0x1f, 0xc1, 0x3c, 0xa7, 0xeb, 0xe3,
	0x4f, 0x27, 0x9c, 0x09, 0x16, 0xed, 0x24, 0xdd, 0x0d, 0x03, 0xd8, 0xa9, 0xa8, 0xa0, 0x8a, 0x54,
	0x46, 0x7c, 0xe2, 0x52, 0xeb, 0x40, 0x93, 0xcd, 0xb4, 0x2a, 0x64, 0x75, 0x49, 0x39, 0x9f, 0x0a,
	0x16, 0x6d, 0x27, 0xeb, 0x00, 0xee, 0xc1, 0xac, 0xf8, 0xa0, 0x2b, 0xcb, 0x3d, 0xc1, 0xa2, 0xfb,
	0x49, 0x7b, 0x41, 0x04, 0x2f, 0x4f, 0x6d, 0xca, 0x67, 0x82, 0x45, 0xf7, 0x12, 0x77, 0xc6, 0x03,
	0x80, 0x73, 0xa3, 0xd5, 0xfb, 0xd3, 0x73, 0xca, 0x2c, 0x9f, 0x3b, 0x99, 0x5e, 0x24, 0x14, 0x70,
	0x30, 0xf4, 0x7c, 0x53, 0x6a, 0x65, 0x28, 0x7c, 0x06, 0x8f, 0xdf, 0x5c, 0xd4, 0xe6, 0x73, 0x2f,
	0x7f, 0x47, 0x6b, 0xa1, 0x0f, 0xfc, 0x26, 0xa5, 0x2b, 0xf7, 0x93, 0xc1, 0x83, 0x0d, 0xa9, 0x5e,
	0x9d, 0xc9, 0xb0, 0x45, 0xd3, 0x51, 0x8b, 0xbc, 0x41, 0x8b, 0xe6, 0xb7, 0x59, 0xb4, 0xd5, 0xb3,
	0x08, 0xc1, 0x2b, 0x5e, 0x29, 0xcb, 0xb7, 0x1d, 0xd0, 0x9d, 0xc3, 0x17, 0xb0, 0xff, 0x4e, 0x1a,
	0xbb, 0xf1, 0x50, 0x73, 0x57, 0xe3, 0x6f, 0x21, 0xb8, 0x9d, 0xd6, 0x36, 0x8f, 0x4b, 0x98, 0xc9,
	0x26, 0xc0, 0x99, 0x98, 0x46, 0xbb, 0x47, 0x7b, 0xab, 0xb4, 0x94, 0xab, 0x4d, 0xe3, 0x5b, 0xc8,
	0xd1, 0xdf, 0x09, 0xec, 0xf6, 0x52, 0xf8, 0x05, 0xb6, 0xba, 0x49, 0x61, 0xe8, 0x78, 0xa3, 0x6b,
	0xe7, 0x1f, 0x8e, 0x62, 0xba, 0x61, 0x3c, 0xfd, 0xfe, 0xfb, 0xcf, 0x8f, 0x89, 0x08, 0xf7, 0xdd,
	0x76, 0xb7, 0x1f, 0xc0, 0xc4, 0x5f, 0xdb, 0x6e, 0xbe, 0xc5, 0x8e, 0xfb, 0x92, 0x2d, 0x51, 0xc2,
	0xcc, 0x0d, 0x14, 0x03, 0x57, 0x75, 0x60, 0x1f, 0xfc, 0xc5, 0x40, 0xb6, 0x53, 0x3b, 0x74, 0x6a,
	0x8b, 0xe5, 0x98, 0x1a, 0x96, 0xe0, 0x35, 0x16, 0xa2, 0x70, 0xb5, 0x46, 0x86, 0xe0, 0x3f, 0x19,
	0x41, 0xfc, 0xaf, 0x88, 0x63, 0x8a, 0xa7, 0x73, 0xf7, 0x93, 0x9f, 0xff, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0xca, 0x4f, 0x78, 0x01, 0x04, 0x00, 0x00,
}
