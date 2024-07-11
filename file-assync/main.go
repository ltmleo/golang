package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sync"
    "time"
)


func readLogFile(logFile string, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    file, err := os.Open(logFile)
    if err != nil {
        fmt.Println("Error opening log file:", err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        ch <- line
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading log file:", err)
    }
    for {
        stat, err := file.Stat()
        if err != nil {
            fmt.Println("Error getting file stats:", err)
            return
        }
        lastModTime := stat.ModTime()
        time.Sleep(time.Second) 
        stat, err = file.Stat()
        if err != nil {
            fmt.Println("Error getting file stats:", err)
            return
        }
        if stat.ModTime().After(lastModTime) {
            file.Seek(0, io.SeekCurrent)
            scanner = bufio.NewScanner(file)

            for scanner.Scan() {
                line := scanner.Text()
                ch <- line
            }

            if err := scanner.Err(); err != nil {
                fmt.Println("Error reading log file:", err)
            }
        }
    }
}
func executeFunction() {
	time.Sleep(time.Second*10)
}

func main() {
    ch := make(chan string)
	ch2 := make(chan string)
    wg := sync.WaitGroup{}
    wg.Add(1)
    go readLogFile("file.txt", ch, &wg)
    for line := range ch {
        fmt.Println("Log line:", line)
	}
	wg.Add(1)
	go readLogFile("file2.txt", ch2, &wg)
    for line := range ch2 {
        fmt.Println("Log line 2:", line)
	}
    wg.Add(1)
    go executeFunction()
    wg.Wait()

    fmt.Println("Log file processing complete.")
}
