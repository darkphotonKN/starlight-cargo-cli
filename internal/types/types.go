package types

type ClientConnector interface {
	AuthenticateWithServer() error
	Connect()
}
