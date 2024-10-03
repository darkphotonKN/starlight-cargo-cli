package types

import "net"

type ClientConnector interface {
	// requests the server authentication and provides an error if this fails
	AuthenticateWithServer() error

	// attempts to connect to a server and returns an error if it fails
	Connect() error

	// provides the current established (or nil) connection
	Conn() *net.Conn

	// provides the current token from the authentication
	Token() string
}
