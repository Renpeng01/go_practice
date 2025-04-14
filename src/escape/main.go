package main

// go build -gcflags="-m" main.go or go tool comploie -m xxx.go

func main() {

	// ** []interface{}数据类型，通过[]赋值必定会出现他逃逸
	// data := []interface{}{100, 200}
	// data := []int{100, 200}
	// data[0] = 100 // 100 发生了逃逸

	// ** map[string]interface{} 类型通过赋值，必定出现逃逸
	// data := make(map[string]interface{})
	// data := make(map[string]int)
	// data["key"] = 200 // 200 发生了逃逸

	// ** map[interface{}]interface{} 类型通过赋值，会导致key和value的赋值，出现逃逸
	// data := make(map[interface{}]interface{})
	// data[100] = 200 // 100 和 200 发生了逃逸

	// **
	// data := make(map[string][]string)
	// data["key"] = []string{"value"} // []string{"value"} 发生了逃逸

	// ** []*int 数据类型，赋值的右值发生逃逸
	// a := 10
	// data := []*int{nil}
	// data[0] = &a // 赋值右值（a）发生逃逸

	// ** func(*int)函数类型，进行函数赋值，会使传递的形参出现逃逸现象
	// data := 10
	// f := foo
	// f(&data)
	// fmt.Println(data)

	// ** func([]string) 函数类型 进行[]string{"val"}赋值，会使传递的参数出现逃逸
	// s := []string{"rp"}
	// foo1(s)
	// fmt.Println(s)

	// chan []string 数据类型 向当前channel中传输[]string{"val"}会发生逃逸
	ch := make(chan []string)
	s := []string{"aceld"}
	go func() {
		ch <- s
	}()
}

func foo1(a []string) {
	return
}

func foo(a *int) {
	return
}
