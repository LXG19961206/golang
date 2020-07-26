package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	var num = rand.Intn(10)
	// 	fmt.Println(num)
	// }
	// test()
	// test2()
	// test3()
	// test4()
	// test5()
	test7()
}

func test() {
	rand.Seed(1)
	var num = rand.Intn(10)
	fmt.Print(num)
}

func test2() {
	var time1 = time.Now()
	fmt.Printf("%T\n", time1)
	//时间戳:指定时间,距离1970计算机元年之间时间的差值
	timeStamp1 := time1.Unix()
	fmt.Println(timeStamp1)
}

func test3() {
	rand.Seed(time.Now().UnixNano()) //给种子数设置为时间戳
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(20))
	}
}

func test4() {
	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	fmt.Println(arr)
	var arr2 = [...]int{12, 3, 4, 5, 6, 7, 3}
	fmt.Println(arr2)
	arr2[len(arr2)-1] = 20
	fmt.Println(arr2)
}
func test5() {
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func test6() {
	var arr = [6]int{3, 4, 5, 7, 4, 23}
	for i, item := range arr {
		fmt.Printf("第%d个,是%d", i, item)
		fmt.Println()
	}
}

func test7() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Print(arr2 == arr)
	arr[0] = 20
	fmt.Print(arr, arr2)
}
