package bin

import (
	"fmt"
	"os"
)

func init() {
	Register("mkdir", Mkdir)
}

func Mkdir(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: mkdir <directory>")
		return
	}

	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("mkdir:", err)
		}
	}
}