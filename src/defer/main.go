package main

import "fmt"

func main() {
	A()

}

func A() (a int) {
	// defermt.Println("a: ", a)  二者效果不同
	defer func() {
		fmt.Println("a: ", a)
	}()
	a = 1

	return a
}
