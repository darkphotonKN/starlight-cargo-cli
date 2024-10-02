package types

import "net"

type ClientConnector interface {
	AuthenticateWithServer() error
	Connect() error
	Conn() *net.Conn
	Token() string
}
