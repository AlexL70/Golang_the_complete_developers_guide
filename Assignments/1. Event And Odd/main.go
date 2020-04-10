package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, s := range slice {
		if s%2 == 0 {
			fmt.Printf("%d is even\n", s)
		} else {
			fmt.Printf("%d is odd\n", s)
		}
	}
}
