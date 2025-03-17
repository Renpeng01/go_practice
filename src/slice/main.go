package main

import "fmt"

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

type E struct {
	Val int
}

func SliceCopy() {
	a := make([]*E, 0, 1)
	a = append(a, &E{
		Val: 1,
	})

	b := make([]*E, 10) // 这里的len不能是0

	copy(b, a)
	fmt.Printf("origin a: %+v, b: %+v\n", a[0].Val, b[0].Val)

	b[0].Val = 2
	fmt.Printf("change a: %+v, b: %+v\n", a[0].Val, b[0].Val)
}

func main() {
	// SliceCopy()

	arr := [3]int{1, 2, 3}
	slice := arr[:]
	// 使用数组初始化另一个数组后，通过index改变元素，会源数组相同位置的元素
	slice[0] = 0
	fmt.Println("arr: ", arr)
	fmt.Println("slice: ", slice)

}
