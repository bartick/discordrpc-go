package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"

	code "github.com/bartick/discordrpc-go/opcode"
)

func CloseSocket() error {
	if socket != nil {
		err := socket.Close()
		if err != nil {
			return err
		}

		socket = nil
	}
	return nil
}

// Read the socket response
func Read() string {
	buf := make([]byte, 512)
	payloadlength, err := socket.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buffer := new(bytes.Buffer)
	for i := 8; i < payloadlength; i++ {
		buffer.WriteByte(buf[i])
	}

	return buffer.String()
}

// Send opcode and payload to the unix socket
func Send(opcode code.Gateway, payload string) string {
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, int32(opcode))
	if err != nil {
		fmt.Println(err)
	}

	err = binary.Write(buffer, binary.LittleEndian, int32(len(payload)))
	if err != nil {
		fmt.Println(err)
	}

	buffer.Write([]byte(payload))
	_, err = socket.Write(buffer.Bytes())
	if err != nil {
		fmt.Println(err)
	}

	return Read()
}
