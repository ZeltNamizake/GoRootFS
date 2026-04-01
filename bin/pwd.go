package bin

import (
	"fmt"
	"os"
)

func init() {
	Register("pwd", Pwd)
}

func Pwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(dir)
}