package bin

import (
	"fmt"
	"os/user"
)

func init() {
	Register("whoami", Whoami)
}

func Whoami(args []string) {
	u, err := user.Current()
	if err != nil {
		fmt.Println("whoami: error getting username")
		return
	}
	fmt.Println(u.Username)
}