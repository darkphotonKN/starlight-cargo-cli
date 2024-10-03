package tcpclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/types"
)

const (
	AUTHENTICATED = "AUTHENTICATED"
)

/**
* File responsible connecting to the server and initializing the CLI application.
**/

type TcpClient struct {
	addr        string
	conn        net.Conn
	accessToken string
	console     *console.Console
}

func NewTcpClient(addr string, console *console.Console) types.ClientConnector {
	return &TcpClient{
		addr:    addr,
		console: console,
	}
}

/**
* Connect to the server and set it for instance access.
**/
func (t *TcpClient) Connect() error {
	conn, err := net.Dial("tcp", t.addr)

	// make sure connection is accessible in the entire struct instance
	t.conn = conn

	if err != nil {
		fmt.Printf("Error when attempting to dial to tcp connection at %s, error: %s", t.addr, err)
		return err
	}
	// defer t.conn.Close()
	return nil
}

/**
* Read for email and password to authenticates user with a max attempt of 3 attempts.
**/
func (t *TcpClient) AuthenticateWithServer() error {
	// read in arguments to send over this tcp connection
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("@AuthenticateWithServer conn:", t.conn)

	for {
		// read response and log it
		res, err := bufio.NewReader(t.conn).ReadString('\n')

		if err != nil {
			t.console.WriteConsole(fmt.Sprintf("Error when attempting to read from connection: %s", err), console.BLUE, console.BOLD)
		}

		t.console.WriteConsole(fmt.Sprintf("Starlight Officer: %s", res), console.BLUE, console.BOLD)

		// attempt to check if server provided an authenticated status
		resPair := strings.SplitN(res, ":", 2)

		// check the correct pre-determined format has been received before checking status
		if len(resPair) == 2 {
			status := resPair[0]
			serverMsg := resPair[1]

			// t.writeConsole(fmt.Sprintf("From Server:\nStatus: %s\n\nMessage: %s\n\n", status, serverMsg), MAGENTA, NORMAL)

			if status == AUTHENTICATED {
				t.console.NewLine(2)
				fmt.Println("Server successfully authenticated..")
				t.console.NewLine(1)
				// set access token for global instance usage
				t.accessToken = serverMsg

				// exit out of the infinite auth loop
				return nil
			}
		}

		t.console.WriteConsole("Enter: ", console.BLUE, console.UNDERLINE)
		msg, _ := reader.ReadString('\n')
		t.console.NewLine(1)

		_, err = t.conn.Write([]byte(msg))

		if err != nil {
			t.console.WriteConsole(fmt.Sprintf("Error sending message: %s", err), console.BLUE, console.BOLD)
			return err
		}
	}
}

/**
* Helper methods to allow access to private instance variables that are central to the server.
**/

// connection instance
func (t *TcpClient) Conn() *net.Conn {
	return &t.conn
}

// instance access token
func (t *TcpClient) Token() string {
	return t.accessToken
}

// For setting of mock tests
func (t *TcpClient) SetConnection(conn net.Conn) {
	t.conn = conn
}
