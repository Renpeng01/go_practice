package main

// func main() {
// 	A()

// }

// func A() (a int) {
// 	defer fmt.Println("a1: ", a) // 二者输出不同
// 	defer func() {
// 		fmt.Println("a2: ", a)
// 	}()
// 	a = 1

// 	return a
// }

// import "fmt"

// func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

// 	defer func() {
// 		t = t * 10
// 	}()

// 	return 1
// }

// func main() {
// 	fmt.Println(returnButDefer())
// }

// import (
// 	"fmt"
// )

// func main() {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("11111", err)
// 		} else {
// 			fmt.Println("fatal")
// 		}
// 	}()

// 	defer func() {
// 		panic("defer panic")
// 	}()

// 	panic("panic")
// }

import "fmt"

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}

// 4 1 3 0 2
