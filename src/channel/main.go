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

	ch := make(chan int)
	_, ok := <-ch // <-ch 效果相同 不会提前退出

	fmt.Println(ok, "stop .................")
}
