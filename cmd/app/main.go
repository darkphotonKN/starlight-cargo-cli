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
	tcpClientInstance, ok := tcpClient.(*tcpclient.TcpClient)
	if !ok {
		log.Panic("Tcp client instance errored before starting.")
	}
	communication := communication.NewCommunicationService(tcpClientInstance, console)

	console.PrintIntro()
	tcpClient.Connect()
	communication.CommunicateWithServer()

}
