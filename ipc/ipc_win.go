//go:build windows
// +build windows

package ipc

import (
	"errors"
	"net"
	"strconv"
	"time"

	winio "github.com/Microsoft/go-winio"
)

func namedPipe(id int) (net.Conn, error) {
	sock, err := winio.DialPipe(`\\.\pipe\discord-ipc-`+strconv.Itoa(id), time.Second*2)
	if err != nil {
		if id < 10 {
			return namedPipe(id + 1)
		}
		return nil, errors.New("Failed to connect to windows named pipe")
	}
	return sock, nil

}

func OpenSocket() error {
	var (
		sock net.Conn
		err  error
	)

	sock, err = namedPipe(0)

	if err != nil {
		return err
	}

	socket = sock
	return nil
}