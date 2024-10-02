package communication

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/types"
)

/**
* File responsible for cli to tcp server inter-communications.
**/

type Communication struct {
	console *console.Console
	client  types.ClientConnector
}

func NewCommunicationService(client types.ClientConnector, console *console.Console) *Communication {
	return &Communication{
		console: console,
		client:  client,
	}
}

/**
* Starts a cli application for interaction between the cli and the Starlight Cargo tcp server.
**/
func (c *Communication) CommunicateWithServer() error {
	// -- attempt to authenticate --
	err := c.client.AuthenticateWithServer()

	if err != nil {
		fmt.Println("Error during communication. Exiting...")
		return err
	}

	// -- access server after authenticated --
	err = c.commandMessageLoop()

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
func (c *Communication) commandMessageLoop() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		c.console.WriteConsole("Status: ", console.CYAN, console.NORMAL)
		c.console.WriteConsole("Authenticated\n", console.CYAN, console.ITALIC)

		c.console.NewLine(1)

		// -- menu selection --
		c.console.ShowMainMenu()
		msg, _ := reader.ReadString('\n')
		trimmedMsg := strings.TrimRight(msg, "\n")
		menuChoiceNo, err := strconv.Atoi(trimmedMsg)

		if err != nil {
			fmt.Println("Error when attempting to convert string to int:", err)
			continue
		}

		switch console.MenuChoice(menuChoiceNo) {
		case console.SendMessage:
			// -- message mode --
			c.console.WriteConsole("Enter [cmd] + [message]: ", console.CYAN, console.NORMAL)

			msg, _ := reader.ReadString('\n')
			c.console.WriteConsole(msg, console.CYAN, console.NORMAL)
			c.console.NewLine(1)

			// read response and log it
			conn := c.client.Conn()
			res, err := bufio.NewReader(*conn).ReadString('\n')
			if err != nil {
				c.console.WriteConsole(fmt.Sprintf("Error when attempting to read from connection: %s", err), console.RED, console.BOLD)
			}

			c.console.WriteConsole(fmt.Sprintf("Starlight Officer: %s", res), console.MAGENTA, console.NORMAL)

		case console.BrowseFiles:
			// -- show files --
			fmt.Println("Show files.")

		case console.DownloadFile:
		case console.UploadFile:
			fmt.Println("uploading file")

		}

	}

}
