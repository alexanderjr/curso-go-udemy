package main

import "fmt"

func write(text string, numbers ...int) {
	for c := 0; c < len(numbers); c++ {
		fmt.Printf("%s %d \n", text, numbers[c])
	}
}
func main() {
	write("Hello World", 1, 2, 3)
}
