package main

import "fmt"

func main() {
	A()

}

func A() (a int) {
	defer fmt.Println("a1: ", a) // 二者输出不同
	defer func() {
		fmt.Println("a2: ", a)
	}()
	a = 1

	return a
}
