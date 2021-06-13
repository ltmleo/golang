package main

import (
	"fmt"
	"soma/math"
)

func main() {
	soma := Soma(1, 1)
	fmt.Printf("Soma: %v\n",soma)

	sub := math.Sub(1,1)
	fmt.Printf("Subtracao: %v\n",sub)
}

func Soma(a int, b int) int { // tem que colocar o tipo do retorno
	return a + b
}