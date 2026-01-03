package pipewire

import (
	"bytes"
	"encoding/binary"
)

const (
	updatePropertiesTargetID = 1
	updatePropertiesOpCode   = 2
)

type updateProperties struct {
	writer Writer
}

func NewUpdateProperties(writer Writer) updateProperties {
	if writer == nil {
		writer = Messager{}
	}

	return updateProperties{
		writer: writer,
	}
}

func (up updateProperties) Write(writer PipewireWriter) (int, error) {
	return up.writer.Write(writer, updatePropertiesTargetID, updatePropertiesOpCode, up.toPayload())
}

func (up updateProperties) toPayload() bytes.Buffer {
	var buf bytes.Buffer
	buf.Grow(18 * 4)

	_ = binary.Write(&buf, binary.NativeEndian, uint32(18*4)) // size struct
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeStruct))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(16*4)) // size struct
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeStruct))

	// Pod Int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(1)) // n_items
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	// key
	_ = binary.Write(&buf, binary.NativeEndian, uint32(17)) // size string
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeString))
	_ = binary.Write(&buf, binary.NativeEndian, toFixedBytes("remote.intention"))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // null end string
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	// value
	_ = binary.Write(&buf, binary.NativeEndian, uint32(8)) // size string
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeString))
	_ = binary.Write(&buf, binary.NativeEndian, toFixedBytes("manager"))
	_ = binary.Write(&buf, binary.NativeEndian, uint8(0)) // null end string

	return buf
}

func toFixedBytes(s string) []byte {
	b := make([]byte, len(s))
	copy(b, s)
	return b
}
