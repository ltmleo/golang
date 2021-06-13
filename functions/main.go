package main

import "fmt"

func main() {
	fmt.Println(soma(1,2))
	fmt.Println(somaTudo(1,2,3,4,5,6,7,8))
}

func soma(x int, y int) (res int){
	res = x + y
	return // Como passei o nome na decalaração do retorno posso apenas colocar return
}

func somaTudo(x ... int) (res int){ // Função variatica, Posso pasar quantos args quiser (~ kargs kwargs)
	res = 0
	for _, v := range x { // for tem uma key, não vou usar, coloco o _
		res += v
	}
	return
}