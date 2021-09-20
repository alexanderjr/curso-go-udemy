package main

import "fmt"

func recoverWhenFailed() {
	if fn := recover(); fn != nil {
		fmt.Println("Recuperando fn que caiu no panic")
	}
}

func studentIsApproved(n1, n2 int) bool {
	defer recoverWhenFailed()
	sum := n1 + n2

	if sum == 0 {
		panic("A soma das medias não pode ser <= 0!!!")
	}
	avg := n1 + n2/2
	defer fmt.Println("O resultado sera mostrado a seguir")

	return avg >= 5
}

func main() {
	fmt.Println(studentIsApproved(0, 0))
	fmt.Print("Fim")
}
