package bin

import (
	"fmt"
	"os"
)

const (
  bold = "\033[1m"
	colorBlue  = "\033[34m"
	colorGray  = "\033[90m"
	colorReset = "\033[0m"
)

func init() {
	Register("ls", Ls)
}

func Ls(args []string) {
	dir := "."

	if len(args) > 0 {
		dir = args[0]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for _, f := range files {
		name := f.Name()

		if f.IsDir() {
			fmt.Println(bold + colorBlue + name + colorReset)
		} else {
			fmt.Println(bold + colorGray + name + colorReset)
		}
	}
}