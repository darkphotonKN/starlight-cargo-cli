package tcpclient

import (
	"fmt"
	"net"
	"time"
)

/**
* Responsible connecting to the server and initializing the CLI application.
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

	// make sure connection is accessible in the entire struct
	t.conn = conn

}

/**
* Prints the intro to the CLI app.
**/
func (t *TcpClient) printIntro() {

	t.writeConsole(`  
 _______ _________ _______  _______  _       _________ _______          _________   _______  _______  _______  _______  _______ 
(  ____ \\__   __/(  ___  )(  ____ )( \      \__   __/(  ____ \|\     /|\__   __/  (  ____ \(  ___  )(  ____ )(  ____ \(  ___  )
| (    \/   ) (   | (   ) || (    )|| (         ) (   | (    \/| )   ( |   ) (     | (    \/| (   ) || (    )|| (    \/| (   ) |
| (_____    | |   | (___) || (____)|| |         | |   | |      | (___) |   | |     | |      | (___) || (____)|| |      | |   | |
(_____  )   | |   |  ___  ||     __)| |         | |   | | ____ |  ___  |   | |     | |      |  ___  ||     __)| | ____ | |   | |
      ) |   | |   | (   ) || (\ (   | |         | |   | | \_  )| (   ) |   | |     | |      | (   ) || (\ (   | | \_  )| |   | |
/\____) |   | |   | )   ( || ) \ \__| (____/\___) (___| (___) || )   ( |   | |     | (____/\| )   ( || ) \ \__| (___) || (___) |
\_______)   )_(   |/     \||/   \__/(_______/\_______/(_______)|/     \|   )_(     (_______/|/     \||/   \__/(_______)(_______)
`, CYAN, NORMAL)

	t.newLine(3)
	time.Sleep(time.Millisecond * 800)
	t.writeConsole("Welcome to Starlight Cargo - Your Galactic File Management System!\n\n", CYAN, NORMAL)
	time.Sleep(time.Millisecond * 800)
	t.writeConsole("Initializing the interstellar transport layer...\n\n", CYAN, BOLD)

}
