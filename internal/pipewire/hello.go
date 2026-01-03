package pipewire

import (
	"bytes"
	"encoding/binary"
)

const (
	helloTargetID = 0
	helloOpCode   = 1
)

type hello struct {
	writer  Writer
	version uint32
}

func NewHello(version uint32, writer Writer) hello {
	if writer == nil {
		writer = Messager{}
	}

	return hello{
		writer:  writer,
		version: version,
	}
}

func (h hello) Write(writer PipewireWriter) (int, error) {
	return h.writer.Write(writer, helloTargetID, helloOpCode, h.toPayload())
}

func (h hello) toPayload() bytes.Buffer {
	var buf bytes.Buffer
	buf.Grow(16)

	_ = binary.Write(&buf, binary.NativeEndian, uint32(16)) // size struct
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeStruct))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, h.version)
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0))

	return buf
}
