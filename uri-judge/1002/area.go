package main

import (
	"fmt"
)

const pi = 3.14159


func Area(radius float64) float64 {
	area := Round(pi * radius * radius, 10000)
	return area
}

func Round(number float64, precision float64) float64 {
	trunked := float64(int(number * precision)) / precision
	delta := number - trunked
	if (delta * precision >= 0.5) {
		return trunked + (1/precision)
	} else {
		return trunked
	}
}

func main(){
	var radius float64
	fmt.Scanln(&radius)
	area := Area(radius)
	fmt.Printf("A=%.4f\n", area)
}


