package tcpclient

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
* File responsible for cli to tcp server inter-communications.
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
		t.writeConsole("Status: ", CYAN, NORMAL)
		t.writeConsole("Authenticated\n", CYAN, ITALIC)

		t.newLine(1)

		// -- menu selection --
		t.showMainMenu()
		msg, _ := reader.ReadString('\n')
		trimmedMsg := strings.TrimRight(msg, "\n")
		menuChoiceNo, err := strconv.Atoi(trimmedMsg)

		if err != nil {
			fmt.Println("Error when attempting to convert string to int:", err)
			continue
		}

		switch MenuChoice(menuChoiceNo) {
		case sendMessage:
			// -- message mode --
			t.writeConsole("Enter [cmd] + [message]: ", CYAN, NORMAL)

			msg, _ := reader.ReadString('\n')
			t.newLine(1)

			err := t.writeServerJwt(msg)

			if err != nil {
				t.writeConsole(fmt.Sprintf("Error sending message: %s", err), RED, BOLD)
				return err
			}

			// read response and log it
			res, err := bufio.NewReader(t.conn).ReadString('\n')
			if err != nil {
				t.writeConsole(fmt.Sprintf("Error when attempting to read from connection: %s", err), RED, BOLD)
			}

			t.writeConsole(fmt.Sprintf("Starlight Officer: %s", res), MAGENTA, NORMAL)

		case browseFiles:
			// -- show files --
			fmt.Println("Show files.")

		case downloadFile:
		case uploadFile:
			fmt.Println("uploading file")

		}

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
			t.writeConsole(fmt.Sprintf("Error when attempting to read from connection: %s", err), RED, BOLD)
		}

		t.writeConsole(fmt.Sprintf("Starlight Officer: %s", res), MAGENTA, NORMAL)

		// attempt to check if server provided an authenticated status
		resPair := strings.SplitN(res, ":", 2)

		// check the correct pre-determined format has been received before checking status
		if len(resPair) == 2 {
			status := resPair[0]
			serverMsg := resPair[1]

			// t.writeConsole(fmt.Sprintf("From Server:\nStatus: %s\n\nMessage: %s\n\n", status, serverMsg), MAGENTA, NORMAL)

			if status == AUTHENTICATED {
				t.newLine(2)
				fmt.Println("Server successfully authenticated..")
				t.newLine(1)
				t.accessToken = serverMsg

				// exit out of the infinite auth loop
				return nil
			}
		}

		t.writeConsole("Enter: ", CYAN, NORMAL)
		msg, _ := reader.ReadString('\n')
		t.newLine(1)

		_, err = t.conn.Write([]byte(msg))

		if err != nil {
			t.writeConsole(fmt.Sprintf("Error sending message: %s", err), RED, BOLD)
			return err
		}
	}

}
