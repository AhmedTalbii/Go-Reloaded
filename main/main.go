package main

import (
	"fmt"
	"goreloded"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 || strings.HasSuffix(args[1], ".go") || !strings.HasSuffix(args[1], ".txt") {
		return
	}
	file, errOs := os.Open(args[0])
	if errOs != nil {
		return
	}
	defer file.Close()
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
	defer res.Close()
	
	_ , errRf := res.WriteString(strL)
	if errRf != nil {
		return
	}
	fmt.Println(strL)
}
