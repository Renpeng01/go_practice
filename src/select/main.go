package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println(time.Now())
			time.Sleep(3 * time.Second)
		}
	}()
	// var ch chan int

	ch := make(chan int, 10)
	select {
	case ch <- 1:
		fmt.Println("<-ch")
		// default:
		// 	fmt.Println("exec default")
	}
}
