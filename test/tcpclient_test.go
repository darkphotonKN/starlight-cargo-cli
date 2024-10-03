package test

import (
	"testing"
	"time"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/tcpclient"
	"github.com/stretchr/testify/assert"
)

func TestTcpClient_Connect(t *testing.T) {
	// test tcp connection success
	c := console.NewConsole()
	tcpClient := tcpclient.NewTcpClient(":3600", c)

	err := tcpClient.Connect()

	assert.NoError(t, err, "No error expected during connection.")
}

func TestTcpClient_AuthenticateWithServer(t *testing.T) {
	mockConsole := console.NewConsole()
	mockConnection := NewMockConnection()

	// Simulate the server response to authenticate
	go func() {
		time.Sleep(time.Second * 1)
		mockConnection.Write([]byte("AUTHENTICATED: YourAccessToken\n"))
	}()

	// Create a mock client using the MockConnection instead of net.Conn
	client := tcpclient.NewTcpClient("localhost:3600", mockConsole)
}
