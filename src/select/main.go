package main

import "fmt"

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

func Expr(n int) int {
	fmt.Println(n)
	return n
}

func main() {
	switch Expr(2) {
	case Expr(1), Expr(2), Expr(3):
		fmt.Println("case 1")
		fallthrough // 强行执行下一个代码块且不进行条件判断
	case Expr(4):
		fmt.Println("case 2")
	}
}
