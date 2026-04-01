package bin

import (
	"fmt"
	"os"
)

func init() {
	Register("cd", Cd)
}

func Cd(args []string) {
	var target string

	if len(args) == 0 {
		target = os.Getenv("HOME")
		if target == "" {
			fmt.Println("cd: HOME not set")
			return
		}
	} else {
		target = args[0]
	}

	err := os.Chdir(target)
	if err != nil {
		fmt.Println("cd:", err)
	}
}