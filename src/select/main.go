package main

import (
	"fmt"
	"time"
)

// func main() {

// 	go func() {
// 		for {
// 			fmt.Println(time.Now())
// 			time.Sleep(3 * time.Second)
// 		}
// 	}()
// 	// var ch chan int

// 	// ch := make(chan int)
// 	// select {
// 	// case ch <- 1:
// 	// 	fmt.Println("<-ch")
// 	// default:
// 	// 	fmt.Println("exec default")
// 	// }

// 	select {}

// }

// func Expr(n int) int {
// 	fmt.Println(n)
// 	return n
// }

// func main() {
// 	switch Expr(2) {
// 	case Expr(1), Expr(2), Expr(3):
// 		fmt.Println("case 1")
// 		fallthrough // 强行执行下一个代码块且不进行条件判断
// 	case Expr(4):
// 		fmt.Println("case 2")
// 	}
// }

func getAReadOnlyChannel() <-chan int {
	fmt.Println("invoke getAReadOnlyChannel")
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getASlice() *[5]int {
	fmt.Println("invoke getASlice")
	var a [5]int
	return &a
}

func getAWriteOnlyChannel() chan<- int {
	fmt.Println("invoke getAWriteOnlyChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func main() {
	select {
	// 从channel接收数据
	case (getASlice())[0] = <-getAReadOnlyChannel(): // case 等号左边的从Channel接收数据的表达式不会被先求值（getASlice）。[惰性求值]
		fmt.Println("recv something from a readonly channel")
	// 将数据发送到channel
	case getAWriteOnlyChannel() <- getANumToChannel():
		fmt.Println("send something to a writeonly channel")
	}
}
