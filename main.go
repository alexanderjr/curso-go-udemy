package main

import "fmt"

func main() {
	var fn func(message string)

	fn = func(message string) {
		fmt.Printf("Hello World %s", message)
	}

	fn("Alex")
}
