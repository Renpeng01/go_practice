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

// func main() {
// 	// SliceCopy()

// 	arr := [3]int{1, 2, 3}
// 	slice := arr[:]
// 	// 使用数组初始化另一个数组后，通过index改变元素，会源数组相同位置的元素
// 	// 使用下标初始化切片不会拷贝原数组或者原切片中的数据，只会创建一个指向原数组的切片结构体，修改新切片的数据也会修改原切片
// 	slice[0] = 0
// 	fmt.Println("arr: ", arr)
// 	fmt.Println("slice: ", slice)
// }

// func main() {
// 	a := []int{1, 2, 3, 4}
// 	s1 := a[0:2]

// 	fmt.Printf("s1 len: %+v, cap: %+v, s1: %+v\n", len(s1), cap(s1), s1)

// 	s1 = append(s1, 0)

// 	fmt.Printf("after s1 append a: %+v", a)

// }

func insert(slice []int, index int, value int) []int {
	// 在指定位置插入元素
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	v := 99 // 待插入元素

	a := slice[:2]
	b := slice[2:]

	// 正确方式
	t := append([]int{v}, b...)
	a = append(a, t...)

	// 错误方式
	// a = append(a, v)
	// a = append(a, b...)

	fmt.Println(a)

	// // 在索引 2 的位置插入元素 99
	// slice = insert(slice, 2, 99)

	// fmt.Println(slice) // 输出: [1 2 99 3 4 5]
}
