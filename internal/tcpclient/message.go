package tcpclient

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Color int

const (
	CYAN Color = iota
	RED
	BLUE
	MAGENTA
	WHITE
)

type Style int

const (
	NORMAL Style = iota
	UNDERLINE
	BOLD
	ITALIC
)

/**
* File handles terminal and server message sending formatting, styling.
**/

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

type MenuChoice int

const (
	sendMessage MenuChoice = iota + 1
	browseFiles
	uploadFile
	downloadFile
)

/**
* Shows the main menu.
**/
func (t *TcpClient) showMainMenu() {

	t.writeConsole("Starlight Reception Desk: Welcome. How may I assist you today?", MAGENTA, NORMAL)
	t.newLine(1)
	t.newLine(1)

	t.writeConsole("1. Send a message to the server.\n", MAGENTA, NORMAL)
	t.writeConsole("2. Browse Files.\n", MAGENTA, NORMAL)
	t.writeConsole("3. Upload file.\n", MAGENTA, NORMAL)
	t.writeConsole("4. Download file.\n", MAGENTA, NORMAL)

	t.newLine(1)
}

/**
* Simple new-line method for code cleaniness.
**/
func (t *TcpClient) newLine(n int) {
	if n == 0 || n == 1 {
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		fmt.Println()
	}
}

/**
* Custom helper for styling and formatting messages.
**/
func (t *TcpClient) writeConsole(msg string, colorChoice Color, style Style) {
	var c *color.Color

	switch colorChoice {
	case CYAN:
		c = color.New(color.FgCyan)
	case RED:
		c = color.New(color.FgRed)
	case BLUE:
		c = color.New(color.FgBlue)
	case MAGENTA:
		c = color.New(color.FgMagenta)
	case WHITE:
		c = color.New(color.FgWhite)
	default:
		c = color.New(color.FgWhite)
	}

	switch style {
	case BOLD:
		c.Add(color.Bold)
	case UNDERLINE:
		c.Add(color.Underline)
	case ITALIC:
		c.Add(color.Italic)
	case NORMAL:
	default:
	}

	c.Print(msg)
}

/**
* Writes a message to server in a slice of bytes with the auth token appended.
**/
func (t *TcpClient) writeServerJwt(msg string) error {
	// append jwt to message for authorization
	// message format is [jwt accessToekn] [cmd] [message]
	authMsg := fmt.Sprintf("%s %s", t.accessToken, msg)

	_, err := t.conn.Write([]byte(authMsg))

	return err
}
