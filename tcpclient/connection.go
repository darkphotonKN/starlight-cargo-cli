package tcpclient

import (
	"fmt"
	"net"
	"time"
)

type TcpClient struct {
	addr string
	conn net.Conn
}

func NewTcpClient(addr string) *TcpClient {
	return &TcpClient{
		addr: addr,
	}
}

func (t *TcpClient) ConnectAndComm() {
	t.PrintIntro()

	conn, err := net.Dial("tcp", t.addr)

	if err != nil {
		fmt.Printf("Error when attempting to dial to tcp connection at %s, error: %s", t.addr, err)
		return
	}

	// make sure connection is accessible in the entire struct
	t.conn = conn

	defer t.conn.Close()

	// initialize communication with server
	err = t.CommunicateWithServer()

	if err != nil {
		fmt.Printf("Error with server: %s\n\n", err)
		return
	}

}

func (t *TcpClient) PrintIntro() {

	fmt.Println(`  
 _______ _________ _______  _______  _       _________ _______          _________   _______  _______  _______  _______  _______ 
(  ____ \\__   __/(  ___  )(  ____ )( \      \__   __/(  ____ \|\     /|\__   __/  (  ____ \(  ___  )(  ____ )(  ____ \(  ___  )
| (    \/   ) (   | (   ) || (    )|| (         ) (   | (    \/| )   ( |   ) (     | (    \/| (   ) || (    )|| (    \/| (   ) |
| (_____    | |   | (___) || (____)|| |         | |   | |      | (___) |   | |     | |      | (___) || (____)|| |      | |   | |
(_____  )   | |   |  ___  ||     __)| |         | |   | | ____ |  ___  |   | |     | |      |  ___  ||     __)| | ____ | |   | |
      ) |   | |   | (   ) || (\ (   | |         | |   | | \_  )| (   ) |   | |     | |      | (   ) || (\ (   | | \_  )| |   | |
/\____) |   | |   | )   ( || ) \ \__| (____/\___) (___| (___) || )   ( |   | |     | (____/\| )   ( || ) \ \__| (___) || (___) |
\_______)   )_(   |/     \||/   \__/(_______/\_______/(_______)|/     \|   )_(     (_______/|/     \||/   \__/(_______)(_______)
`)

	time.Sleep(time.Millisecond * 800)
	fmt.Printf("Welcome to Starlight Cargo - Your Galactic File Management System!\n\n")
	time.Sleep(time.Millisecond * 800)
	fmt.Printf("Initializing the interstellar transport layer...\n\n")

}
