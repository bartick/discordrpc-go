//go:build !windows
// +build !windows

package ipc

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"
)

func GetIpcPath() string {
	variablesnames := []string{
		"XDG_RUNTIME_DIR",
		"TMPDIR",
		"TMP",
		"TEMP",
	}

	for _, variablename := range variablesnames {

		if path, exists := os.LookupEnv(variablename); exists {
			return path
		}
	}

	return "/tmp"
}

func ipcPathExists(id int, path string) (string, error) {
	path = path + "/discord-ipc-" + strconv.Itoa(id)
	_, err := os.Stat(path)
	if err == nil {
		return path, nil
	}
	if os.IsNotExist(err) {
		return "", errors.New("IPC path does not exist")
	}
	if id < 10 {
		return ipcPathExists(id+1, path)
	}
	return "", err
}

func OpenSocket() error {
	var (
		sock net.Conn
		err  error
	)
	if path, err := ipcPathExists(0, GetIpcPath()); err == nil {
		sock, err = net.DialTimeout("unix", path, time.Second*2)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	if err != nil {
		return err
	}

	socket = sock
	return nil
}
