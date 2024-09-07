package main

import "github.com/darkphotonKN/starlight-cargo-cli/internal/tcpclient"

func main() {
	tcpConnection := tcpclient.NewTcpClient(":3600")
	tcpConnection.InitCLI()
}
