package bin

import (
	"fmt"
	"os"
)

func init() {
	Register("rm", Rm)
}

func Rm(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: rm [-r] <file/dir>...")
		return
	}

	rec := false
	files := args

	if args[0] == "-r" {
		rec = true
		if len(args) < 2 {
			fmt.Println("rm: missing target after -r")
			return
		}
		files = args[1:]
	}

	for _, f := range files {
		info, err := os.Stat(f)
		if err != nil {
			fmt.Println("rm:", err)
			continue
		}

		if info.IsDir() {
			if rec {
				err = os.RemoveAll(f)
				if err != nil {
					fmt.Println("rm:", err)
				}
			} else {
				fmt.Println("rm:", f, "is a directory (use -r to remove)")
			}
		} else {
			err = os.Remove(f)
			if err != nil {
				fmt.Println("rm:", err)
			}
		}
	}
}