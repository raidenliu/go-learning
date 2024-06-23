package main

import (
	"fmt"
	"strconv"
)

func main() {

	//1.变量不赋值时，有初始
	//
	//值
	//2.变量不使用会报错
	//3.变量不可以重复声明
	var i int = 5
	fmt.Println(i)
	var str = "aaa"
	fmt.Println(str)
	a := 90
	a = 5
	fmt.Println(a)

	chInt := make(chan int, 2)

	chInt <- 2
	chInt <- 5

	b := <-chInt
	fmt.Println(b)

	var fn func(i int) string

	fn = func(i int) string {
		return strconv.Itoa(i)
	}
	fmt.Println(fn(7))

}

func test() {

}
