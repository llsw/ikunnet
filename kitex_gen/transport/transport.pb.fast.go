// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package transport

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Transport) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Transport[number], err)
}

func (x *Transport) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var ov Transport_Addr
	x.Svc = &ov
	ov.Addr, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Transport) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov Transport_Name
	x.Svc = &ov
	ov.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Transport) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Session, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Transport) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *Transport) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Transport) fastWriteField1(buf []byte) (offset int) {
	if x.GetAddr() == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetAddr())
	return offset
}

func (x *Transport) fastWriteField2(buf []byte) (offset int) {
	if x.GetName() == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Transport) fastWriteField3(buf []byte) (offset int) {
	if x.Session == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetSession())
	return offset
}

func (x *Transport) fastWriteField4(buf []byte) (offset int) {
	if len(x.Msg) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 4, x.GetMsg())
	return offset
}

func (x *Transport) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Transport) sizeField1() (n int) {
	if x.GetAddr() == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetAddr())
	return n
}

func (x *Transport) sizeField2() (n int) {
	if x.GetName() == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Transport) sizeField3() (n int) {
	if x.Session == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetSession())
	return n
}

func (x *Transport) sizeField4() (n int) {
	if len(x.Msg) == 0 {
		return n
	}
	n += fastpb.SizeBytes(4, x.GetMsg())
	return n
}

var fieldIDToName_Transport = map[int32]string{
	1: "Addr",
	2: "Name",
	3: "Session",
	4: "Msg",
}
