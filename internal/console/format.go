package console

import "fmt"

/**
* Simple new-line method for code cleaniness.
**/
func (c *Console) NewLine(n int) {
	if n == 0 || n == 1 {
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		fmt.Println()
	}
}
