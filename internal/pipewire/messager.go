package pipewire

import (
	"bytes"
	"encoding/binary"
)

const (
	numberFD = 0
)

type Writer interface {
	Write(writer PipewireWriter, targetID uint32, opCode uint32, payload bytes.Buffer) (int, error)
}

type Messager struct{}

func (m Messager) buildHeader(targetID uint32, opCode uint32, sequence uint32, payloadSize int) bytes.Buffer {
	var buf bytes.Buffer
	buf.Grow(16)

	_ = binary.Write(&buf, binary.NativeEndian, uint32(targetID))
	_ = binary.Write(&buf, binary.NativeEndian, uint32(payloadSize)|uint32(opCode)<<24)
	_ = binary.Write(&buf, binary.NativeEndian, sequence)
	_ = binary.Write(&buf, binary.NativeEndian, uint32(numberFD))

	return buf
}

func (m Messager) Write(writer PipewireWriter, targetID uint32, opCode uint32, payload bytes.Buffer) (int, error) {
	header := m.buildHeader(targetID, opCode, writer.Seq(), payload.Len())

	nHeader, errHeader := writer.Write(header.Bytes())

	if errHeader != nil {
		return nHeader, errHeader
	}

	nPayload, errPayload := writer.Write(payload.Bytes())

	if errPayload != nil {
		return nHeader + nPayload, errPayload
	}

	return nHeader + nPayload, nil
}
