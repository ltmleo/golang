package main
import "fmt"

func main() {
    var age int
    fmt.Scanln(&age)

    mars := mars_age(age)
    fmt.Println(mars)
}

func mars_age(age int) int {
	days := age * 365
	return int(days / 687) 
}