package main

import (
	"fmt"
	"time"
)

func repeat(s string, n int) chan string {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- s
		}
	}()
	return ch
}

func main() {
	fastChannel := repeat("fast", 550)
	slowChannel := repeat("slow", 1550)

	stop := time.After(5 * time.Second)
loop:
	for {
		select {
		case s := <-fastChannel:
			fmt.Println(s)
		case s := <-slowChannel:
			fmt.Println(s)
		case <-stop:
			break loop
		}
	}

}
