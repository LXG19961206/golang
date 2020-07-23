package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("请输入两个值")
	fmt.Scanln(&x, &y)
	fmt.Printf("x的值是 %d,y的值是%d\n", x, y)
	fmt.Println(2 == 3)
	fmt.Print(4 << 2)
}
