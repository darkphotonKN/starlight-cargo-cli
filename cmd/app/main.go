package main

import "github.com/darkphotonKN/starlight-cargo-cli/tcpclient"

func main() {
	tcpConnection := tcpclient.NewTcpClient(":3600")
	tcpConnection.ConnectAndComm()
}
