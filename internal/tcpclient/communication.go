package tcpclient

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
* Responsible for cli to tcp server inter-communications.
**/

const (
	AUTHENTICATED = "AUTHENTICATED"
)

/**
* Starts a cli application for interaction between the cli and the Starlight Cargo tcp server.
**/
func (t *TcpClient) communicateWithServer() error {

	// -- attempt to authenticate --
	err := t.authenticateWithServer()

	if err != nil {
		fmt.Println("Error during communication. Exiting...")
		return err
	}

	// -- access server after authenticated --
	err = t.commandMessageLoop()

	if err != nil {
		fmt.Println("Error during communication. Exiting...")

		return err
	}

	return nil
}

/**
* Communicates with server with a command message loop until errored or
* TODO: connection stays idle for too long.
**/
func (t *TcpClient) commandMessageLoop() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Entering Authenticated Message Loop")

		fmt.Printf("accessToken: %s\n\n", t.accessToken)

		fmt.Print("Enter [cmd] + [message]: ")
		msg, _ := reader.ReadString('\n')
		fmt.Println()

		// append jwt to message for authorization
		// message format is [jwt accessToekn] [cmd] [message]
		authMsg := fmt.Sprintf("%s %s", t.accessToken, msg)

		// TODO: use jwt access token

		_, err := t.conn.Write([]byte(authMsg))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return err
		}

		// read response and log it
		res, err := bufio.NewReader(t.conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error when attempting to read from connection.")
		}

		fmt.Println("Starlight Officer:", res)
	}

}

/**
* Read for email and password to authenticates user with a max attempt of 3 attempts.
**/
func (t *TcpClient) authenticateWithServer() error {
	// read in arguments to send over this tcp connection
	reader := bufio.NewReader(os.Stdin)

	for {
		// read response and log it
		res, err := bufio.NewReader(t.conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error when attempting to read from connection during authentication loop.")
		}

		fmt.Println("Starlight Officer:", res)

		// attempt to check if server provided an authenticated status
		resPair := strings.SplitN(res, ":", 2)

		// check the correct pre-determined format has been received before checking status
		if len(resPair) == 2 {
			status := resPair[0]
			serverMsg := resPair[1]

			fmt.Printf("From Server:\nStatus: %s\nMessage: %s\n\n", status, serverMsg)

			if status == AUTHENTICATED {
				fmt.Println("Server authenticated.")
				t.accessToken = serverMsg // temp
				// exit out of the infinite auth loop

				return nil
			}
		}

		fmt.Print("Enter: ")
		msg, _ := reader.ReadString('\n')
		fmt.Println()

		_, err = t.conn.Write([]byte(msg))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return err
		}
	}
}
