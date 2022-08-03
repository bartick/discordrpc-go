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

// Check for named pipe
func namedPipe(id int) (net.Conn, error) {
	var d = time.Duration(2 * time.Second)
	sock, err := winio.DialPipe(`\\.\pipe\discord-ipc-`+strconv.Itoa(id), &d)
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
