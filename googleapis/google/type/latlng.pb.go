// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/latlng.proto

package google_type

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (m *LatLng) Reset()      { *m = LatLng{} }
func (*LatLng) ProtoMessage() {}
func (*LatLng) Descriptor() ([]byte, []int) {
	return fileDescriptor_a087c9a103c281a9, []int{0}
}
func (m *LatLng) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatLng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatLng.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatLng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatLng.Merge(m, src)
}
func (m *LatLng) XXX_Size() int {
	return m.Size()
}
func (m *LatLng) XXX_DiscardUnknown() {
	xxx_messageInfo_LatLng.DiscardUnknown(m)
}

var xxx_messageInfo_LatLng proto.InternalMessageInfo

func (m *LatLng) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LatLng) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*LatLng)(nil), "google.type.LatLng")
}

func init() { proto.RegisterFile("google/type/latlng.proto", fileDescriptor_a087c9a103c281a9) }

var fileDescriptor_a087c9a103c281a9 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x49, 0x2c, 0xc9, 0xc9, 0x4b, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xe8, 0x81, 0x64, 0x94, 0x9c, 0xb8, 0xd8, 0x7c,
	0x12, 0x4b, 0x7c, 0xf2, 0xd2, 0x85, 0xa4, 0xb8, 0x38, 0x72, 0x12, 0x4b, 0x32, 0x4b, 0x4a, 0x53,
	0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x18, 0x83, 0xe0, 0x7c, 0x21, 0x19, 0x2e, 0xce, 0x9c, 0xfc,
	0xbc, 0x74, 0x88, 0x24, 0x13, 0x58, 0x12, 0x21, 0xe0, 0x14, 0x7b, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d,
	0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e,
	0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x81, 0x8b, 0x3f, 0x39, 0x3f, 0x57, 0x0f, 0xc9, 0x29, 0x4e, 0xdc, 0x10, 0x87, 0x04, 0x80,
	0x1c, 0x19, 0xc0, 0xf8, 0x83, 0x91, 0x71, 0x11, 0x13, 0xb3, 0x7b, 0x48, 0x40, 0x12, 0x1b, 0xd8,
	0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x88, 0x65, 0xda, 0xd2, 0x00, 0x00, 0x00,
}

func (this *LatLng) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LatLng)
	if !ok {
		that2, ok := that.(LatLng)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Latitude != that1.Latitude {
		return false
	}
	if this.Longitude != that1.Longitude {
		return false
	}
	return true
}
func (this *LatLng) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&google_type.LatLng{")
	s = append(s, "Latitude: "+fmt.Sprintf("%#v", this.Latitude)+",\n")
	s = append(s, "Longitude: "+fmt.Sprintf("%#v", this.Longitude)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLatlng(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LatLng) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatLng) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatLng) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Longitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i--
		dAtA[i] = 0x11
	}
	if m.Latitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintLatlng(dAtA []byte, offset int, v uint64) int {
	offset -= sovLatlng(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LatLng) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Latitude != 0 {
		n += 9
	}
	if m.Longitude != 0 {
		n += 9
	}
	return n
}

func sovLatlng(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLatlng(x uint64) (n int) {
	return sovLatlng(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LatLng) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LatLng{`,
		`Latitude:` + fmt.Sprintf("%v", this.Latitude) + `,`,
		`Longitude:` + fmt.Sprintf("%v", this.Longitude) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLatlng(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LatLng) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLatlng
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
			return fmt.Errorf("proto: LatLng: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatLng: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipLatlng(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLatlng
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLatlng
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
func skipLatlng(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLatlng
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
					return 0, ErrIntOverflowLatlng
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
					return 0, ErrIntOverflowLatlng
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
				return 0, ErrInvalidLengthLatlng
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLatlng
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLatlng
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
				next, err := skipLatlng(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLatlng
				}
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
	ErrInvalidLengthLatlng = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLatlng   = fmt.Errorf("proto: integer overflow")
)
