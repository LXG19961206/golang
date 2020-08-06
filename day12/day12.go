package main

import "fmt"

func main() {
	test2()
}
func test() {
	map1 := make(map[string]string)
	map1["name"] = "TOM"
	map1["gender"] = "male"
	map1["address"] = "xxx"

	map2 := make(map[string]string)
	map2["name"] = "jenny"
	map2["gender"] = "female"
	map2["address"] = "xxx"

	map3 := map[string]string{"name": "lalla", "age": "12", "adress": "bj"}
	// 如何用一个切片来存储map
	slice := make([]map[string]string, 0, 3)
	slice = append(slice, map1)
	slice = append(slice, map2)
	slice = append(slice, map3)
	for _, item := range slice {
		fmt.Println(item)
	}
}

func test2() {
	map1 := map[string]string{"age": "12", "gender": "0"}
	// map2表示这个map的key是字符串,值还是一个map
	map2 := make(map[string]map[string]string)
	map2["prop"] = map1
	map2["info"] = map[string]string{"地址": "万寿路", "爱好": "唱跳rap"}
	// fmt.Printf("%T", map2)
	fmt.Println(map2)
}
