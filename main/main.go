package main

import (
	"fmt"
	"goreloded"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 || args[1] == "main.go" || args[1] == "sample.txt" || args[1] == "../helperFuncs.go" || args[1] == "../hexBinToDec.go" || args[1] == "../ordersHandle.go" || args[1] == "../punctuationMarks.go" || args[1] == "../singleQuotes.go" || args[1] == "../solve.go" {
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
	strL := goreloded.Solve(str)
	res , errRs := os.Create(args[1])
	if errRs != nil {
		return
	}
	_ , errRf := res.WriteString(strL)
	if errRf != nil {
		return
	}
	fmt.Println(strL)
}
