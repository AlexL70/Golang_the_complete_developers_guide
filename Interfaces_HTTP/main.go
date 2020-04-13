package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	_, err = resp.Body.Read(bs)
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
}
