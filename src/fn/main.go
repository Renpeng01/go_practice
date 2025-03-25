package main

import "fmt"

// func main() {
// 	t := &A{
// 		a: 10,
// 	}

// 	fmt.Println("sss", WithT(2, t.TTT))

// }

// func WithT(m int, fn func(int) int) int {
// 	return fn(m)
// }

// type A struct {
// 	a int
// }

// func (a *A) TTT(m int) int {
// 	fmt.Println(a.a, m)

// 	return 3
// }

func main() {

	a := 1

	fn := func() {
		a++
		fmt.Println(a)
	}

	fn()
	fn()
}
