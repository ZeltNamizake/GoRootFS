package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"GoRootFS/bin"
)

func getPrompt() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "unknown"
	}
	
	home := os.Getenv("HOME")
	if strings.HasPrefix(dir, home) {
		dir = "~" + strings.TrimPrefix(dir, home)
	}

	return fmt.Sprintf("%s:$ ", dir)
}

var history []string

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
	  time.Sleep(100 * time.Millisecond)
		fmt.Print("\n" + getPrompt())

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		
		history = append(history, input)

		if input == "exit" {
			break
		}

		args := strings.Fields(input)
		cmd := args[0]

		if fn, ok := bin.Commands[cmd]; ok {
			fn(args[1:])
		} else {
			fmt.Println("unknown command")
		}
	}
}