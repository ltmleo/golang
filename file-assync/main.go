package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func tail(filename string, out io.Writer, done chan bool) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		select {
		case <-done:
			return
		default:
			for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
				if prefix {
					fmt.Fprint(out, filename, string(line))
				} else {
					fmt.Fprintln(out, filename, string(line))
				}
			}
		}
	}
}

func function(done chan bool) {
	time.Sleep(time.Second * 10)
	done <- true
}

func main() {
	done := make(chan bool)
	go tail("file.txt", os.Stdout, done)
	go tail("file2.txt", os.Stdout, done)
	function(done)
}
