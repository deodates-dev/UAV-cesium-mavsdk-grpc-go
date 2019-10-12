// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MAVSDK-Proto/protos/geofence/geofence.proto

package mavsdk_rpc_geofence

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

type Polygon_Type int32

const (
	Polygon_INCLUSION Polygon_Type = 0
	Polygon_EXCLUSION Polygon_Type = 1
)

var Polygon_Type_name = map[int32]string{
	0: "INCLUSION",
	1: "EXCLUSION",
}

var Polygon_Type_value = map[string]int32{
	"INCLUSION": 0,
	"EXCLUSION": 1,
}

func (x Polygon_Type) String() string {
	return proto.EnumName(Polygon_Type_name, int32(x))
}

func (Polygon_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{1, 0}
}

// Possible results returned for geofence requests.
type GeofenceResult_Result int32

const (
	GeofenceResult_UNKNOWN                 GeofenceResult_Result = 0
	GeofenceResult_SUCCESS                 GeofenceResult_Result = 1
	GeofenceResult_ERROR                   GeofenceResult_Result = 2
	GeofenceResult_TOO_MANY_GEOFENCE_ITEMS GeofenceResult_Result = 3
	GeofenceResult_BUSY                    GeofenceResult_Result = 4
	GeofenceResult_TIMEOUT                 GeofenceResult_Result = 5
	GeofenceResult_INVALID_ARGUMENT        GeofenceResult_Result = 6
)

var GeofenceResult_Result_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "ERROR",
	3: "TOO_MANY_GEOFENCE_ITEMS",
	4: "BUSY",
	5: "TIMEOUT",
	6: "INVALID_ARGUMENT",
}

var GeofenceResult_Result_value = map[string]int32{
	"UNKNOWN":                 0,
	"SUCCESS":                 1,
	"ERROR":                   2,
	"TOO_MANY_GEOFENCE_ITEMS": 3,
	"BUSY":                    4,
	"TIMEOUT":                 5,
	"INVALID_ARGUMENT":        6,
}

func (x GeofenceResult_Result) String() string {
	return proto.EnumName(GeofenceResult_Result_name, int32(x))
}

func (GeofenceResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{4, 0}
}

// Point type.
type Point struct {
	LatitudeDeg          float64  `protobuf:"fixed64,1,opt,name=latitude_deg,json=latitudeDeg,proto3" json:"latitude_deg,omitempty"`
	LongitudeDeg         float64  `protobuf:"fixed64,2,opt,name=longitude_deg,json=longitudeDeg,proto3" json:"longitude_deg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitudeDeg() float64 {
	if m != nil {
		return m.LatitudeDeg
	}
	return 0
}

func (m *Point) GetLongitudeDeg() float64 {
	if m != nil {
		return m.LongitudeDeg
	}
	return 0
}

// Polygon type.
type Polygon struct {
	Points               []*Point     `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Type                 Polygon_Type `protobuf:"varint,2,opt,name=type,proto3,enum=mavsdk.rpc.geofence.Polygon_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Polygon) Reset()         { *m = Polygon{} }
func (m *Polygon) String() string { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()    {}
func (*Polygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{1}
}

func (m *Polygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Polygon.Unmarshal(m, b)
}
func (m *Polygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Polygon.Marshal(b, m, deterministic)
}
func (m *Polygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Polygon.Merge(m, src)
}
func (m *Polygon) XXX_Size() int {
	return xxx_messageInfo_Polygon.Size(m)
}
func (m *Polygon) XXX_DiscardUnknown() {
	xxx_messageInfo_Polygon.DiscardUnknown(m)
}

var xxx_messageInfo_Polygon proto.InternalMessageInfo

func (m *Polygon) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Polygon) GetType() Polygon_Type {
	if m != nil {
		return m.Type
	}
	return Polygon_INCLUSION
}

type UploadGeofenceRequest struct {
	Polygons             []*Polygon `protobuf:"bytes,1,rep,name=polygons,proto3" json:"polygons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UploadGeofenceRequest) Reset()         { *m = UploadGeofenceRequest{} }
func (m *UploadGeofenceRequest) String() string { return proto.CompactTextString(m) }
func (*UploadGeofenceRequest) ProtoMessage()    {}
func (*UploadGeofenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{2}
}

func (m *UploadGeofenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadGeofenceRequest.Unmarshal(m, b)
}
func (m *UploadGeofenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadGeofenceRequest.Marshal(b, m, deterministic)
}
func (m *UploadGeofenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadGeofenceRequest.Merge(m, src)
}
func (m *UploadGeofenceRequest) XXX_Size() int {
	return xxx_messageInfo_UploadGeofenceRequest.Size(m)
}
func (m *UploadGeofenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadGeofenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadGeofenceRequest proto.InternalMessageInfo

func (m *UploadGeofenceRequest) GetPolygons() []*Polygon {
	if m != nil {
		return m.Polygons
	}
	return nil
}

type UploadGeofenceResponse struct {
	GeofenceResult       *GeofenceResult `protobuf:"bytes,1,opt,name=geofence_result,json=geofenceResult,proto3" json:"geofence_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UploadGeofenceResponse) Reset()         { *m = UploadGeofenceResponse{} }
func (m *UploadGeofenceResponse) String() string { return proto.CompactTextString(m) }
func (*UploadGeofenceResponse) ProtoMessage()    {}
func (*UploadGeofenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{3}
}

func (m *UploadGeofenceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadGeofenceResponse.Unmarshal(m, b)
}
func (m *UploadGeofenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadGeofenceResponse.Marshal(b, m, deterministic)
}
func (m *UploadGeofenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadGeofenceResponse.Merge(m, src)
}
func (m *UploadGeofenceResponse) XXX_Size() int {
	return xxx_messageInfo_UploadGeofenceResponse.Size(m)
}
func (m *UploadGeofenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadGeofenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadGeofenceResponse proto.InternalMessageInfo

func (m *UploadGeofenceResponse) GetGeofenceResult() *GeofenceResult {
	if m != nil {
		return m.GeofenceResult
	}
	return nil
}

// Result type.
type GeofenceResult struct {
	Result               GeofenceResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.geofence.GeofenceResult_Result" json:"result,omitempty"`
	ResultStr            string                `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GeofenceResult) Reset()         { *m = GeofenceResult{} }
func (m *GeofenceResult) String() string { return proto.CompactTextString(m) }
func (*GeofenceResult) ProtoMessage()    {}
func (*GeofenceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed091f248bcd496, []int{4}
}

func (m *GeofenceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeofenceResult.Unmarshal(m, b)
}
func (m *GeofenceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeofenceResult.Marshal(b, m, deterministic)
}
func (m *GeofenceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeofenceResult.Merge(m, src)
}
func (m *GeofenceResult) XXX_Size() int {
	return xxx_messageInfo_GeofenceResult.Size(m)
}
func (m *GeofenceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GeofenceResult.DiscardUnknown(m)
}

var xxx_messageInfo_GeofenceResult proto.InternalMessageInfo

func (m *GeofenceResult) GetResult() GeofenceResult_Result {
	if m != nil {
		return m.Result
	}
	return GeofenceResult_UNKNOWN
}

func (m *GeofenceResult) GetResultStr() string {
	if m != nil {
		return m.ResultStr
	}
	return ""
}

func init() {
	proto.RegisterEnum("mavsdk.rpc.geofence.Polygon_Type", Polygon_Type_name, Polygon_Type_value)
	proto.RegisterEnum("mavsdk.rpc.geofence.GeofenceResult_Result", GeofenceResult_Result_name, GeofenceResult_Result_value)
	proto.RegisterType((*Point)(nil), "mavsdk.rpc.geofence.Point")
	proto.RegisterType((*Polygon)(nil), "mavsdk.rpc.geofence.Polygon")
	proto.RegisterType((*UploadGeofenceRequest)(nil), "mavsdk.rpc.geofence.UploadGeofenceRequest")
	proto.RegisterType((*UploadGeofenceResponse)(nil), "mavsdk.rpc.geofence.UploadGeofenceResponse")
	proto.RegisterType((*GeofenceResult)(nil), "mavsdk.rpc.geofence.GeofenceResult")
}

func init() {
	proto.RegisterFile("MAVSDK-Proto/protos/geofence/geofence.proto", fileDescriptor_6ed091f248bcd496)
}

var fileDescriptor_6ed091f248bcd496 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0x36, 0x87, 0x66, 0xdc, 0xb8, 0xd6, 0x72, 0xaa, 0x0a, 0x48, 0xad, 0xc3, 0x45,
	0x45, 0x85, 0x2b, 0x19, 0x21, 0x71, 0x9b, 0x83, 0x89, 0xac, 0xc6, 0x76, 0xd8, 0xb5, 0x0b, 0xbd,
	0xb2, 0x42, 0xb2, 0xb5, 0xa2, 0x1a, 0xaf, 0xb1, 0x37, 0x45, 0xb9, 0xe1, 0x41, 0x78, 0x46, 0x1e,
	0x02, 0x79, 0xe3, 0x0d, 0x0d, 0x0a, 0xa8, 0x37, 0xd6, 0xce, 0xec, 0xff, 0x7f, 0x33, 0x9e, 0xd5,
	0xc0, 0xb9, 0xdb, 0xbd, 0x22, 0x83, 0xcb, 0x37, 0xe3, 0x9c, 0x71, 0x76, 0x91, 0x95, 0xdf, 0xe2,
	0x22, 0xa6, 0xec, 0x86, 0xa6, 0x53, 0xba, 0x3e, 0x98, 0xe2, 0x02, 0x3d, 0xfa, 0x3a, 0xb9, 0x2b,
	0x66, 0xb7, 0x66, 0x9e, 0x4d, 0x4d, 0x79, 0x65, 0xf8, 0x50, 0x1f, 0xb3, 0x79, 0xca, 0xd1, 0x29,
	0x1c, 0x24, 0x13, 0x3e, 0xe7, 0x8b, 0x19, 0x8d, 0x66, 0x34, 0x3e, 0x52, 0x4e, 0x94, 0x33, 0x05,
	0xab, 0x32, 0x37, 0xa0, 0x31, 0xea, 0x40, 0x3b, 0x61, 0x69, 0xfc, 0x47, 0xb3, 0x2b, 0x34, 0x07,
	0xeb, 0xe4, 0x80, 0xc6, 0xc6, 0x4f, 0x05, 0x9a, 0x63, 0x96, 0x2c, 0x63, 0x96, 0x22, 0x0b, 0x1a,
	0x59, 0x09, 0x2f, 0x8e, 0x94, 0x93, 0xbd, 0x33, 0xd5, 0x3a, 0x36, 0xb7, 0xb4, 0x60, 0x8a, 0xfa,
	0xb8, 0x52, 0xa2, 0x77, 0x50, 0xe3, 0xcb, 0x8c, 0x0a, 0xb6, 0x66, 0x9d, 0xfe, 0xc3, 0x21, 0xf8,
	0x66, 0xb0, 0xcc, 0x28, 0x16, 0x72, 0xe3, 0x15, 0xd4, 0xca, 0x08, 0xb5, 0xa1, 0xe5, 0x78, 0xfd,
	0x51, 0x48, 0x1c, 0xdf, 0xd3, 0x77, 0xca, 0xd0, 0xfe, 0x2c, 0x43, 0xc5, 0xf8, 0x08, 0x4f, 0xc2,
	0x2c, 0x61, 0x93, 0xd9, 0xb0, 0x42, 0x61, 0xfa, 0x6d, 0x41, 0x0b, 0x8e, 0xde, 0xc3, 0x7e, 0xb6,
	0x82, 0xca, 0x5e, 0x5f, 0xfc, 0xaf, 0x32, 0x5e, 0xab, 0x8d, 0x1b, 0x78, 0xfa, 0x37, 0xb2, 0xc8,
	0x58, 0x5a, 0x50, 0x34, 0x82, 0x43, 0xe9, 0x8b, 0x72, 0x5a, 0x2c, 0x12, 0x2e, 0x86, 0xaa, 0x5a,
	0x9d, 0xad, 0xe8, 0x7b, 0xfe, 0x45, 0xc2, 0xb1, 0x16, 0x6f, 0xc4, 0xc6, 0x2f, 0x05, 0xb4, 0x4d,
	0x09, 0xea, 0x41, 0xe3, 0x1e, 0x57, 0xb3, 0x5e, 0x3f, 0x80, 0x6b, 0x56, 0xf8, 0xca, 0x89, 0x5e,
	0x02, 0xac, 0x4e, 0x51, 0xc1, 0x73, 0x31, 0xf4, 0x16, 0x6e, 0xad, 0x32, 0x84, 0xe7, 0xc6, 0x77,
	0x68, 0x54, 0xc5, 0x54, 0x68, 0x86, 0xde, 0xa5, 0xe7, 0x7f, 0x2a, 0xc7, 0xaa, 0x42, 0x93, 0x84,
	0xfd, 0xbe, 0x4d, 0x88, 0xae, 0xa0, 0x16, 0xd4, 0x6d, 0x8c, 0x7d, 0xac, 0xef, 0xa2, 0xe7, 0xf0,
	0x2c, 0xf0, 0xfd, 0xc8, 0xed, 0x7a, 0xd7, 0xd1, 0xd0, 0xf6, 0x3f, 0xd8, 0x5e, 0xdf, 0x8e, 0x9c,
	0xc0, 0x76, 0x89, 0xbe, 0x87, 0xf6, 0xa1, 0xd6, 0x0b, 0xc9, 0xb5, 0x5e, 0x2b, 0xed, 0x81, 0xe3,
	0xda, 0x7e, 0x18, 0xe8, 0x75, 0xf4, 0x18, 0x74, 0xc7, 0xbb, 0xea, 0x8e, 0x9c, 0x41, 0xd4, 0xc5,
	0xc3, 0xd0, 0xb5, 0xbd, 0x40, 0x6f, 0x58, 0x3f, 0xe0, 0x50, 0x36, 0x4e, 0x68, 0x7e, 0x37, 0x9f,
	0x52, 0x74, 0x0b, 0xda, 0xe6, 0xa4, 0xd1, 0xf6, 0x1f, 0xde, 0xfa, 0xc2, 0xc7, 0xe7, 0x0f, 0xd2,
	0xae, 0x9e, 0xce, 0xd8, 0xe9, 0x75, 0x00, 0xcd, 0x99, 0xb4, 0x48, 0x79, 0xaf, 0x2d, 0x95, 0x62,
	0xe1, 0xbe, 0x34, 0xc4, 0x62, 0xbd, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x85, 0x1c, 0xd2, 0x2c,
	0x87, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeofenceServiceClient is the client API for GeofenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeofenceServiceClient interface {
	//
	// Upload a geofence.
	//
	// Polygons are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error)
}

type geofenceServiceClient struct {
	cc *grpc.ClientConn
}

func NewGeofenceServiceClient(cc *grpc.ClientConn) GeofenceServiceClient {
	return &geofenceServiceClient{cc}
}

func (c *geofenceServiceClient) UploadGeofence(ctx context.Context, in *UploadGeofenceRequest, opts ...grpc.CallOption) (*UploadGeofenceResponse, error) {
	out := new(UploadGeofenceResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeofenceServiceServer is the server API for GeofenceService service.
type GeofenceServiceServer interface {
	//
	// Upload a geofence.
	//
	// Polygons are uploaded to a drone. Once uploaded, the geofence will remain
	// on the drone even if a connection is lost.
	UploadGeofence(context.Context, *UploadGeofenceRequest) (*UploadGeofenceResponse, error)
}

// UnimplementedGeofenceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeofenceServiceServer struct {
}

func (*UnimplementedGeofenceServiceServer) UploadGeofence(ctx context.Context, req *UploadGeofenceRequest) (*UploadGeofenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadGeofence not implemented")
}

func RegisterGeofenceServiceServer(s *grpc.Server, srv GeofenceServiceServer) {
	s.RegisterService(&_GeofenceService_serviceDesc, srv)
}

func _GeofenceService_UploadGeofence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadGeofenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.geofence.GeofenceService/UploadGeofence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeofenceServiceServer).UploadGeofence(ctx, req.(*UploadGeofenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeofenceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.geofence.GeofenceService",
	HandlerType: (*GeofenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadGeofence",
			Handler:    _GeofenceService_UploadGeofence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MAVSDK-Proto/protos/geofence/geofence.proto",
}
