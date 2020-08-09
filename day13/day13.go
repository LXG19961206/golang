package main

import "fmt"

func main() {
	test()
	test2()
}
func test() {
	for i, string1 := 0, "helloworld"; i < len(string1); i++ {
		fmt.Printf("%c\n", string1[i])
	}

	for i, string1 := 0, "中国"; i < len(string1); i++ {
		fmt.Printf("%c\n", string1[i])
	}
}

func test2() {
	var string1 = "abcde"
	var slice1 = []byte(string1)
	fmt.Println(slice1)
}
