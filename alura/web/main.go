package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome string `template:"nome"`
	Descricao string `template:"descricao"`
	Preco float64 `template:"preco"`
	Quantidade int `template:"quantidade"`
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 39., Quantidade: 2},
		{"Tenis", "Confortavel", 309.90, 5},
		{"Fone", "Topzera", 899., 50},
	}
	temp.ExecuteTemplate(w, "index", produtos)
}
