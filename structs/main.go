package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar(){   // Um bind entre o metodo e a struct
	fmt.Println(c.Nome, "Andou!")
}

func main() {
	carro := Carro{
		Nome: "Fuscão",
	}
	carro.andar()
}