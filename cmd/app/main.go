package main

import (
	"github.com/darkphotonKN/starlight-cargo-cli/internal/communication"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/tcpclient"
)

func main() {
	console := console.NewConsole()
	tcpClient := tcpclient.NewTcpClient(":3600", console)
	communication := communication.NewCommunicationService(tcpClient, console)

	console.PrintIntro()
	tcpClient.Connect()
	communication.CommunicateWithServer()

}
