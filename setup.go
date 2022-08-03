package client

type RPCInterface interface {
	Login(clientid string) error
	Logout()
	SetActivity(activity Activity) error
	ClearActivity() error
}
