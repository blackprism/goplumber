package pipewire

import (
	"bytes"
	"encoding/binary"
)

const (
	getRegistryTargetID = 0
	getRegistryOpCode   = 5
	registryVersion     = 3
	newID               = 1
)

type getRegistry struct {
	writer Writer
}

func NewGetRegistry(writer Writer) getRegistry {
	if writer == nil {
		writer = Messager{}
	}

	return getRegistry{
		writer: writer,
	}
}

func (s getRegistry) Write(writer PipewireWriter) (int, error) {
	return s.writer.Write(writer, getRegistryTargetID, getRegistryOpCode, s.toPayload())
}

func (s getRegistry) toPayload() bytes.Buffer {
	var buf bytes.Buffer
	buf.Grow(32)

	_ = binary.Write(&buf, binary.NativeEndian, uint32(32)) // size struct
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeStruct))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(registryVersion))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(newID))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	return buf
}
