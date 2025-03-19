package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(1)
	s, ok := v.Interface().([]int)
	fmt.Println(s, ok)

	// i := 1
	// v := reflect.ValueOf(&i)
	// v.Elem().SetInt(10)
	// fmt.Println(i)
}
