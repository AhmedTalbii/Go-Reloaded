package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"goreloded"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid Number of arguments expected 2.")
		return
	}
	if !strings.HasSuffix(args[1], ".txt") || !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("Files Extentions need to be .txt")
		return
	}
	file, errOs := os.Open(args[0])
	if errOs != nil {
		fmt.Println("Error in opening file :", errOs)
		return
	}
	defer file.Close()
	data, errIo := io.ReadAll(file)
	if errIo != nil {
		fmt.Println("Error in read the file :", errIo)
		return
	}
	str := string(data)
	strL := goreloded.Solve(str)
	res, errRs := os.Create(args[1])
	if errRs != nil {
		fmt.Println("Error in creating file :", errRs)
		return
	}
	defer res.Close()

	_, errRf := res.WriteString(strL)

	if errRf != nil {
		fmt.Println("Error in Writing in the file :", errRf)
		return
	}
}
