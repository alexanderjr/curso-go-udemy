package main

import "fmt"

func toNegative(n *int) {
	*n = *n * -1
}

func main() {
	var n int = 10
	toNegative(&n)
	fmt.Print(n)
}
