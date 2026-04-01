package bin

import (
	"fmt"
	"os"
)

func init() {
	Register("mv", Mv)
}

func Mv(args []string) {
	if len(args) < 2 {
		fmt.Println("mv: missing source or destination")
		return
	}

	src := args[0]
	dst := args[1]

	err := os.Rename(src, dst)
	if err != nil {
		fmt.Println("mv:", err)
	}
}