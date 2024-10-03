package console

import (
	"github.com/fatih/color"
	"time"
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

type Console struct{}

func NewConsole() *Console {
	return &Console{}
}

/**
* Prints the intro to the CLI app.
**/
func (c *Console) PrintIntro() {

	c.WriteConsole(`  
 _______ _________ _______  _______  _       _________ _______          _________   _______  _______  _______  _______  _______ 
(  ____ \\__   __/(  ___  )(  ____ )( \      \__   __/(  ____ \|\     /|\__   __/  (  ____ \(  ___  )(  ____ )(  ____ \(  ___  )
| (    \/   ) (   | (   ) || (    )|| (         ) (   | (    \/| )   ( |   ) (     | (    \/| (   ) || (    )|| (    \/| (   ) |
| (_____    | |   | (___) || (____)|| |         | |   | |      | (___) |   | |     | |      | (___) || (____)|| |      | |   | |
(_____  )   | |   |  ___  ||     __)| |         | |   | | ____ |  ___  |   | |     | |      |  ___  ||     __)| | ____ | |   | |
      ) |   | |   | (   ) || (\ (   | |         | |   | | \_  )| (   ) |   | |     | |      | (   ) || (\ (   | | \_  )| |   | |
/\____) |   | |   | )   ( || ) \ \__| (____/\___) (___| (___) || )   ( |   | |     | (____/\| )   ( || ) \ \__| (___) || (___) |
\_______)   )_(   |/     \||/   \__/(_______/\_______/(_______)|/     \|   )_(     (_______/|/     \||/   \__/(_______)(_______)
`, CYAN, NORMAL)

	c.NewLine(3)
	time.Sleep(time.Millisecond * 800)
	c.WriteConsole("Welcome to Starlight Cargo - Your Galactic File Management System!\n\n", CYAN, NORMAL)
	time.Sleep(time.Millisecond * 800)
	c.WriteConsole("Initializing the interstellar transport layer...\n\n", CYAN, BOLD)
	c.WriteConsole("Please login to be authenticated.\n\n", CYAN, BOLD)
}

type MenuChoice int

const (
	SendMessage MenuChoice = iota + 1
	BrowseFiles
	UploadFile
	DownloadFile
)

/**
* Shows the main menu.
**/
func (c *Console) ShowMainMenu() {
	c.WriteConsole("Starlight Reception: Welcome. How may I assist you today?", MAGENTA, NORMAL)
	c.NewLine(1)
	c.NewLine(1)

	c.WriteConsole("1. Send a message to the server.\n", MAGENTA, NORMAL)
	c.WriteConsole("2. Browse Files.\n", MAGENTA, NORMAL)
	c.WriteConsole("3. Upload file.\n", MAGENTA, NORMAL)
	c.WriteConsole("4. Download file.\n", MAGENTA, NORMAL)

	c.NewLine(1)
}

/**
* Custom helper for styling and formatting messages.
**/
func (c *Console) WriteConsole(msg string, colorChoice Color, style Style) {
	var col *color.Color

	switch colorChoice {
	case CYAN:
		col = color.New(color.FgCyan)
	case RED:
		col = color.New(color.FgRed)
	case BLUE:
		col = color.New(color.FgBlue)
	case MAGENTA:
		col = color.New(color.FgMagenta)
	case WHITE:
		col = color.New(color.FgWhite)
	default:
		col = color.New(color.FgWhite)
	}

	switch style {
	case BOLD:
		col.Add(color.Bold)
	case UNDERLINE:
		col.Add(color.Underline)
	case ITALIC:
		col.Add(color.Italic)
	case NORMAL:
	default:
	}

	col.Print(msg)
}
