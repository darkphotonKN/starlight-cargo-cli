package tcpclient

import (
	"fmt"

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
* Handles formatting terminal and server message formatting and styling.
**/

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
