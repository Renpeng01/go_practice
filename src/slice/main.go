package main

import "fmt"

func main() {
	subSlice()
}

// https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/
// 需要注意的是使用下标初始化切片不会拷贝原数组或者原切片中的数据，它只会创建一个指向原数组的切片结构体，所以修改新切片的数据也会修改原切片
func subSlice() {
	s := make([]int, 0, 3)

	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Printf("origin slice: %+v\n", s)

	sub := s[0:1]
	fmt.Printf("sub slice: %+v\n", sub)

	sub[0] = 10
	fmt.Printf("change sub slice: %+v, origin slice: %+v\n", sub, s)
}
