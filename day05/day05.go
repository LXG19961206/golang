package main

import "fmt"

func main() {
	// num := 20
	// if num > 16 {
	// 	fmt.Printf("num大于16")
	// } else {
	// 	fmt.Printf("num小于16")
	// }
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	test10()
}

func test1() {
	var num int
	fmt.Print("请输入一个数字")
	fmt.Scanln(&num)
	if num > 60 {
		fmt.Print("及格")
	} else {
		fmt.Print("没及格")
	}
}

func test2() {
	var num int
	fmt.Println("请输入一串数字")
	fmt.Scanln(&num)
	if num < 60 {
		fmt.Println("不及格")
	} else if num < 70 {
		fmt.Println("良好")
	} else if num < 80 {
		fmt.Println("优秀")
	} else if num < 90 {
		fmt.Println("nice")
	} else {
		fmt.Println("极佳")
	}
}

func test3() {
	fmt.Println("--------------------------")
	if num := 4; num > 0 {
		fmt.Println("正数")
		fmt.Println(num)
	} else {
		fmt.Println("负数")
	}
	fmt.Println(111)
}

func test4() {
	var num int
	fmt.Println("please input a number")
	fmt.Scanln(&num)
	switch num {
	case 6:
		fmt.Println("不及格")
		break
	case 7:
		fmt.Println("及格")
		break
	case 8:
		fmt.Println("还可以")
	case 9:
		fmt.Println("良好")
	case 10:
		fmt.Println("优秀")
	default:
		fmt.Println("未知的分数")
	}
}

func test5() {
	var (
		num1 int
		num2 int
		opa  string
	)
	fmt.Println("请输入两个数,一个运算符号")
	fmt.Scanln(&num1, &num2, &opa)
	switch opa {
	case "+":
		fmt.Print(num1 + num2)
	case "-":
		fmt.Print(num1 - num2)
	case "*":
		fmt.Print(num1 * num2)
	case "/":
		fmt.Print(num1 / num2)
	}
}

func test6() {
	var num int
	fmt.Println("请输入一个分数")
	fmt.Scanln(&num)
	switch {
	case num < 60:
		fmt.Println("没及格")
	case num < 70:
		fmt.Println("刚及格")
	case num < 80:
		fmt.Println("还可以")
	case num < 90:
		fmt.Println("良好")
	case num < 100:
		fmt.Println("优秀")
	default:
		fmt.Println("满分")
	}
}

func test7() {
	var month int
	fmt.Println("请输入1-12几个数")
	fmt.Scanln(&month)
	switch month {
	case 1, 2, 3:
		fmt.Println("一季度")
	case 4, 5, 6:
		fmt.Println("二季度")
	case 7, 8, 9:
		fmt.Println("三季度")
	case 10, 11, 12:
		fmt.Println("四季度")
	}
}

func test10() {
	n := 1
	switch n {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	}
}
