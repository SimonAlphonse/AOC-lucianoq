package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		vm := NewVM("input", input, output)
		vm.Code[0] = 2
		vm.Run()
		close(output)
		done <- struct{}{}
	}()

	go func() {
		last := 0
		for x := range output {
			if x > 127 {
				fmt.Println(x)
				return
			}
			fmt.Print(string(x))
			if x == '\n' && last == '\n' {
				time.Sleep(100 * time.Millisecond)
				fmt.Print("\033[2J\033[H")
			}
			last = x
		}
	}()

	// Send main
	send(input, "A,B,A,C,A,A,C,B,C,B")

	// Send A
	send(input, "L,12,L,8,R,12")

	// Send B
	send(input, "L,10,L,8,L,12,R,12")

	// Send C
	send(input, "R,12,L,8,L,10")

	// continuous video feed
	send(input, "y")

	<-done
}

func send(input chan<- int, s string) {
	for _, r := range s {
		input <- int(r)
	}
	input <- '\n'
}
