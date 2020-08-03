package main

import "fmt"

func main() {
	test5()
}
func test() {
	var map1 = map[string]int{"chinese": 99, "math": 98, "english": 97}
	fmt.Println(map1)
	var map2 map[int]string
	map2 = make(map[int]string)
	map2[2] = "22"
	fmt.Println(map2)
}
func test2() {
	var map1 = make(map[string]int) //首先得是一个不为nil的map
	map1["语文"] = 99
	fmt.Println(map1)
}

func test3() {
	var map1 = map[string]int{"语文": 99, "数学": 99}
	if _, ok := map1["英语"]; ok == true {
		fmt.Println("key值存在")
	} else {
		fmt.Println("key值不存在")
	}
}

func test4() {
	var map1 = map[string]int{"语文": 99, "数学": 90, "英语": 100}
	for key, value := range map1 {
		fmt.Println(map1[key])
		fmt.Println(value)
		fmt.Println(map1[key] == value)
	}
}

func test5() {
	var map1 = map[int]string{1: "孙悟空", 2: "猪八戒", 3: "唐僧", 4: "比克大魔王"}
	for i := 1; i <= 4; i++ {
		fmt.Printf("%d:%s\n", i, map1[i])
	}
}
