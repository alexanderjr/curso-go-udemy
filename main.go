package main

import "fmt"

func closure() func() {
	text := "Clojure function"

	fn := func() {
		fmt.Print(text)
	}

	return fn
}

func main() {
	text := "Main function "
	fmt.Print(text)

	fn := closure()
	fn()
}
