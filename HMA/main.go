package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error in opening file")
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
