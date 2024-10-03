package main

import (
	"log"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/communication"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	fileservice "github.com/darkphotonKN/starlight-cargo-cli/internal/file_service"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/tcpclient"
)

func main() {
	fileService := fileservice.NewFileService()
	console := console.NewConsole()
	tcpClient := tcpclient.NewTcpClient(":3600", console)
	communication := communication.NewCommunicationService(tcpClient, fileService, console)

	console.PrintIntro()
	err := tcpClient.Connect()
	if err != nil {
		log.Panic("Could not connect to server.")
	}
	communication.CommunicateWithServer()

}
