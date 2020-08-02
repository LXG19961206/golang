// 冷落了好久天 golan ,首先复习下之前的东西
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("开始")
	test6()
}
func test() {
	rand.Seed(time.Now().Unix())
	slice := make([]int, 10, 10)
	for i, _ := range slice {
		slice[i] = rand.Intn(20)
	}
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func test2() {
	var slice = []int{1, 2, 3}
	var slice2 = slice
	slice2[0] = 2
	fmt.Println(slice)                // 2,2,3
	fmt.Println(slice2)               // 2,2,3
	fmt.Printf("%p,%p", slice, slice) // 就引用类型而言,%p可取其在内存中的地址
}

func test3() {
	var slice = make([]int, 3, 3)
	fmt.Printf("%p", slice)
	fmt.Println()
	slice = []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Printf("%d,%d", len(slice), cap(slice))
	fmt.Printf("%p", slice)
}

func test4() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var slice = arr[:10]
	fmt.Println(slice)
	var slice2 = arr[0:10]
	fmt.Println(slice2)
	var slice3 = arr[:]
	fmt.Println(slice3)
	fmt.Printf("%d\n,%d\n", len(slice), cap(slice))
}

func test5() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var slice = arr[3:7]
	fmt.Printf("%d,%d", len(slice), cap(slice))
	var slice2 = arr[3:10]
	fmt.Printf("%d,%d", len(slice2), cap(slice2))
}

func test6() {
	var slice = []int{1, 2, 3, 4}
	var slice2 = make([]int, 5, 5)
	slice2[0] = 0
	var slice3 = make([]int, 4)
	copy(slice2, slice)
	copy(slice3, slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
