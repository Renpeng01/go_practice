package main

import (
	"fmt"
	"sync"
)

// 并发append导致底层数组变化的问题
func main() {
	s := make([]int, 0, 6) // 未触发扩容 有问题
	// s := make([]int, 0, 2) // 触发扩容 没问题

	s = append(s, 1)
	s = append(s, 2)

	var wg sync.WaitGroup

	var m1 []int
	var m2 []int

	wg.Add(1)
	go func() {
		defer wg.Done()
		m2 = test(s, 4)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		m1 = test(s, 3)
	}()

	wg.Wait()

	fmt.Println(m1)
	fmt.Println(m2)
}

func test(a []int, i int) []int {
	a = append(a, i)
	return a
}
