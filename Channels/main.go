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
		//	does not work as expected because main go routine exits before
		//	child go routines finish
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err == nil {
		fmt.Printf("%v site is up and running.\n", link)
	} else {
		fmt.Printf("%v site seems to be down.\n", link)
	}
}
