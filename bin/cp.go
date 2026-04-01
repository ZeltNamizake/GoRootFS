package bin

import (
	"fmt"
	"io"
	"os"
)

func init() {
	Register("cp", Cp)
}

func Cp(args []string) {
	if len(args) < 2 {
		fmt.Println("cp: missing source or destination")
		return
	}

	src := args[0]
	dst := args[1]

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		fmt.Println("cp:", err)
		return
	}

	if sourceFileStat.IsDir() {
		fmt.Println("cp: copying directories not supported yet")
		return
	}

	source, err := os.Open(src)
	if err != nil {
		fmt.Println("cp:", err)
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("cp:", err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("cp:", err)
	}
}