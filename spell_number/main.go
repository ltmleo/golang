package main
import "fmt"

var n int

var dict = map[int]string {
    0: "Zero",
    1: "One",
    2: "Two",
    3: "Three",
    4: "Four",
    5: "Five",
    6: "Six",
    7: "Seven",
    8: "Eight",
    9: "Nine",
    10: "Ten",
}

func main() {
    for i := 0; i < 3; i++{
        fmt.Scan(&n)
        fmt.Println(dict[n])
    }
}