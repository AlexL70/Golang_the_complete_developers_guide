package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("Hex color for %v is %v\n", key, value)
	}
}
