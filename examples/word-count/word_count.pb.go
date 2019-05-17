// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: word_count.proto

package word_count

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import protocol "go.gazette.dev/core/protocol"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_LiveRamp_gazette_v2_pkg_consumer "go.gazette.dev/core/consumer"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NGramCount struct {
	NGram                NGram    `protobuf:"bytes,1,opt,name=n_gram,json=nGram,proto3,casttype=NGram" json:"n_gram,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NGramCount) Reset()         { *m = NGramCount{} }
func (m *NGramCount) String() string { return proto.CompactTextString(m) }
func (*NGramCount) ProtoMessage()    {}
func (*NGramCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_word_count_d83b0a860cc001ff, []int{0}
}
func (m *NGramCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NGramCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NGramCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NGramCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NGramCount.Merge(dst, src)
}
func (m *NGramCount) XXX_Size() int {
	return m.ProtoSize()
}
func (m *NGramCount) XXX_DiscardUnknown() {
	xxx_messageInfo_NGramCount.DiscardUnknown(m)
}

var xxx_messageInfo_NGramCount proto.InternalMessageInfo

func (m *NGramCount) GetNGram() NGram {
	if m != nil {
		return m.NGram
	}
	return ""
}

func (m *NGramCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type PublishRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_word_count_d83b0a860cc001ff, []int{1}
}
func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(dst, src)
}
func (m *PublishRequest) XXX_Size() int {
	return m.ProtoSize()
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PublishResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_word_count_d83b0a860cc001ff, []int{2}
}
func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(dst, src)
}
func (m *PublishResponse) XXX_Size() int {
	return m.ProtoSize()
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

type QueryRequest struct {
	// Header attached by a proxy-ing peer. Not directly set by clients.
	Header *protocol.Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// NGram prefix to query.
	Prefix NGram `protobuf:"bytes,2,opt,name=prefix,proto3,casttype=NGram" json:"prefix,omitempty"`
	// Shard to query. Optional; if not set, shard is inferred from |prefix|'s current mapping.
	Shard                github_com_LiveRamp_gazette_v2_pkg_consumer.ShardID `protobuf:"bytes,3,opt,name=shard,proto3,casttype=go.gazette.dev/core/consumer.ShardID" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_word_count_d83b0a860cc001ff, []int{3}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.ProtoSize()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetHeader() *protocol.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *QueryRequest) GetPrefix() NGram {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *QueryRequest) GetShard() github_com_LiveRamp_gazette_v2_pkg_consumer.ShardID {
	if m != nil {
		return m.Shard
	}
	return ""
}

type QueryResponse struct {
	Grams                []NGramCount `protobuf:"bytes,1,rep,name=grams" json:"grams"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_word_count_d83b0a860cc001ff, []int{4}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(dst, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.ProtoSize()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetGrams() []NGramCount {
	if m != nil {
		return m.Grams
	}
	return nil
}

func init() {
	proto.RegisterType((*NGramCount)(nil), "word_count.NGramCount")
	proto.RegisterType((*PublishRequest)(nil), "word_count.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "word_count.PublishResponse")
	proto.RegisterType((*QueryRequest)(nil), "word_count.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "word_count.QueryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NGramClient is the client API for NGram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NGramClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type nGramClient struct {
	cc *grpc.ClientConn
}

func NewNGramClient(cc *grpc.ClientConn) NGramClient {
	return &nGramClient{cc}
}

func (c *nGramClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/word_count.NGram/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nGramClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/word_count.NGram/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NGramServer is the server API for NGram service.
type NGramServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

func RegisterNGramServer(s *grpc.Server, srv NGramServer) {
	s.RegisterService(&_NGram_serviceDesc, srv)
}

func _NGram_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NGramServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word_count.NGram/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NGramServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NGram_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NGramServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word_count.NGram/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NGramServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NGram_serviceDesc = grpc.ServiceDesc{
	ServiceName: "word_count.NGram",
	HandlerType: (*NGramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _NGram_Publish_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _NGram_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "word_count.proto",
}

func (m *NGramCount) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NGramCount) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NGram) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(len(m.NGram)))
		i += copy(dAtA[i:], m.NGram)
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PublishRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublishRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PublishResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublishResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(m.Header.ProtoSize()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Prefix) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(len(m.Prefix)))
		i += copy(dAtA[i:], m.Prefix)
	}
	if len(m.Shard) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWordCount(dAtA, i, uint64(len(m.Shard)))
		i += copy(dAtA[i:], m.Shard)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Grams) > 0 {
		for _, msg := range m.Grams {
			dAtA[i] = 0xa
			i++
			i = encodeVarintWordCount(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintWordCount(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NGramCount) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NGram)
	if l > 0 {
		n += 1 + l + sovWordCount(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovWordCount(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PublishRequest) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovWordCount(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PublishResponse) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueryRequest) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.ProtoSize()
		n += 1 + l + sovWordCount(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovWordCount(uint64(l))
	}
	l = len(m.Shard)
	if l > 0 {
		n += 1 + l + sovWordCount(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueryResponse) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Grams) > 0 {
		for _, e := range m.Grams {
			l = e.ProtoSize()
			n += 1 + l + sovWordCount(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWordCount(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWordCount(x uint64) (n int) {
	return sovWordCount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NGramCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordCount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NGramCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NGramCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NGram", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NGram = NGram(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWordCount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWordCount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublishRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordCount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublishRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublishRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWordCount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWordCount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PublishResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordCount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PublishResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublishResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipWordCount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWordCount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordCount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &protocol.Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = NGram(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shard", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shard = github_com_LiveRamp_gazette_v2_pkg_consumer.ShardID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWordCount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWordCount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordCount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWordCount
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grams = append(m.Grams, NGramCount{})
			if err := m.Grams[len(m.Grams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWordCount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWordCount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWordCount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWordCount
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
					return 0, ErrIntOverflowWordCount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWordCount
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWordCount
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWordCount
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWordCount(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWordCount = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWordCount   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("word_count.proto", fileDescriptor_word_count_d83b0a860cc001ff) }

var fileDescriptor_word_count_d83b0a860cc001ff = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x75, 0x7c, 0x05, 0xf3, 0x6e, 0xab, 0xd6, 0x89, 0x31, 0x88, 0x49, 0x8b, 0xc4, 0x05, 0x1b,
	0x21, 0xa1, 0x0b, 0x13, 0xe3, 0x8a, 0x36, 0x51, 0x13, 0x35, 0x3a, 0x7e, 0x40, 0x03, 0x74, 0x0a,
	0xc4, 0xc2, 0xe0, 0x30, 0xd4, 0xea, 0x1f, 0xf8, 0x17, 0xae, 0xfd, 0x92, 0x2e, 0xfd, 0x82, 0x2e,
	0xea, 0x5f, 0x74, 0x65, 0x98, 0xc1, 0x96, 0x26, 0x5d, 0xbc, 0xd5, 0xdc, 0x7b, 0xcf, 0x39, 0x77,
	0xce, 0x9c, 0x81, 0xe1, 0x37, 0xc6, 0x17, 0xf3, 0x98, 0xd5, 0x85, 0x70, 0x4b, 0xce, 0x04, 0xc3,
	0x70, 0x9a, 0x98, 0x2f, 0x93, 0x4c, 0xa4, 0x75, 0xe4, 0xc6, 0x2c, 0xf7, 0xde, 0x65, 0x6b, 0x4a,
	0xc2, 0xbc, 0xf4, 0x92, 0xf0, 0x07, 0x15, 0x82, 0x7a, 0x6b, 0xdf, 0x2b, 0xbf, 0x24, 0x9e, 0x54,
	0xc5, 0x6c, 0x75, 0x2c, 0xd4, 0x1e, 0xf3, 0x79, 0x47, 0x9b, 0xb0, 0x84, 0x29, 0x3c, 0xaa, 0x97,
	0xb2, 0x93, 0x8d, 0xac, 0x14, 0xdd, 0x9e, 0x01, 0x7c, 0x78, 0xcd, 0xc3, 0x7c, 0xda, 0x5c, 0x8c,
	0x2d, 0xd0, 0x8b, 0x79, 0xc2, 0xc3, 0xdc, 0x40, 0x16, 0x72, 0xae, 0x83, 0xeb, 0xc3, 0x6e, 0xac,
	0x49, 0x9c, 0x68, 0x45, 0x73, 0xe0, 0x87, 0xa0, 0x49, 0x8f, 0xc6, 0x6d, 0x0b, 0x39, 0x3d, 0xa2,
	0x1a, 0xfb, 0x19, 0xdc, 0xfb, 0x58, 0x47, 0xab, 0xac, 0x4a, 0x09, 0xfd, 0x5a, 0xd3, 0x4a, 0x60,
	0x0c, 0x3d, 0x41, 0x37, 0x42, 0xed, 0x21, 0xb2, 0xb6, 0x1f, 0xc0, 0xfd, 0x23, 0xab, 0x2a, 0x59,
	0x51, 0x51, 0xfb, 0x37, 0x82, 0xc1, 0xa7, 0x9a, 0xf2, 0xef, 0xff, 0x75, 0x0e, 0xe8, 0x29, 0x0d,
	0x17, 0x94, 0x4b, 0x65, 0xdf, 0x1f, 0xba, 0xc7, 0xf7, 0xbd, 0x91, 0x73, 0xd2, 0xe2, 0xf8, 0x29,
	0xe8, 0x25, 0xa7, 0xcb, 0x6c, 0x23, 0xad, 0x9c, 0x79, 0x6d, 0x01, 0xfc, 0x1e, 0xb4, 0x2a, 0x0d,
	0xf9, 0xc2, 0xb8, 0x92, 0x8c, 0x17, 0x87, 0xdd, 0x78, 0x72, 0x83, 0x68, 0x63, 0x56, 0x54, 0x75,
	0x4e, 0xb9, 0xfb, 0xb9, 0x11, 0xbf, 0x9d, 0x11, 0xb5, 0xc5, 0x9e, 0xc2, 0xdd, 0xd6, 0xab, 0x72,
	0x8f, 0x7d, 0xd0, 0x9a, 0xb0, 0x2a, 0x03, 0x59, 0x57, 0x4e, 0xdf, 0x7f, 0xe4, 0x76, 0x7e, 0xf5,
	0x94, 0x6a, 0xd0, 0xdb, 0xee, 0xc6, 0xb7, 0x88, 0xa2, 0xfa, 0x3f, 0x11, 0x28, 0x97, 0x38, 0x80,
	0x3b, 0x6d, 0x1c, 0xd8, 0xec, 0x2a, 0xcf, 0x93, 0x34, 0x9f, 0x5c, 0xc4, 0x5a, 0x07, 0xaf, 0x40,
	0x93, 0x96, 0xb0, 0xd1, 0x65, 0x75, 0x13, 0x35, 0x1f, 0x5f, 0x40, 0x94, 0x3a, 0x18, 0x6c, 0xf7,
	0x23, 0xf4, 0x67, 0x3f, 0x42, 0xbf, 0xfe, 0x8e, 0x50, 0xa4, 0xcb, 0xa4, 0x27, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x8a, 0xa8, 0xd2, 0x0b, 0x9c, 0x02, 0x00, 0x00,
}