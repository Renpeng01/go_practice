package main

import "fmt"

func main() {
	defer fmt.Println("in main1")
	defer fmt.Println("in main2")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
