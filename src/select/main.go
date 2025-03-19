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
	var ch chan int

	select {
	case <-ch:
		fmt.Println("<-ch")
	default:
		fmt.Println("exec default")
	}
}
