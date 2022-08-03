package main

import (
	"fmt"
	"time"

	rpc "github.com/bartick/discordrpc-go"
)

func main() {
	var err error
	var answer string
	conn := rpc.RPC{
		ClientID: "<PUT YOUR CLIEND ID FROM DISCORD DEVELOPERS PORTAL>",
	}
	_, err = conn.Login()
	if err != nil {
		panic(err)
	}

	_, err = conn.SetActivity(rpc.Activity{
		State:   "Playing",
		Details: "Playing a game",
		Assets: &rpc.ActivityAssets{
			LargeImage: "foo",
			LargeText:  "This is the image",
		},
		Timestamps: &rpc.ActivityTimestamps{
			Start: time.Now().Unix(),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Scanf("%s", &answer)
	_, err = conn.ClearActivity()
	if err != nil {
		panic(err)
	}
	fmt.Scanf("%s", &answer)
	conn.Logout()
}
