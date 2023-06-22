package main

import (
	"fmt"
	"sync"
	"time"
)

func test() string {
	var wg sync.WaitGroup

	// Definir o número de iterações do loop
	numIterations := 100

	wg.Add(numIterations)

	resultChan := make(chan string)

	for i := 0; i < numIterations; i++ {
		go func(id int) {
			defer wg.Done()
			result := minhaFuncao(id)

			if result != "" {
				resultChan <- result
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		return result
	}

	return ""
}

func minhaFuncao(id int) string {
	time.Sleep(1*time.Second)
	if id == 17 {
		return "OK"
	}

	return ""
}

func main() {
	t := test()
	fmt.Print(t)
}
