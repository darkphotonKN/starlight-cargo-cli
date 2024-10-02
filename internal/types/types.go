package types

import "net"

type ClientConnector interface {
	AuthenticateWithServer() error
	Connect()
	Conn() *net.Conn
	Token() string
}
