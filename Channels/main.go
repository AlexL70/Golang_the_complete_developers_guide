package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		if checkLink(link) {
			fmt.Printf("%v site is up and running.\n", link)
		} else {
			fmt.Printf("%v site seems to be down.\n", link)
		}
	}
}

func checkLink(link string) bool {
	_, err := http.Get(link)
	return err == nil
}
