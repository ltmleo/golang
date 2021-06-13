package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a, &a) // &a mostra o endereco de MEM em que a está alocada
	// 0xc0000180c8(endereco-memoria) <---- a <---- 10
	var ponteiro *int = &a
	fmt.Println(ponteiro, *ponteiro) // *pont retorna o valor guardado nessa endereco
	*ponteiro = 50 //altera o valor contido no endereco
	fmt.Println(a) 
	b := &a //  b alocado no mesmo end de memoria
	fmt.Println(*b)
	*b = 100 // altera o valor de b, do endereco e por fim de a
	fmt.Println(a)
	abc(&a) // altera o valor de a, alterando o valor do ponteiro e não retornando nada
	fmt.Println(a)

	carro := Carro{
		Nome: "Fuscão",
	}
	carro.andar()
	println(carro.Nome)  // sem o *Carro, retorna 'fuscao', como *Carro retonar komboza pois alterou o ponteiro (GLOBAL)
}


func abc(a *int){
	*a = 200
}

type Carro struct {
	Nome string
}

func (c *Carro) andar(){    // sem o *Carro, retorna 'fuscao', como *Carro retonar komboza pois alterou o ponteiro  (GLOBAL)
	c.Nome = "Komboza"
	fmt.Println(c.Nome, "Andou!")
}
