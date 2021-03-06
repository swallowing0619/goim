// Code generated by protoc-gen-gogo.
// source: router.proto
// DO NOT EDIT!

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		router.proto

	It has these top-level messages:
		NoArg
		NoReply
		ConnArg
		ConnReply
		DisconnArg
		DisconnReply
		GetArg
		GetReply
		GetAllReply
		MGetArg
		MGetReply
		GetSeqCountArg
		GetSeqCountReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

// discarding unused import gogoproto "gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type NoArg struct {
}

func (m *NoArg) Reset()         { *m = NoArg{} }
func (m *NoArg) String() string { return proto1.CompactTextString(m) }
func (*NoArg) ProtoMessage()    {}

type NoReply struct {
}

func (m *NoReply) Reset()         { *m = NoReply{} }
func (m *NoReply) String() string { return proto1.CompactTextString(m) }
func (*NoReply) ProtoMessage()    {}

type ConnArg struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Server int32 `protobuf:"varint,2,opt,name=server,proto3" json:"server,omitempty"`
}

func (m *ConnArg) Reset()         { *m = ConnArg{} }
func (m *ConnArg) String() string { return proto1.CompactTextString(m) }
func (*ConnArg) ProtoMessage()    {}

type ConnReply struct {
	Seq int32 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (m *ConnReply) Reset()         { *m = ConnReply{} }
func (m *ConnReply) String() string { return proto1.CompactTextString(m) }
func (*ConnReply) ProtoMessage()    {}

type DisconnArg struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Seq    int32 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (m *DisconnArg) Reset()         { *m = DisconnArg{} }
func (m *DisconnArg) String() string { return proto1.CompactTextString(m) }
func (*DisconnArg) ProtoMessage()    {}

type DisconnReply struct {
	Has bool `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
}

func (m *DisconnReply) Reset()         { *m = DisconnReply{} }
func (m *DisconnReply) String() string { return proto1.CompactTextString(m) }
func (*DisconnReply) ProtoMessage()    {}

type GetArg struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (m *GetArg) Reset()         { *m = GetArg{} }
func (m *GetArg) String() string { return proto1.CompactTextString(m) }
func (*GetArg) ProtoMessage()    {}

type GetReply struct {
	Seqs    []int32 `protobuf:"varint,1,rep,name=seqs" json:"seqs,omitempty"`
	Servers []int32 `protobuf:"varint,2,rep,name=servers" json:"servers,omitempty"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto1.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}

type GetAllReply struct {
	UserIds  []int64     `protobuf:"varint,1,rep,name=userIds" json:"userIds,omitempty"`
	Sessions []*GetReply `protobuf:"bytes,2,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *GetAllReply) Reset()         { *m = GetAllReply{} }
func (m *GetAllReply) String() string { return proto1.CompactTextString(m) }
func (*GetAllReply) ProtoMessage()    {}

func (m *GetAllReply) GetSessions() []*GetReply {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type MGetArg struct {
	UserIds []int64 `protobuf:"varint,1,rep,name=userIds" json:"userIds,omitempty"`
}

func (m *MGetArg) Reset()         { *m = MGetArg{} }
func (m *MGetArg) String() string { return proto1.CompactTextString(m) }
func (*MGetArg) ProtoMessage()    {}

type MGetReply struct {
	UserIds  []int64     `protobuf:"varint,1,rep,name=userIds" json:"userIds,omitempty"`
	Sessions []*GetReply `protobuf:"bytes,2,rep,name=sessions" json:"sessions,omitempty"`
}

func (m *MGetReply) Reset()         { *m = MGetReply{} }
func (m *MGetReply) String() string { return proto1.CompactTextString(m) }
func (*MGetReply) ProtoMessage()    {}

func (m *MGetReply) GetSessions() []*GetReply {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type GetSeqCountArg struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (m *GetSeqCountArg) Reset()         { *m = GetSeqCountArg{} }
func (m *GetSeqCountArg) String() string { return proto1.CompactTextString(m) }
func (*GetSeqCountArg) ProtoMessage()    {}

type GetSeqCountReply struct {
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *GetSeqCountReply) Reset()         { *m = GetSeqCountReply{} }
func (m *GetSeqCountReply) String() string { return proto1.CompactTextString(m) }
func (*GetSeqCountReply) ProtoMessage()    {}

func init() {
}
func (m *NoArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		switch fieldNum {
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *NoReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		switch fieldNum {
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ConnArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UserId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Server |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ConnReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seq |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DisconnArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UserId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seq", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seq |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DisconnReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Has", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Has = bool(v != 0)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *GetArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UserId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *GetReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seqs", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Seqs = append(m.Seqs, v)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Servers", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Servers = append(m.Servers, v)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *GetAllReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UserIds = append(m.UserIds, v)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sessions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sessions = append(m.Sessions, &GetReply{})
			if err := m.Sessions[len(m.Sessions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *MGetArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UserIds = append(m.UserIds, v)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *MGetReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UserIds = append(m.UserIds, v)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sessions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sessions = append(m.Sessions, &GetReply{})
			if err := m.Sessions[len(m.Sessions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *GetSeqCountArg) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UserId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *GetSeqCountReply) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipRouter(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipRouter(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRouter(data[start:])
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
func (m *NoArg) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *NoReply) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ConnArg) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovRouter(uint64(m.UserId))
	}
	if m.Server != 0 {
		n += 1 + sovRouter(uint64(m.Server))
	}
	return n
}

func (m *ConnReply) Size() (n int) {
	var l int
	_ = l
	if m.Seq != 0 {
		n += 1 + sovRouter(uint64(m.Seq))
	}
	return n
}

func (m *DisconnArg) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovRouter(uint64(m.UserId))
	}
	if m.Seq != 0 {
		n += 1 + sovRouter(uint64(m.Seq))
	}
	return n
}

func (m *DisconnReply) Size() (n int) {
	var l int
	_ = l
	if m.Has {
		n += 2
	}
	return n
}

func (m *GetArg) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovRouter(uint64(m.UserId))
	}
	return n
}

func (m *GetReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Seqs) > 0 {
		for _, e := range m.Seqs {
			n += 1 + sovRouter(uint64(e))
		}
	}
	if len(m.Servers) > 0 {
		for _, e := range m.Servers {
			n += 1 + sovRouter(uint64(e))
		}
	}
	return n
}

func (m *GetAllReply) Size() (n int) {
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, e := range m.UserIds {
			n += 1 + sovRouter(uint64(e))
		}
	}
	if len(m.Sessions) > 0 {
		for _, e := range m.Sessions {
			l = e.Size()
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	return n
}

func (m *MGetArg) Size() (n int) {
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, e := range m.UserIds {
			n += 1 + sovRouter(uint64(e))
		}
	}
	return n
}

func (m *MGetReply) Size() (n int) {
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, e := range m.UserIds {
			n += 1 + sovRouter(uint64(e))
		}
	}
	if len(m.Sessions) > 0 {
		for _, e := range m.Sessions {
			l = e.Size()
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	return n
}

func (m *GetSeqCountArg) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovRouter(uint64(m.UserId))
	}
	return n
}

func (m *GetSeqCountReply) Size() (n int) {
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovRouter(uint64(m.Count))
	}
	return n
}

func sovRouter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRouter(x uint64) (n int) {
	return sovRouter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NoArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NoArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *NoReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *NoReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ConnArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConnArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.UserId))
	}
	if m.Server != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintRouter(data, i, uint64(m.Server))
	}
	return i, nil
}

func (m *ConnReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConnReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seq != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.Seq))
	}
	return i, nil
}

func (m *DisconnArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DisconnArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.UserId))
	}
	if m.Seq != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintRouter(data, i, uint64(m.Seq))
	}
	return i, nil
}

func (m *DisconnReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DisconnReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Has {
		data[i] = 0x8
		i++
		if m.Has {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *GetArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.UserId))
	}
	return i, nil
}

func (m *GetReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Seqs) > 0 {
		for _, num := range m.Seqs {
			data[i] = 0x8
			i++
			i = encodeVarintRouter(data, i, uint64(num))
		}
	}
	if len(m.Servers) > 0 {
		for _, num := range m.Servers {
			data[i] = 0x10
			i++
			i = encodeVarintRouter(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *GetAllReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetAllReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, num := range m.UserIds {
			data[i] = 0x8
			i++
			i = encodeVarintRouter(data, i, uint64(num))
		}
	}
	if len(m.Sessions) > 0 {
		for _, msg := range m.Sessions {
			data[i] = 0x12
			i++
			i = encodeVarintRouter(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MGetArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MGetArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, num := range m.UserIds {
			data[i] = 0x8
			i++
			i = encodeVarintRouter(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *MGetReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *MGetReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserIds) > 0 {
		for _, num := range m.UserIds {
			data[i] = 0x8
			i++
			i = encodeVarintRouter(data, i, uint64(num))
		}
	}
	if len(m.Sessions) > 0 {
		for _, msg := range m.Sessions {
			data[i] = 0x12
			i++
			i = encodeVarintRouter(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GetSeqCountArg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetSeqCountArg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.UserId))
	}
	return i, nil
}

func (m *GetSeqCountReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetSeqCountReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintRouter(data, i, uint64(m.Count))
	}
	return i, nil
}

func encodeFixed64Router(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Router(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRouter(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
