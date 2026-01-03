package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"goplumber/internal/pipewire"
)

const (
	pipewireVersion = 4
)

func main() {
	socketPath := "/run/user/1000/pipewire-0-manager"
	connDial, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer connDial.Close()

	conn := pipewire.Connection{
		Conn: connDial,
	}

	_ = pipewire.NewHello(pipewireVersion, nil)

	hello := pipewire.NewHello(pipewireVersion, nil)
	_, _ = hello.Write(&conn)

	updateProperties := pipewire.NewUpdateProperties(nil)
	_, _ = updateProperties.Write(&conn)

	getRegistry := pipewire.NewGetRegistry(nil)
	_, _ = getRegistry.Write(&conn)

	sync := pipewire.NewSync(nil)
	_, _ = sync.Write(&conn)

	for {
		var id uint32
		binary.Read(&conn, binary.NativeEndian, &id)

		var opandsize uint32
		binary.Read(&conn, binary.NativeEndian, &opandsize)

		var seq uint32
		binary.Read(&conn, binary.NativeEndian, &seq)

		var fds uint32
		binary.Read(&conn, binary.NativeEndian, &fds)
		_ = fds

		op := opandsize >> 24
		sizeMessage := opandsize & 0xFFF

		fmt.Printf("seq: %d, id: %d, op: %d, size: %d", seq, id, op, sizeMessage)

		// Registry Event Global add
		if id == 1 && op == 0 {
			var size uint32
			binary.Read(&conn, binary.NativeEndian, &size)
			var podType uint32
			binary.Read(&conn, binary.NativeEndian, &podType)

			if podType != uint32(pipewire.PODTypeStruct) {
				trash := make([]byte, size)
				binary.Read(&conn, binary.NativeEndian, &trash)
			} else {
				var idSize uint32
				binary.Read(&conn, binary.NativeEndian, &idSize)
				_ = idSize
				var idType uint32
				binary.Read(&conn, binary.NativeEndian, &idType)
				_ = idType
				var idValue uint32
				binary.Read(&conn, binary.NativeEndian, &idValue)
				fmt.Printf(", id: %d\n", idValue)
				trash := make([]byte, sizeMessage-5*4)
				binary.Read(&conn, binary.NativeEndian, &trash)
			}
		} else {
			fmt.Println("")
		}

		if id == 1 && op == 1 {
		}

		if id != 1 || op != 0 {
			trash := make([]byte, sizeMessage)
			_ = binary.Read(&conn, binary.NativeEndian, &trash)
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
