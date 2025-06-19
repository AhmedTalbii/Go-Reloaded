package main

import (
	"fmt"
	"io"
	"os"
	"goreloded"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	file, errOs := os.Open(args[0])
	if errOs != nil {
		return
	}
	data, errIo := io.ReadAll(file)
	if errIo != nil {
		return
	}
	str := string(data)
	fmt.Print(goreloded.Solve(str))
}
