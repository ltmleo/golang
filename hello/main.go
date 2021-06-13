package main

import "fmt"

var a string // Declaração

func main() {
	fmt.Println("Hello World!!")
	
	a = "Leo" // Atriuição
	b := "Montero" // Declaração e atribuição
	fmt.Printf("%v %v!!\n",a,b)

	c := 123
	d := true
	e := 1.4543

	fmt.Printf("%v=%T\n%v=%T\n%v=%T\n",c,c,d,d,e,e)


}
