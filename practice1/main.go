package main

import (
	"fmt"
	"io"
	"os"
)

type myfileLoger struct{}

func main() {
	args := os.Args
	file1 := args[1]
	file, err := os.Open(file1)
	if err != nil {
		fmt.Println(err)
	}
	fileLogger := myfileLoger{}
	io.Copy(fileLogger, file)
}

func (lw myfileLoger) Write(file []byte) (int, error) {
	fmt.Println(string(file))
	return len(file), nil
}
