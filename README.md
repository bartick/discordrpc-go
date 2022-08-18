# discordrpc-go
[![Go Reference](https://pkg.go.dev/badge/github.com/bartick/discordrpc-go.svg)](https://pkg.go.dev/github.com/bartick/discordrpc-go)

## Usage

You need to create a client.     
If you don't know how to get your client Id check [this](./GET_APP_ID.md)
```go
conn := rpc.RPC{
		ClientID: "APP_ID",
	}
```
Now you need to login to your discord client. Make sure your discord client is open and not your browser discord.
```go
_, err = conn.Login()  // Returns the login command with data with error (if any)
if err != nil {
  panic(err)
}
  ```
Now you can set an activity. Make sure that your program runs in background or else your status won't persist
```go
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
})      // returns the string returned by discord on successful connection and error (if any)
if err != nil {
  panic(err)
}
```
You can also clear your activity like this
```go
_, err = conn.ClearActivity()
if err != nil {
  panic(err)
}
```
You can also logout of your connection to discord client like this
```go
conn.Logout()
```

This library currently supports RPC status change for windows, linux and mac.            
It currently dosen't support/have features to handel events returned by discord. 
