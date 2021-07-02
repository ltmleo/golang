package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go out(0, 5, ch)
	go out(6, 10, ch)
	<- ch
	evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSum(0, 10, evenCh)
    go squareSum(0, 100, sqCh)

	// select {
	// case x := <- evenCh:
	// 	fmt.Println("x:", x)
	// case y := <- sqCh:
	// 	fmt.Println("y:", y)
	// }
	n := 0
	var X int
	var Y int
	for {
        select {
            case x := <- evenCh:
                fmt.Println("x:", x)
				X = x
				n += 1
            case y := <- sqCh:
                fmt.Println("y:", y)
				Y = y
				n += 1				
            default:
                fmt.Println("Nothing available")
                time.Sleep(50 * time.Millisecond)
        }
		if (n == 2){
			fmt.Println(X + Y) // == fmt.Println(<- evenCh + <- sqCh)
			return
		}
    }

}

func out(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- true // close(ch) This is also done in the sender.

	
}

func evenSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i
        }    
    }
    ch <- result
}

func squareSum(from, to int, ch chan int) {
    result := 0
    for i:=from; i<=to; i++ {
        if i%2 == 0 {
            result += i*i
			//time.Sleep(500 * time.Millisecond)
        }    
    }
    ch <- result
}

