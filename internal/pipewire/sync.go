package pipewire

import (
	"bytes"
	"encoding/binary"
)

const (
	syncTargetID = 0
	syncOpCode   = 5
)

type sync struct {
	writer Writer
}

func NewSync(writer Writer) sync {
	if writer == nil {
		writer = Messager{}
	}

	return sync{
		writer: writer,
	}
}

func (s sync) Write(writer PipewireWriter) (int, error) {
	return s.writer.Write(writer, syncTargetID, syncOpCode, s.toPayload(writer))
}

func (s sync) toPayload(writer PipewireWriter) bytes.Buffer {
	var buf bytes.Buffer
	buf.Grow(32)

	_ = binary.Write(&buf, binary.NativeEndian, uint32(32)) // size struct
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeStruct))

	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // id
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	_ = binary.Write(&buf, binary.NativeEndian, uint32(4)) // size int
	_ = binary.Write(&buf, binary.NativeEndian, uint32(PODTypeInt))
	_ = binary.Write(&buf, binary.NativeEndian, writer.Seq())
	_ = binary.Write(&buf, binary.NativeEndian, uint32(0)) // padding

	return buf
}
