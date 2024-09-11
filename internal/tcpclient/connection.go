package tcpclient

import (
	"fmt"
	"net"
)

/**
* File responsible connecting to the server and initializing the CLI application.
**/

type TcpClient struct {
	addr        string
	conn        net.Conn
	accessToken string
}

func NewTcpClient(addr string) *TcpClient {
	return &TcpClient{
		addr: addr,
	}
}

func (t *TcpClient) InitCLI() {

	// 1. print intro
	t.printIntro()

	// 2. connects to tcp server
	t.connect()

	// 3. read-write to the server in order communicate
	t.communicateWithServer()

	defer t.conn.Close()
}

/**
* Connect to the server and set it for instance access.
**/
func (t *TcpClient) connect() {

	conn, err := net.Dial("tcp", t.addr)

	if err != nil {
		fmt.Printf("Error when attempting to dial to tcp connection at %s, error: %s", t.addr, err)
		return
	}

	// make sure connection is accessible in the entire struct instance
	t.conn = conn
}
