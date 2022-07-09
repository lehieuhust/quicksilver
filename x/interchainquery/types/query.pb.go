// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/interchainquery/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryRequestsRequest struct {
	Pagination   *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	ConnectionId string             `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (m *QueryRequestsRequest) Reset()         { *m = QueryRequestsRequest{} }
func (m *QueryRequestsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequestsRequest) ProtoMessage()    {}
func (*QueryRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4aadfdae61bcbb1, []int{0}
}
func (m *QueryRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequestsRequest.Merge(m, src)
}
func (m *QueryRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequestsRequest proto.InternalMessageInfo

func (m *QueryRequestsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *QueryRequestsRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryRequestsResponse struct {
	// params defines the parameters of the module.
	Queries    []Query             `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRequestsResponse) Reset()         { *m = QueryRequestsResponse{} }
func (m *QueryRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRequestsResponse) ProtoMessage()    {}
func (*QueryRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4aadfdae61bcbb1, []int{1}
}
func (m *QueryRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequestsResponse.Merge(m, src)
}
func (m *QueryRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequestsResponse proto.InternalMessageInfo

func (m *QueryRequestsResponse) GetQueries() []Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *QueryRequestsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequestsRequest)(nil), "quicksilver.interchainquery.v1.QueryRequestsRequest")
	proto.RegisterType((*QueryRequestsResponse)(nil), "quicksilver.interchainquery.v1.QueryRequestsResponse")
}

func init() {
	proto.RegisterFile("quicksilver/interchainquery/v1/query.proto", fileDescriptor_e4aadfdae61bcbb1)
}

var fileDescriptor_e4aadfdae61bcbb1 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0xeb, 0xd3, 0x30,
	0x1c, 0xc6, 0x9b, 0x29, 0x8e, 0x65, 0x7a, 0x29, 0x13, 0xc6, 0x90, 0x3a, 0x26, 0xea, 0x18, 0x9a,
	0xb0, 0xa9, 0x88, 0x17, 0x85, 0x81, 0x8a, 0x37, 0x57, 0x2f, 0xe2, 0x45, 0xd2, 0xee, 0x4b, 0x16,
	0xdc, 0x92, 0xae, 0x49, 0x8b, 0x43, 0xbc, 0xe8, 0x1b, 0x10, 0x7c, 0x11, 0xbe, 0x04, 0x5f, 0x81,
	0xb0, 0xe3, 0xc0, 0x8b, 0x27, 0x91, 0xcd, 0x17, 0x22, 0x4d, 0x5b, 0xf6, 0x47, 0xf8, 0x6d, 0xbf,
	0x53, 0xd3, 0xf6, 0xf3, 0x3c, 0xcf, 0xf7, 0x49, 0x82, 0x7b, 0xf3, 0x44, 0x84, 0xef, 0xb4, 0x98,
	0xa6, 0x10, 0x53, 0x21, 0x0d, 0xc4, 0xe1, 0x84, 0x09, 0x39, 0x4f, 0x20, 0x5e, 0xd0, 0xb4, 0x4f,
	0xed, 0x82, 0x44, 0xb1, 0x32, 0xca, 0xf5, 0x76, 0x58, 0x72, 0xc0, 0x92, 0xb4, 0xdf, 0x6a, 0x70,
	0xc5, 0x95, 0x45, 0x69, 0xb6, 0xca, 0x55, 0xad, 0x6b, 0x5c, 0x29, 0x3e, 0x05, 0xca, 0x22, 0x41,
	0x99, 0x94, 0xca, 0x30, 0x23, 0x94, 0xd4, 0xc5, 0xdf, 0x3b, 0x47, 0xf2, 0x39, 0x48, 0xd0, 0xa2,
	0xa4, 0x7b, 0xa1, 0xd2, 0x33, 0xa5, 0x69, 0xc0, 0x34, 0xd0, 0x92, 0x09, 0xc0, 0xb0, 0x3e, 0x8d,
	0x18, 0x17, 0xd2, 0x5a, 0xe7, 0x6c, 0xe7, 0x33, 0xc2, 0x8d, 0x51, 0x86, 0xf8, 0x30, 0x4f, 0x40,
	0x1b, 0x5d, 0x3c, 0xdd, 0x67, 0x18, 0x6f, 0xe1, 0x26, 0x6a, 0xa3, 0x6e, 0x7d, 0x70, 0x8b, 0xe4,
	0xce, 0x24, 0x73, 0x26, 0x65, 0x23, 0xeb, 0x4c, 0x5e, 0x32, 0x0e, 0x85, 0xd6, 0xdf, 0x51, 0xba,
	0x37, 0xf0, 0x95, 0x50, 0x49, 0x09, 0x61, 0xf6, 0xf6, 0x56, 0x8c, 0x9b, 0x95, 0x36, 0xea, 0xd6,
	0xfc, 0xcb, 0xdb, 0x8f, 0x2f, 0xc6, 0x9d, 0x6f, 0x08, 0x5f, 0x3d, 0x98, 0x42, 0x47, 0x4a, 0x6a,
	0x70, 0x9f, 0xe2, 0x6a, 0x96, 0x23, 0x40, 0x37, 0x51, 0xfb, 0x42, 0xb7, 0x3e, 0xb8, 0x49, 0xce,
	0xde, 0x5f, 0x62, 0x7d, 0x86, 0x17, 0x97, 0xbf, 0xaf, 0x3b, 0x7e, 0xa9, 0x75, 0x9f, 0xef, 0xb5,
	0xa9, 0xd8, 0x36, 0xb7, 0x8f, 0xb6, 0xc9, 0x67, 0xd8, 0xad, 0x33, 0xf8, 0x81, 0x70, 0xcd, 0x26,
	0xbc, 0x8a, 0xd3, 0xd8, 0xfd, 0x8e, 0x70, 0x75, 0x54, 0x44, 0xdc, 0x3f, 0x69, 0xb0, 0x83, 0x6d,
	0x6e, 0x3d, 0x38, 0xa7, 0x2a, 0x1f, 0xa9, 0xf3, 0xe4, 0xd3, 0xcf, 0xbf, 0x5f, 0x2b, 0x8f, 0xdc,
	0x87, 0xf4, 0x84, 0x9b, 0x29, 0x40, 0xd3, 0x0f, 0x7b, 0x87, 0xf0, 0x71, 0xf8, 0x7a, 0xb9, 0xf6,
	0xd0, 0x6a, 0xed, 0xa1, 0x3f, 0x6b, 0x0f, 0x7d, 0xd9, 0x78, 0xce, 0x6a, 0xe3, 0x39, 0xbf, 0x36,
	0x9e, 0xf3, 0xe6, 0x31, 0x17, 0x66, 0x92, 0x04, 0x24, 0x54, 0x33, 0x2a, 0x24, 0x07, 0x99, 0x08,
	0xb3, 0xb8, 0x1b, 0x24, 0x62, 0x3a, 0xde, 0x0b, 0x7b, 0xff, 0x5f, 0x9c, 0x59, 0x44, 0xa0, 0x83,
	0x4b, 0xf6, 0x62, 0xdd, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x07, 0xfc, 0xc2, 0x34, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuerySrvrClient is the client API for QuerySrvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuerySrvrClient interface {
	// Params returns the total set of minting parameters.
	Queries(ctx context.Context, in *QueryRequestsRequest, opts ...grpc.CallOption) (*QueryRequestsResponse, error)
}

type querySrvrClient struct {
	cc grpc1.ClientConn
}

func NewQuerySrvrClient(cc grpc1.ClientConn) QuerySrvrClient {
	return &querySrvrClient{cc}
}

func (c *querySrvrClient) Queries(ctx context.Context, in *QueryRequestsRequest, opts ...grpc.CallOption) (*QueryRequestsResponse, error) {
	out := new(QueryRequestsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.interchainquery.v1.QuerySrvr/Queries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerySrvrServer is the server API for QuerySrvr service.
type QuerySrvrServer interface {
	// Params returns the total set of minting parameters.
	Queries(context.Context, *QueryRequestsRequest) (*QueryRequestsResponse, error)
}

// UnimplementedQuerySrvrServer can be embedded to have forward compatible implementations.
type UnimplementedQuerySrvrServer struct {
}

func (*UnimplementedQuerySrvrServer) Queries(ctx context.Context, req *QueryRequestsRequest) (*QueryRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Queries not implemented")
}

func RegisterQuerySrvrServer(s grpc1.Server, srv QuerySrvrServer) {
	s.RegisterService(&_QuerySrvr_serviceDesc, srv)
}

func _QuerySrvr_Queries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerySrvrServer).Queries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.interchainquery.v1.QuerySrvr/Queries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerySrvrServer).Queries(ctx, req.(*QueryRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuerySrvr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.interchainquery.v1.QuerySrvr",
	HandlerType: (*QuerySrvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Queries",
			Handler:    _QuerySrvr_Queries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/interchainquery/v1/query.proto",
}

func (m *QueryRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Queries) > 0 {
		for iNdEx := len(m.Queries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Queries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequestsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRequestsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
