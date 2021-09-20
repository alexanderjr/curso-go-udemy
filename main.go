package main

import "fmt"

func fibonnaci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

func main() {

	posicao := uint(7)
	for i := uint(0); i <= posicao; i++ {
		fmt.Printf("%d ", fibonnaci(i))
	}
}
