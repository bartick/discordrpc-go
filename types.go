package client

type rpcVersion string

type rpcCommand string

type ActivityType int

type ActivityJoinRequestReply int

type ActivityActionType int

const (
	RPCVersion1 rpcVersion = "1"
)

const (
	PLAYING ActivityType = iota
	STREAMING
	LISTENING
	WATCHING
	CUSTOM
	COMPETING
)

const (
	NO ActivityJoinRequestReply = iota
	YES
	IGNORE
)

const (
	JOIN ActivityActionType = iota
	SPECTATE
)

const (
	DISPATCH                   rpcCommand = "DISPATCH"
	AUTHORIZE                  rpcCommand = "AUTHORIZE"
	AUTHENTICATE               rpcCommand = "AUTHENTICATE"
	GET_GUILD                  rpcCommand = "GET_GUILD"
	GET_GUILDS                 rpcCommand = "GET_GUILDS"
	GET_CHANNEL                rpcCommand = "GET_CHANNEL"
	GET_CHANNELS               rpcCommand = "GET_CHANNELS"
	SUBSCRIBE                  rpcCommand = "SUBSCRIBE"
	UNSUBSCRIBE                rpcCommand = "UNSUBSCRIBE"
	SET_USER_VOICE_SETTINGS    rpcCommand = "SET_USER_VOICE_SETTINGS"
	SELECT_VOICE_CHANNEL       rpcCommand = "SELECT_VOICE_CHANNEL"
	GET_SELECTED_VOICE_CHANNEL rpcCommand = "GET_SELECTED_VOICE_CHANNEL"
	SELECT_TEXT_CHANNEL        rpcCommand = "SELECT_TEXT_CHANNEL"
	GET_VOICE_SETTINGS         rpcCommand = "GET_VOICE_SETTINGS"
	SET_VOICE_SETTINGS         rpcCommand = "SET_VOICE_SETTINGS"
	SET_CERTIFIED_DEVICES      rpcCommand = "SET_CERTIFIED_DEVICES"
	SET_ACTIVITY               rpcCommand = "SET_ACTIVITY"
	SEND_ACTIVITY_JOIN_INVITE  rpcCommand = "SEND_ACTIVITY_JOIN_INVITE"
	CLOSE_ACTIVITY_REQUEST     rpcCommand = "CLOSE_ACTIVITY_REQUEST"
)

type Handshake struct {
	V        rpcVersion `json:"v"`
	ClientId string     `json:"client_id"`
}

type Frame struct {
	Cmd   rpcCommand `json:"cmd"`
	Args  Args       `json:"args"`
	Nonce string     `json:"nonce,omitempty"`
}

type Args struct {
	Pid      int       `json:"pid"`
	Activity *Activity `json:"activity,omitempty"`
}

type User struct {
	Id            int64  `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

type Activity struct {
	State      string              `json:"state,omitempty"`
	Details    string              `json:"details,omitempty"`
	Timestamps *ActivityTimestamps `json:"timestamps,omitempty"`
	Assets     *ActivityAssets     `json:"assets,omitempty"`
	Party      *ActivityParty      `json:"party,omitempty"`
	Secrets    *ActivitySecrets    `json:"secrets,omitempty"`
	Buttons    []*ActivityButton   `json:"buttons,omitempty"`
	Instance   bool                `json:"instance,omitempty"`
}

type ActivityTimestamps struct {
	Start int64 `json:"start"`
	End   int64 `json:"end,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivityParty struct {
	ID   string     `json:"id"`
	Size *PartySize `json:"size"`
}

type PartySize struct {
	Current int32 `json:"current"`
	Max     int32 `json:"max,omitempty"`
}

type ActivitySecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}
