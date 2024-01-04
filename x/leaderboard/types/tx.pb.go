// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSendIbcTopRank struct {
	Creator          string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	PlayerId         string `protobuf:"bytes,5,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Rank             uint64 `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
	Score            uint64 `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *MsgSendIbcTopRank) Reset()         { *m = MsgSendIbcTopRank{} }
func (m *MsgSendIbcTopRank) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcTopRank) ProtoMessage()    {}
func (*MsgSendIbcTopRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{0}
}
func (m *MsgSendIbcTopRank) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcTopRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcTopRank.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcTopRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcTopRank.Merge(m, src)
}
func (m *MsgSendIbcTopRank) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcTopRank) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcTopRank.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcTopRank proto.InternalMessageInfo

func (m *MsgSendIbcTopRank) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendIbcTopRank) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *MsgSendIbcTopRank) GetRank() uint64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *MsgSendIbcTopRank) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type MsgSendIbcTopRankResponse struct {
}

func (m *MsgSendIbcTopRankResponse) Reset()         { *m = MsgSendIbcTopRankResponse{} }
func (m *MsgSendIbcTopRankResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendIbcTopRankResponse) ProtoMessage()    {}
func (*MsgSendIbcTopRankResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abcbe4eb090e075c, []int{1}
}
func (m *MsgSendIbcTopRankResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendIbcTopRankResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendIbcTopRankResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendIbcTopRankResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendIbcTopRankResponse.Merge(m, src)
}
func (m *MsgSendIbcTopRankResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendIbcTopRankResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendIbcTopRankResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendIbcTopRankResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendIbcTopRank)(nil), "cosmonaut.leaderboard.leaderboard.MsgSendIbcTopRank")
	proto.RegisterType((*MsgSendIbcTopRankResponse)(nil), "cosmonaut.leaderboard.leaderboard.MsgSendIbcTopRankResponse")
}

func init() { proto.RegisterFile("leaderboard/tx.proto", fileDescriptor_abcbe4eb090e075c) }

var fileDescriptor_abcbe4eb090e075c = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4f, 0x72, 0x31,
	0x14, 0xc6, 0xe9, 0xcb, 0xbf, 0x97, 0x0e, 0x46, 0x1b, 0x86, 0x8a, 0xa6, 0x41, 0x26, 0xe2, 0x70,
	0x49, 0x94, 0xd1, 0xc9, 0xb8, 0x30, 0xe0, 0x70, 0x65, 0x72, 0xeb, 0xed, 0x3d, 0x01, 0x02, 0xb7,
	0xa7, 0x69, 0x4b, 0x02, 0xbb, 0xa3, 0x83, 0x1f, 0xcb, 0x11, 0x37, 0x47, 0x03, 0x5f, 0xc4, 0x58,
	0x03, 0x5e, 0xbd, 0x83, 0x89, 0xdb, 0xf3, 0xfc, 0x4e, 0xce, 0x39, 0x3d, 0x7d, 0x68, 0x73, 0x0e,
	0x32, 0x05, 0x9b, 0xa0, 0xb4, 0x69, 0xcf, 0x2f, 0x23, 0x63, 0xd1, 0x23, 0x3b, 0x53, 0xe8, 0x32,
	0xd4, 0x72, 0xe1, 0xa3, 0x5c, 0x3d, 0xaf, 0x3b, 0x2f, 0x84, 0x1e, 0x0d, 0xdd, 0xf8, 0x0e, 0x74,
	0x3a, 0x48, 0xd4, 0x08, 0x4d, 0x2c, 0xf5, 0x8c, 0x71, 0x5a, 0x57, 0x16, 0xa4, 0x47, 0xcb, 0x49,
	0x9b, 0x74, 0x1b, 0xf1, 0xce, 0x32, 0x46, 0x2b, 0x06, 0xad, 0xe7, 0xff, 0x02, 0x0e, 0x9a, 0x9d,
	0xd2, 0x86, 0x9a, 0x48, 0xad, 0x61, 0x3e, 0xb8, 0xe1, 0xe5, 0x50, 0xf8, 0x02, 0xec, 0x9c, 0x1e,
	0xfa, 0x69, 0x06, 0xb8, 0xf0, 0xa3, 0x69, 0x06, 0xce, 0xcb, 0xcc, 0xf0, 0x4a, 0x9b, 0x74, 0x2b,
	0x71, 0x81, 0xb3, 0x16, 0xfd, 0x6f, 0xe6, 0x72, 0x05, 0x76, 0x90, 0xf2, 0x6a, 0x18, 0xb4, 0xf7,
	0x1f, 0x9b, 0xad, 0xd4, 0x33, 0x5e, 0x0b, 0xbd, 0x41, 0xb3, 0x26, 0xad, 0x3a, 0x85, 0x16, 0x78,
	0x3d, 0xc0, 0x4f, 0xd3, 0x39, 0xa1, 0xc7, 0x85, 0x93, 0x62, 0x70, 0x06, 0xb5, 0x83, 0x8b, 0x47,
	0x42, 0xcb, 0x43, 0x37, 0x66, 0x0f, 0x84, 0x1e, 0xfc, 0xb8, 0xba, 0x1f, 0xfd, 0xfa, 0x5f, 0x51,
	0x61, 0x70, 0xeb, 0xea, 0x2f, 0x5d, 0xbb, 0xe7, 0x5c, 0xdf, 0x3e, 0x6f, 0x04, 0x59, 0x6f, 0x04,
	0x79, 0xdb, 0x08, 0xf2, 0xb4, 0x15, 0xa5, 0xf5, 0x56, 0x94, 0x5e, 0xb7, 0xa2, 0x74, 0xdf, 0x1f,
	0x4f, 0xfd, 0x64, 0x91, 0x44, 0x0a, 0xb3, 0xde, 0x7e, 0x43, 0x2f, 0x9f, 0xf3, 0xf2, 0x9b, 0xf3,
	0x2b, 0x03, 0x2e, 0xa9, 0x85, 0xe4, 0x2f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x80, 0x4a, 0xd0,
	0x58, 0x11, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendIbcTopRank(ctx context.Context, in *MsgSendIbcTopRank, opts ...grpc.CallOption) (*MsgSendIbcTopRankResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendIbcTopRank(ctx context.Context, in *MsgSendIbcTopRank, opts ...grpc.CallOption) (*MsgSendIbcTopRankResponse, error) {
	out := new(MsgSendIbcTopRankResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.leaderboard.leaderboard.Msg/SendIbcTopRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendIbcTopRank(context.Context, *MsgSendIbcTopRank) (*MsgSendIbcTopRankResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendIbcTopRank(ctx context.Context, req *MsgSendIbcTopRank) (*MsgSendIbcTopRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIbcTopRank not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendIbcTopRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendIbcTopRank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendIbcTopRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.leaderboard.leaderboard.Msg/SendIbcTopRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendIbcTopRank(ctx, req.(*MsgSendIbcTopRank))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.leaderboard.leaderboard.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIbcTopRank",
			Handler:    _Msg_SendIbcTopRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaderboard/tx.proto",
}

func (m *MsgSendIbcTopRank) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcTopRank) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcTopRank) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Score != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x38
	}
	if m.Rank != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PlayerId) > 0 {
		i -= len(m.PlayerId)
		copy(dAtA[i:], m.PlayerId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PlayerId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendIbcTopRankResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendIbcTopRankResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendIbcTopRankResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendIbcTopRank) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.PlayerId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovTx(uint64(m.Rank))
	}
	if m.Score != 0 {
		n += 1 + sovTx(uint64(m.Score))
	}
	return n
}

func (m *MsgSendIbcTopRankResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendIbcTopRank) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendIbcTopRank: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcTopRank: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSendIbcTopRankResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendIbcTopRankResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendIbcTopRankResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
