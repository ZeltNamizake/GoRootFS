package bin

import (
	"fmt"
	"strings"
)

func init() {
	Register("echo", Echo)
}

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}