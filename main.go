package main

import "fmt"

func generic(a interface{}) {
	fmt.Print(a)
}

func main() {
	generic(1)
	generic("Alex")
	generic(true)
}
