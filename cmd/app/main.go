package main

import (
	"log"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/communication"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/tcpclient"
)

func main() {
	console := console.NewConsole()
	tcpClient := tcpclient.NewTcpClient(":3600", console)
	communication := communication.NewCommunicationService(tcpClient, console)

	console.PrintIntro()
	err := tcpClient.Connect()
	if err != nil {
		log.Panic("Could not connect to server.")
	}
	communication.CommunicateWithServer()

}
