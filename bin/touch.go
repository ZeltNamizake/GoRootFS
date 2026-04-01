package bin

import (
	"fmt"
	"os"
	"time"
)

func init() {
	Register("touch", Touch)
}

func Touch(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: touch <file>...")
		return
	}

	for _, file := range args {

		if _, err := os.Stat(file); os.IsNotExist(err) {

			f, err := os.Create(file)
			if err != nil {
				fmt.Println("touch:", err)
				continue
			}
			f.Close()
		} else if err != nil {
			fmt.Println("touch:", err)
			continue
		} else {

			now := time.Now()
			err := os.Chtimes(file, now, now)
			if err != nil {
				fmt.Println("touch:", err)
			}
		}
	}
}