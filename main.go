package main

import "fmt"

var globalValue int

func init() {
	fmt.Print("Print init function por arquivo")
	globalValue = 10
}

func main() {
	fmt.Print("Main")
	fmt.Print(globalValue)
}
