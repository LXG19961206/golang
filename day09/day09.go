// 今天没有精力学习新的内容,复习之前的东西
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("欢迎开始")
	test2()
}
func test() {
	var time = time.Now()
	rand.Seed(time.Unix())
	var arr [21]int
	var slice = make([]int, 21, 21)
	for i := 0; i <= 20; i++ {
		arr[i] = rand.Intn(20)
		slice[i] = rand.Intn(30)
	}
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func test2() {
	var time = time.Now().Unix()
	rand.Seed(time)
	var slice = make([]int, 20, 20)
	for i := 0; i < 20; i++ {
		slice[i] = rand.Intn(20)
	}
	slice = append(slice, slice...)
	fmt.Println(slice)
}
