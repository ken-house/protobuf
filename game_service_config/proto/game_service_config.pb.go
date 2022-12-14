// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_service_config.proto

package game_service_config

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

// 请求参数
type GameServiceConfigRequest struct {
	GameId               int32    `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	GameAreaId           int32    `protobuf:"varint,2,opt,name=game_area_id,json=gameAreaId,proto3" json:"game_area_id,omitempty"`
	GameSign             string   `protobuf:"bytes,3,opt,name=game_sign,json=gameSign,proto3" json:"game_sign,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	GameDispatchId       int32    `protobuf:"varint,5,opt,name=game_dispatch_id,json=gameDispatchId,proto3" json:"game_dispatch_id,omitempty"`
	OpType               int32    `protobuf:"varint,6,opt,name=op_type,json=opType,proto3" json:"op_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServiceConfigRequest) Reset()         { *m = GameServiceConfigRequest{} }
func (m *GameServiceConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GameServiceConfigRequest) ProtoMessage()    {}
func (*GameServiceConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9682838bfcd9ad2, []int{0}
}

func (m *GameServiceConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServiceConfigRequest.Unmarshal(m, b)
}
func (m *GameServiceConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServiceConfigRequest.Marshal(b, m, deterministic)
}
func (m *GameServiceConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServiceConfigRequest.Merge(m, src)
}
func (m *GameServiceConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GameServiceConfigRequest.Size(m)
}
func (m *GameServiceConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServiceConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GameServiceConfigRequest proto.InternalMessageInfo

func (m *GameServiceConfigRequest) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *GameServiceConfigRequest) GetGameAreaId() int32 {
	if m != nil {
		return m.GameAreaId
	}
	return 0
}

func (m *GameServiceConfigRequest) GetGameSign() string {
	if m != nil {
		return m.GameSign
	}
	return ""
}

func (m *GameServiceConfigRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GameServiceConfigRequest) GetGameDispatchId() int32 {
	if m != nil {
		return m.GameDispatchId
	}
	return 0
}

func (m *GameServiceConfigRequest) GetOpType() int32 {
	if m != nil {
		return m.OpType
	}
	return 0
}

// 响应参数
type GameServiceConfigResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServiceConfigResponse) Reset()         { *m = GameServiceConfigResponse{} }
func (m *GameServiceConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GameServiceConfigResponse) ProtoMessage()    {}
func (*GameServiceConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9682838bfcd9ad2, []int{1}
}

func (m *GameServiceConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServiceConfigResponse.Unmarshal(m, b)
}
func (m *GameServiceConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServiceConfigResponse.Marshal(b, m, deterministic)
}
func (m *GameServiceConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServiceConfigResponse.Merge(m, src)
}
func (m *GameServiceConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GameServiceConfigResponse.Size(m)
}
func (m *GameServiceConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServiceConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GameServiceConfigResponse proto.InternalMessageInfo

func (m *GameServiceConfigResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*GameServiceConfigRequest)(nil), "game_service_config.GameServiceConfigRequest")
	proto.RegisterType((*GameServiceConfigResponse)(nil), "game_service_config.GameServiceConfigResponse")
}

func init() {
	proto.RegisterFile("game_service_config.proto", fileDescriptor_e9682838bfcd9ad2)
}

var fileDescriptor_e9682838bfcd9ad2 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x89, 0xda, 0xd8, 0x0e, 0x22, 0xb2, 0x8a, 0x4d, 0x2d, 0x42, 0xe8, 0x29, 0x17, 0x23,
	0xd8, 0xa3, 0x27, 0xff, 0x80, 0xe4, 0xe2, 0x21, 0xf5, 0xe4, 0x25, 0xac, 0xc9, 0x34, 0x2e, 0xb4,
	0xd9, 0x35, 0xbb, 0x29, 0xe6, 0x3d, 0x7c, 0x2b, 0x5f, 0x4a, 0x76, 0x36, 0x37, 0x57, 0xf0, 0xf8,
	0xfb, 0xbe, 0xcc, 0x64, 0x66, 0x16, 0x66, 0x35, 0xdf, 0x62, 0xa1, 0xb1, 0xdd, 0x89, 0x12, 0x8b,
	0x52, 0x36, 0x6b, 0x51, 0xa7, 0xaa, 0x95, 0x46, 0xb2, 0x53, 0x8f, 0x5a, 0x7c, 0x07, 0x10, 0x3d,
	0xf1, 0x2d, 0xae, 0x1c, 0x7e, 0x20, 0x9a, 0xe3, 0x47, 0x87, 0xda, 0xb0, 0x29, 0x1c, 0x52, 0x8d,
	0xa8, 0xa2, 0x20, 0x0e, 0x92, 0x51, 0x1e, 0xda, 0x98, 0x55, 0x2c, 0x86, 0x23, 0x12, 0xbc, 0x45,
	0x6e, 0xed, 0x1e, 0x59, 0xb0, 0xec, 0xae, 0x45, 0x9e, 0x55, 0x6c, 0x0e, 0x13, 0xf7, 0x3b, 0x51,
	0x37, 0xd1, 0x7e, 0x1c, 0x24, 0x93, 0x7c, 0x6c, 0xc1, 0x4a, 0xd4, 0x0d, 0x3b, 0x83, 0xd1, 0x8e,
	0x6f, 0x3a, 0x8c, 0x0e, 0x48, 0xb8, 0xc0, 0x12, 0x38, 0xa1, 0x92, 0x4a, 0x68, 0xc5, 0x4d, 0xf9,
	0x6e, 0x1b, 0x8f, 0xa8, 0xf1, 0xb1, 0xe5, 0x8f, 0x03, 0xce, 0x2a, 0x3b, 0x97, 0x54, 0x85, 0xe9,
	0x15, 0x46, 0xa1, 0x9b, 0x4b, 0xaa, 0x97, 0x5e, 0xe1, 0x62, 0x09, 0x33, 0xcf, 0x32, 0x5a, 0xc9,
	0x46, 0x23, 0x3b, 0x87, 0xb0, 0x45, 0xdd, 0x6d, 0x0c, 0x2d, 0x33, 0xce, 0x87, 0x74, 0xf3, 0xe5,
	0x3b, 0xc1, 0x10, 0xd8, 0x27, 0x4c, 0x7f, 0xb9, 0x67, 0x69, 0xc4, 0xba, 0x67, 0x57, 0xa9, 0xef,
	0xd6, 0x7f, 0x1d, 0xf3, 0x22, 0xfd, 0xef, 0xe7, 0x6e, 0xdc, 0xfb, 0xcb, 0xd7, 0x79, 0x7a, 0x4d,
	0x2f, 0x77, 0xeb, 0x29, 0x7c, 0x0b, 0x49, 0x2d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xe0,
	0x72, 0x50, 0xf1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServiceConfigServiceClient is the client API for GameServiceConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServiceConfigServiceClient interface {
	GameServiceConfigNotify(ctx context.Context, in *GameServiceConfigRequest, opts ...grpc.CallOption) (*GameServiceConfigResponse, error)
}

type gameServiceConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceConfigServiceClient(cc grpc.ClientConnInterface) GameServiceConfigServiceClient {
	return &gameServiceConfigServiceClient{cc}
}

func (c *gameServiceConfigServiceClient) GameServiceConfigNotify(ctx context.Context, in *GameServiceConfigRequest, opts ...grpc.CallOption) (*GameServiceConfigResponse, error) {
	out := new(GameServiceConfigResponse)
	err := c.cc.Invoke(ctx, "/game_service_config.GameServiceConfigService/GameServiceConfigNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceConfigServiceServer is the server API for GameServiceConfigService service.
type GameServiceConfigServiceServer interface {
	GameServiceConfigNotify(context.Context, *GameServiceConfigRequest) (*GameServiceConfigResponse, error)
}

// UnimplementedGameServiceConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServiceConfigServiceServer struct {
}

func (*UnimplementedGameServiceConfigServiceServer) GameServiceConfigNotify(ctx context.Context, req *GameServiceConfigRequest) (*GameServiceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GameServiceConfigNotify not implemented")
}

func RegisterGameServiceConfigServiceServer(s *grpc.Server, srv GameServiceConfigServiceServer) {
	s.RegisterService(&_GameServiceConfigService_serviceDesc, srv)
}

func _GameServiceConfigService_GameServiceConfigNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameServiceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceConfigServiceServer).GameServiceConfigNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_service_config.GameServiceConfigService/GameServiceConfigNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceConfigServiceServer).GameServiceConfigNotify(ctx, req.(*GameServiceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameServiceConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game_service_config.GameServiceConfigService",
	HandlerType: (*GameServiceConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GameServiceConfigNotify",
			Handler:    _GameServiceConfigService_GameServiceConfigNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_service_config.proto",
}
