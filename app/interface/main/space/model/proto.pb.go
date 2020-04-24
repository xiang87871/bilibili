// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AidReason struct {
	Aid                  int64    `protobuf:"varint,1,opt,name=Aid,proto3" json:"aid"`
	Reason               string   `protobuf:"bytes,2,opt,name=Reason,proto3" json:"reason"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AidReason) Reset()         { *m = AidReason{} }
func (m *AidReason) String() string { return proto.CompactTextString(m) }
func (*AidReason) ProtoMessage()    {}
func (*AidReason) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_bb83dde5bdf5b6ae, []int{0}
}
func (m *AidReason) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AidReason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AidReason.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AidReason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AidReason.Merge(dst, src)
}
func (m *AidReason) XXX_Size() int {
	return m.Size()
}
func (m *AidReason) XXX_DiscardUnknown() {
	xxx_messageInfo_AidReason.DiscardUnknown(m)
}

var xxx_messageInfo_AidReason proto.InternalMessageInfo

type AidReasons struct {
	List                 []*AidReason `protobuf:"bytes,1,rep,name=List" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AidReasons) Reset()         { *m = AidReasons{} }
func (m *AidReasons) String() string { return proto.CompactTextString(m) }
func (*AidReasons) ProtoMessage()    {}
func (*AidReasons) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_bb83dde5bdf5b6ae, []int{1}
}
func (m *AidReasons) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AidReasons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AidReasons.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AidReasons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AidReasons.Merge(dst, src)
}
func (m *AidReasons) XXX_Size() int {
	return m.Size()
}
func (m *AidReasons) XXX_DiscardUnknown() {
	xxx_messageInfo_AidReasons.DiscardUnknown(m)
}

var xxx_messageInfo_AidReasons proto.InternalMessageInfo

type ThemeDetail struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Icon                 string   `protobuf:"bytes,3,opt,name=Icon,proto3" json:"icon"`
	TopPhoto             string   `protobuf:"bytes,4,opt,name=TopPhoto,proto3" json:"top_photo"`
	BgImg                string   `protobuf:"bytes,5,opt,name=BgImg,proto3" json:"bg_img"`
	IsActivated          int64    `protobuf:"varint,6,opt,name=IsActivated,proto3" json:"is_activated"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThemeDetail) Reset()         { *m = ThemeDetail{} }
func (m *ThemeDetail) String() string { return proto.CompactTextString(m) }
func (*ThemeDetail) ProtoMessage()    {}
func (*ThemeDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_bb83dde5bdf5b6ae, []int{2}
}
func (m *ThemeDetail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThemeDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThemeDetail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ThemeDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThemeDetail.Merge(dst, src)
}
func (m *ThemeDetail) XXX_Size() int {
	return m.Size()
}
func (m *ThemeDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ThemeDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ThemeDetail proto.InternalMessageInfo

type ThemeDetails struct {
	List                 []*ThemeDetail `protobuf:"bytes,1,rep,name=List" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ThemeDetails) Reset()         { *m = ThemeDetails{} }
func (m *ThemeDetails) String() string { return proto.CompactTextString(m) }
func (*ThemeDetails) ProtoMessage()    {}
func (*ThemeDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_bb83dde5bdf5b6ae, []int{3}
}
func (m *ThemeDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThemeDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThemeDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ThemeDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThemeDetails.Merge(dst, src)
}
func (m *ThemeDetails) XXX_Size() int {
	return m.Size()
}
func (m *ThemeDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ThemeDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ThemeDetails proto.InternalMessageInfo

type Notice struct {
	Notice               string   `protobuf:"bytes,1,opt,name=Notice,proto3" json:"notice"`
	IsForbid             int64    `protobuf:"varint,2,opt,name=IsForbid,proto3" json:"is_forbid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notice) Reset()         { *m = Notice{} }
func (m *Notice) String() string { return proto.CompactTextString(m) }
func (*Notice) ProtoMessage()    {}
func (*Notice) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_bb83dde5bdf5b6ae, []int{4}
}
func (m *Notice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Notice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Notice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Notice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notice.Merge(dst, src)
}
func (m *Notice) XXX_Size() int {
	return m.Size()
}
func (m *Notice) XXX_DiscardUnknown() {
	xxx_messageInfo_Notice.DiscardUnknown(m)
}

var xxx_messageInfo_Notice proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AidReason)(nil), "AidReason")
	proto.RegisterType((*AidReasons)(nil), "AidReasons")
	proto.RegisterType((*ThemeDetail)(nil), "ThemeDetail")
	proto.RegisterType((*ThemeDetails)(nil), "ThemeDetails")
	proto.RegisterType((*Notice)(nil), "Notice")
}
func (m *AidReason) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AidReason) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Aid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.Aid))
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AidReasons) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AidReasons) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProto(dAtA, i, uint64(msg.Size()))
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

func (m *ThemeDetail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThemeDetail) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.ID))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Icon) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Icon)))
		i += copy(dAtA[i:], m.Icon)
	}
	if len(m.TopPhoto) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.TopPhoto)))
		i += copy(dAtA[i:], m.TopPhoto)
	}
	if len(m.BgImg) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.BgImg)))
		i += copy(dAtA[i:], m.BgImg)
	}
	if m.IsActivated != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.IsActivated))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ThemeDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThemeDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProto(dAtA, i, uint64(msg.Size()))
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

func (m *Notice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Notice) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProto(dAtA, i, uint64(len(m.Notice)))
		i += copy(dAtA[i:], m.Notice)
	}
	if m.IsForbid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProto(dAtA, i, uint64(m.IsForbid))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintProto(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AidReason) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Aid != 0 {
		n += 1 + sovProto(uint64(m.Aid))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AidReasons) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovProto(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ThemeDetail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovProto(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.TopPhoto)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	l = len(m.BgImg)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	if m.IsActivated != 0 {
		n += 1 + sovProto(uint64(m.IsActivated))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ThemeDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovProto(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Notice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Notice)
	if l > 0 {
		n += 1 + l + sovProto(uint64(l))
	}
	if m.IsForbid != 0 {
		n += 1 + sovProto(uint64(m.IsForbid))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProto(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProto(x uint64) (n int) {
	return sovProto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AidReason) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: AidReason: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AidReason: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aid", wireType)
			}
			m.Aid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Aid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *AidReasons) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: AidReasons: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AidReasons: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &AidReason{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *ThemeDetail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: ThemeDetail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThemeDetail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopPhoto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TopPhoto = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BgImg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BgImg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActivated", wireType)
			}
			m.IsActivated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsActivated |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *ThemeDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: ThemeDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThemeDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &ThemeDetail{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func (m *Notice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto
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
			return fmt.Errorf("proto: Notice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
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
				return ErrInvalidLengthProto
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Notice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsForbid", wireType)
			}
			m.IsForbid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsForbid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto
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
func skipProto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
					return 0, ErrIntOverflowProto
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
				return 0, ErrInvalidLengthProto
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProto
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
				next, err := skipProto(dAtA[start:])
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
	ErrInvalidLengthProto = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto.proto", fileDescriptor_proto_bb83dde5bdf5b6ae) }

var fileDescriptor_proto_bb83dde5bdf5b6ae = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x8a, 0x9b, 0x40,
	0x14, 0xc6, 0x33, 0x6a, 0x6c, 0x72, 0x4c, 0xa1, 0xcc, 0x45, 0xb1, 0xa5, 0xa8, 0x78, 0x95, 0x85,
	0xd6, 0x2d, 0xdb, 0x27, 0x88, 0x84, 0x82, 0xa5, 0x2c, 0x65, 0x58, 0x28, 0xf4, 0x46, 0xfc, 0x33,
	0x6b, 0x06, 0x56, 0x47, 0x32, 0xb3, 0x7d, 0x96, 0x3e, 0x52, 0x2e, 0xf3, 0x04, 0xd2, 0xd8, 0x3b,
	0x9f, 0xa2, 0x38, 0x1a, 0x93, 0xbd, 0x91, 0x73, 0x7e, 0xdf, 0x87, 0xe7, 0x9c, 0x8f, 0x01, 0xab,
	0xde, 0x73, 0xc9, 0x03, 0xf5, 0x7d, 0xff, 0xa9, 0x60, 0x72, 0xf7, 0x9c, 0x06, 0x19, 0x2f, 0x6f,
	0x0b, 0x5e, 0xf0, 0x5b, 0x85, 0xd3, 0xe7, 0x47, 0xd5, 0xa9, 0x46, 0x55, 0x83, 0xdd, 0xff, 0x06,
	0xcb, 0x0d, 0xcb, 0x09, 0x4d, 0x04, 0xaf, 0xf0, 0x3b, 0xd0, 0x37, 0x2c, 0xb7, 0x91, 0x87, 0xd6,
	0x7a, 0xf8, 0xaa, 0x6b, 0x5c, 0x3d, 0x61, 0x39, 0xe9, 0x19, 0xf6, 0xc1, 0x1c, 0x4c, 0xb6, 0xe6,
	0xa1, 0xf5, 0x32, 0x84, 0xae, 0x71, 0xcd, 0xbd, 0x22, 0x64, 0x54, 0xfc, 0x8f, 0x00, 0xd3, 0xbf,
	0x04, 0x76, 0xc0, 0xf8, 0xce, 0x84, 0xb4, 0x91, 0xa7, 0xaf, 0xad, 0x3b, 0x08, 0x26, 0x89, 0x28,
	0xee, 0xb7, 0x08, 0xac, 0x87, 0x1d, 0x2d, 0xe9, 0x96, 0xca, 0x84, 0x3d, 0xe1, 0xb7, 0xa0, 0x45,
	0xdb, 0x71, 0xb6, 0xd9, 0x35, 0xae, 0xc6, 0x72, 0xa2, 0x45, 0x5b, 0xfc, 0x01, 0x8c, 0xfb, 0xa4,
	0xa4, 0xe3, 0xdc, 0x45, 0xd7, 0xb8, 0x46, 0x95, 0x94, 0x94, 0x28, 0xda, 0xab, 0x51, 0xc6, 0x2b,
	0x5b, 0xbf, 0xa8, 0x2c, 0xeb, 0x67, 0xf4, 0x14, 0xdf, 0xc0, 0xe2, 0x81, 0xd7, 0x3f, 0x76, 0x5c,
	0x72, 0xdb, 0x50, 0x8e, 0xd7, 0x5d, 0xe3, 0x2e, 0x25, 0xaf, 0xe3, 0xba, 0x87, 0x64, 0x92, 0xb1,
	0x07, 0xf3, 0xb0, 0x88, 0xca, 0xc2, 0x9e, 0x5f, 0xee, 0x4b, 0x8b, 0x98, 0x95, 0x05, 0x19, 0x04,
	0x7c, 0x07, 0x56, 0x24, 0x36, 0x99, 0x64, 0xbf, 0x13, 0x49, 0x73, 0xdb, 0x54, 0x9b, 0xbe, 0xe9,
	0x1a, 0x77, 0xc5, 0x44, 0x9c, 0x9c, 0x39, 0xb9, 0x36, 0xf9, 0x9f, 0x61, 0x75, 0x75, 0xa3, 0xc0,
	0xde, 0x8b, 0x50, 0x56, 0xc1, 0x95, 0x38, 0xc6, 0xf2, 0x13, 0xcc, 0x7b, 0x2e, 0x59, 0x46, 0xfb,
	0xc8, 0x87, 0x4a, 0x85, 0x32, 0xae, 0x54, 0x29, 0x42, 0xce, 0x9e, 0x1b, 0x58, 0x44, 0xe2, 0x2b,
	0xdf, 0xa7, 0x2c, 0x57, 0x01, 0xe9, 0xc3, 0x81, 0x4c, 0xc4, 0x8f, 0x0a, 0x92, 0x49, 0x0e, 0xdd,
	0xc3, 0xc9, 0x99, 0x1d, 0x4f, 0xce, 0xec, 0xd0, 0x3a, 0xe8, 0xd8, 0x3a, 0xe8, 0x6f, 0xeb, 0xa0,
	0x3f, 0xff, 0x9c, 0xd9, 0xaf, 0x79, 0xc9, 0x73, 0xfa, 0x94, 0x9a, 0xea, 0x45, 0x7c, 0xf9, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x18, 0x9e, 0x06, 0x4f, 0x02, 0x00, 0x00,
}
