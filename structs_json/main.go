package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome string
	Email string
	CPF int
}


type ClienteInt struct {
	Cliente
	Pais string `json:"pais-example_de_tag"` // tag para quando rodar algm metodo json transformar o nome
}



func main() {
	cliente := Cliente{
		Nome: "Leonardo",
		Email: "l@l.com",
		CPF: 12345678900,
	}
	fmt.Println(cliente)

	cliente2 := Cliente{"Lais", "la@l.com", 12345678901}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d \n", cliente2.Nome, cliente2.Email, cliente2.CPF)

	cliente3 := ClienteInt{
		Cliente: Cliente {
			Nome: "Dani",
			Email: "d@d.com",
			CPF: 12345678902,
		},
		Pais: "Portugal",
	}
	fmt.Printf("Nome: %s, Email: %s, CPF: %d, Pais: %s\n", cliente3.Cliente.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)
	// Funciona acessando todos os niveis ou como herança

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson))

	jsonClient4 := `{"Nome":"Lucca","Email":"lu@l.com","CPF":12345678903,"pais-example_de_tag":"Inglaterra"}`
	cliente4 := ClienteInt {}

	json.Unmarshal([]byte(jsonClient4), &cliente4) // & para alterar o ponteiro
	fmt.Println(cliente4)

}