package client

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bartick/discordrpc-go/ipc"
	code "github.com/bartick/discordrpc-go/opcode"
)

type RPC struct {
	ClientID string
	RPCInterface
}

var logged bool = false

func (conn *RPC) Login() (string, error) {

	if logged {
		return "", nil
	}

	payload, err := json.Marshal(Handshake{RPCVersion1, conn.ClientID})
	if err != nil {
		return "", err
	}

	err = ipc.OpenSocket()
	if err != nil {
		return "", err
	}

	res := ipc.Send(code.DISPATCH, string(payload))

	logged = true

	return res, nil
}

func (conn *RPC) Logout() error {

	if !logged {
		return nil
	}

	logged = false

	err := ipc.CloseSocket()
	if err != nil {
		return err
	}

	return nil
}

func (conn *RPC) SetActivity(activity Activity) (string, error) {
	if !logged {
		return "", errors.New("not logged in to discord client")
	}

	payload, err := json.Marshal(Frame{
		Cmd: SET_ACTIVITY,
		Args: Args{
			Pid:      os.Getpid(),
			Activity: &activity,
		},
		Nonce: getNonce(),
	})

	if err != nil {
		return "", err
	}

	res := ipc.Send(code.HEARTBEAT, string(payload))
	return res, nil
}

func (conn *RPC) ClearActivity() (string, error) {

	payload, err := json.Marshal(Frame{
		Cmd: SET_ACTIVITY,
		Args: Args{
			Pid: os.Getpid(),
		},
		Nonce: getNonce(),
	})

	if err != nil {
		return "", err
	}

	res := ipc.Send(code.HEARTBEAT, string(payload))
	return res, nil
}

func getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
