package bin

import (
	"fmt"
	"io"
	"os"
)

func init() {
	Register("cat", Cat)
}

func Cat(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: cat <file>")
		return
	}

	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			fmt.Println("err:", err)
		}

		f.Close()
	}
}