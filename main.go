package main

import "fmt"

func studentIsApproved(avg int) bool {
	//a execucacao dessa funcao ira para o final da funcao, portanto executada antes do return
	defer fmt.Println("O resultado sera mostrado a seguir")

	return avg >= 5
}

func main() {
	fmt.Print(studentIsApproved(5))
}
