package bin

import "fmt"

func init() {
	Register("clear", Clear)
}

func Clear(args []string) {
	fmt.Print("\033[H\033[2J")
}