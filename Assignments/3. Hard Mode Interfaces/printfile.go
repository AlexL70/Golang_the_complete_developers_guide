package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: printfile <file_name>")
		os.Exit(1)
	}

	file, error := os.Open(os.Args[1])
	if error != nil {
		log.Fatal(error)
	}

	io.Copy(os.Stdout, file)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
