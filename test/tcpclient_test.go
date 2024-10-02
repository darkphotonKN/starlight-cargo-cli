package test

import (
	"testing"

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
