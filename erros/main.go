package main

import (
	"fmt"
	"net/http"
	"log"
	"errors"
)

func main(){
	res, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Erro ao fazer a comunicação!! "+err.Error())
	}

	fmt.Println(res.Header)
	resSoma, errSoma := soma(10, 0)
	if errSoma != nil {
		log.Fatal("Erro: "+errSoma.Error())
	}
	fmt.Println(resSoma)
	resSoma, _ = soma(10, 10) // Joga o valor "fora"
	fmt.Println(resSoma)

}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}
	return res, nil
}