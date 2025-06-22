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
		fmt.Println("Invalid Input.")
		return
	}
	file, errOs := os.Open(args[0])
	if errOs != nil {
		fmt.Println("Error in opening file :",errOs)
		return
	}
	defer file.Close()
	data, errIo := io.ReadAll(file)
	if errIo != nil {
		fmt.Println("Error in read the file :",errOs)
		return
	}
	str := string(data)
	strL := goreloded.Solve(str)
	res , errRs := os.Create(args[1])
	if errRs != nil {
		fmt.Println("Error in creating file :",errOs)
		return
	}
	defer res.Close()
	
	_ , errRf := res.WriteString(strL)
	if errRf != nil {
		fmt.Println("Error in Writing in the file :",errRf)
		return
	}
}
