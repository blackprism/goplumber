package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"goplumber/pipewire"
)

const (
	pipewireVersion     = 3
	pipewireMethodHello = 1
)

func main() {
	socketPath := "/run/user/1000/pipewire-0-manager"
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	// payload := pcore.Hello{
	// 	Version: pipewireVersion,
	// }

	// m := message.Message{
	// 	Payload: payload,
	// }
	// _ = m

	// _, errWrite := m.Write(conn, 0)

	// if errWrite != nil {
	// 	fmt.Println("Bouh 1", errWrite)
	// 	os.Exit(1)
	// }

	// payload2 := pclient.UpdateProperties{
	// 	Items: [][2]string{},
	// }

	// m2 := message.Message{
	// 	Payload: payload2,
	// }

	// _, errWrite2 := m2.Write(conn, 1)

	// if errWrite2 != nil {
	// 	fmt.Println("Bouh 1", errWrite2)
	// 	os.Exit(1)
	// }

	// strace -e trace=writev -xx pw-dump 2>&1 | head -100

	// payloadR := pcore.GetRegistry{
	// 	Version: pipewireVersion,
	// 	NewID:   1,
	// }

	// m3 := message.Message{
	// 	Payload: payloadR,
	// }
	// _ = m3

	// _, errWrite3 := m3.Write(conn, 2)

	// if errWrite3 != nil {
	// 	fmt.Println("Bouh 3", errWrite3)
	// 	os.Exit(1)
	// }

	seq := 0

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint32(0))
	binary.Write(buf, binary.LittleEndian, uint32(0x01000018)) // 24
	binary.Write(buf, binary.LittleEndian, uint32(seq))
	binary.Write(buf, binary.LittleEndian, uint32(0))  // fd
	binary.Write(buf, binary.LittleEndian, uint32(16)) // size
	binary.Write(buf, binary.LittleEndian, uint32(14)) // struct

	// int
	binary.Write(buf, binary.LittleEndian, uint32(4)) // taille
	binary.Write(buf, binary.LittleEndian, uint32(4)) // type int
	binary.Write(buf, binary.LittleEndian, uint32(4)) // value
	binary.Write(buf, binary.LittleEndian, uint32(0)) // padding
	conn.Write(buf.Bytes())

	// update properties
	if true {
		seq += 1
		buf1 := new(bytes.Buffer)
		binary.Write(buf1, binary.LittleEndian, uint32(1))
		binary.Write(buf1, binary.LittleEndian, uint32(0x02000050)) // 20hex = 32 50hex = 80
		binary.Write(buf1, binary.LittleEndian, uint32(seq))
		binary.Write(buf1, binary.LittleEndian, uint32(0))    // fd
		binary.Write(buf1, binary.LittleEndian, uint32(18*4)) // 6 ou 18*4 taille
		binary.Write(buf1, binary.LittleEndian, uint32(14))   // struct
		binary.Write(buf1, binary.LittleEndian, uint32(16*4)) // 4 ou 16*4 size struct
		binary.Write(buf1, binary.LittleEndian, uint32(14))   // type struct

		// 4 x uint 32
		binary.Write(buf1, binary.LittleEndian, uint32(4)) // size int
		binary.Write(buf1, binary.LittleEndian, uint32(4)) // type int
		binary.Write(buf1, binary.LittleEndian, uint32(1)) // n_items
		binary.Write(buf1, binary.LittleEndian, uint32(0))

		// key remote.intention (8 x uint32)
		binary.Write(buf1, binary.LittleEndian, uint32(17))                       // size string
		binary.Write(buf1, binary.LittleEndian, uint32(8))                        // type string
		binary.Write(buf1, binary.LittleEndian, toFixedBytes("remote.intention")) // 16 byte
		binary.Write(buf1, binary.LittleEndian, uint32(0))                        // null end string
		binary.Write(buf1, binary.LittleEndian, uint32(0))                        // padding

		// value manager (4 x uint32)
		binary.Write(buf1, binary.LittleEndian, uint32(8))               // size string
		binary.Write(buf1, binary.LittleEndian, uint32(8))               // type string
		binary.Write(buf1, binary.LittleEndian, toFixedBytes("manager")) // 7 byte
		binary.Write(buf1, binary.LittleEndian, uint8(0))                // 1 byte

		conn.Write(buf1.Bytes())
	}

	// register
	if true {
		seq += 1
		buf2 := new(bytes.Buffer)
		_ = binary.Write(buf2, binary.NativeEndian, uint32(0))
		_ = binary.Write(buf2, binary.NativeEndian, uint32(0x05000028)) // 40
		_ = binary.Write(buf2, binary.NativeEndian, uint32(seq))
		_ = binary.Write(buf2, binary.NativeEndian, uint32(0))  // fd
		_ = binary.Write(buf2, binary.NativeEndian, uint32(32)) // size struct
		_ = binary.Write(buf2, binary.NativeEndian, uint32(14))

		_ = binary.Write(buf2, binary.NativeEndian, uint32(4)) // size int
		_ = binary.Write(buf2, binary.NativeEndian, uint32(4)) // type int
		_ = binary.Write(buf2, binary.NativeEndian, uint32(3)) // value
		_ = binary.Write(buf2, binary.NativeEndian, uint32(0))

		_ = binary.Write(buf2, binary.NativeEndian, uint32(4)) // size int
		_ = binary.Write(buf2, binary.NativeEndian, uint32(4)) // type int
		_ = binary.Write(buf2, binary.NativeEndian, uint32(1)) // value
		_ = binary.Write(buf2, binary.NativeEndian, uint32(0))
		conn.Write(buf2.Bytes())
	}

	// ============ SYNC ============
	seq += 1
	buf3 := new(bytes.Buffer)
	binary.Write(buf3, binary.LittleEndian, uint32(0))
	binary.Write(buf3, binary.LittleEndian, uint32(0x02000028))
	binary.Write(buf3, binary.LittleEndian, uint32(seq))
	binary.Write(buf3, binary.LittleEndian, uint32(0))
	binary.Write(buf3, binary.LittleEndian, uint32(32))
	binary.Write(buf3, binary.LittleEndian, uint32(14))
	binary.Write(buf3, binary.LittleEndian, uint32(4))
	binary.Write(buf3, binary.LittleEndian, uint32(4))
	binary.Write(buf3, binary.LittleEndian, uint32(0))
	binary.Write(buf3, binary.LittleEndian, uint32(0))
	binary.Write(buf3, binary.LittleEndian, uint32(4))
	binary.Write(buf3, binary.LittleEndian, uint32(4))
	binary.Write(buf3, binary.LittleEndian, uint32(42))
	binary.Write(buf3, binary.LittleEndian, uint32(0))
	conn.Write(buf3.Bytes())

	// var data []byte
	for {
		var id uint32
		binary.Read(conn, binary.NativeEndian, &id)

		var opandsize uint32
		binary.Read(conn, binary.NativeEndian, &opandsize)

		var seq uint32
		binary.Read(conn, binary.NativeEndian, &seq)

		var fds uint32
		binary.Read(conn, binary.NativeEndian, &fds)
		_ = fds

		op := opandsize >> 24
		sizeMessage := opandsize & 0xFFF

		fmt.Printf("seq: %d, id: %d, op: %d, size: %d", seq, id, op, sizeMessage)

		// Registry Event Global add
		if id == 1 && op == 0 {
			var size uint32
			binary.Read(conn, binary.NativeEndian, &size)
			var podType uint32
			binary.Read(conn, binary.NativeEndian, &podType)

			if podType != uint32(pipewire.PODTypeStruct) {
				trash := make([]byte, size)
				binary.Read(conn, binary.NativeEndian, &trash)
			} else {
				var idSize uint32
				binary.Read(conn, binary.NativeEndian, &idSize)
				_ = idSize
				var idType uint32
				binary.Read(conn, binary.NativeEndian, &idType)
				_ = idType
				var idValue uint32
				binary.Read(conn, binary.NativeEndian, &idValue)
				fmt.Printf(", id: %d\n", idValue)
				trash := make([]byte, sizeMessage-5*4)
				binary.Read(conn, binary.NativeEndian, &trash)
			}
		} else {
			fmt.Println("")
		}

		if id == 1 && op == 1 {
		}

		if id != 1 || op != 0 {
			trash := make([]byte, sizeMessage)
			_ = binary.Read(conn, binary.NativeEndian, &trash)
		}

		// fmt.Printf("head:\n%s", hex.Dump(buf))

		if err != nil {
			fmt.Printf("error: %s", err.Error())
			os.Exit(1)
		}

		// n, errConn := conn.Read(buf)

		// fmt.Println(n)

		// if errConn != nil {
		// 	fmt.Println("errConn", errConn)
		// 	os.Exit(1)
		// }

		// data = append(data, buf[:n]...)

		// fmt.Println(string(data))
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("")

		// // Prevent reading too much
		// if len(data) > 65536 {
		// 	break
		// }
	}
}

func toFixedBytes(s string) []byte {
	b := make([]byte, len(s))
	copy(b, s)
	return b
}
