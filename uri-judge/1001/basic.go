package main

import (
	"fmt"
	"io"
	"os"
)

func Add(x int, y int) int {
	return x + y
}

func Show(writer io.Writer, result int) {
    fmt.Fprintf(writer, "X = %d\n", result)
}

func main(){
    var x int
	var y int
    fmt.Scanln(&x)
	fmt.Scanln(&y)
	result := Add(x, y)
	Show(os.Stdout, result)
}