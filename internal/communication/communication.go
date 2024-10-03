package communication

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/darkphotonKN/starlight-cargo-cli/internal/console"
	fileservice "github.com/darkphotonKN/starlight-cargo-cli/internal/file_service"
	"github.com/darkphotonKN/starlight-cargo-cli/internal/types"
)

const (
	fileSizeLimit int = 2048
)

/**
* File responsible for cli to tcp server inter-communications.
**/

type Communication struct {
	client      types.ClientConnector
	fileService *fileservice.FileService
	console     *console.Console
}

func NewCommunicationService(client types.ClientConnector, fileService *fileservice.FileService, console *console.Console) *Communication {
	return &Communication{
		client:      client,
		fileService: fileService,
		console:     console,
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

		// -- Show Menu Selection --
		c.console.ShowMainMenu()
		msg, _ := reader.ReadString('\n')
		trimmedMsg := strings.TrimRight(msg, "\n")
		menuChoiceNo, err := strconv.Atoi(trimmedMsg)

		if err != nil {
			fmt.Println("Error when attempting to convert string to int:", err)
			continue
		}

		switch console.MenuChoice(menuChoiceNo) {
		// -- message mode --
		case console.SendMessage:
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

			// -- show files --
		case console.BrowseFiles:
			fmt.Println("Show files.")

		case console.DownloadFile:
			// -- upload file --
		case console.UploadFile:
			// request for file
			c.console.WriteConsole("Starlight Reception Desk: Please enter your exact filepath:", console.MAGENTA, console.NORMAL)

			// read file input
			filePath, _ := reader.ReadString('\n')

			fmt.Println("filePath:", filePath)

			trimmedFilePath := strings.Trim(filePath, "\n")

			file, err := os.Open(trimmedFilePath)

			buf := make([]byte, fileSizeLimit)

			n, err := file.Read(buf)

			c.console.NewLine(2)

			if err != nil {
				c.console.WriteConsole("Starlight Reception Desk: File size was too large.", console.MAGENTA, console.BOLD)
				c.console.NewLine(2)

				fmt.Println("Err:", err)
				return err
			}

			readFile := buf[:n]

			c.fileService.UploadFile(readFile)
		}
	}

}
