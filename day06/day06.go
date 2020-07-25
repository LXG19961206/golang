package main

import (
	"fmt"
)

func main() {
	fmt.Println("开始")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d", i, j, i*j)
		}
		fmt.Println()
	}
	// test2()
	// test3()
	// test5()
	// test6()
	// test8()
	// test9()
}
func test2() {
	for i, sum := 1, 0; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			sum++
			if sum%5 == 0 && sum != 0 {
				fmt.Println()
			}
		}
	}
}

func test3() {
	fmt.Println()
	for i := 5; i >= 0; i-- {
		for j := 4; j >= i-1; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}

func test5() {
	for i := 0; true; i++ {
		fmt.Println(i)
		if i%5 == 0 {
			continue
		}
		if i >= 20 {
			break
		}
	}
}

func test6() {
	for x := 1; x <= 9; x++ {
		for y := 0; y <= 9; y++ {
			for z := 0; z <= 9; z++ {
				var temp = 100*x + 10*y + z
				if x*x*x+y*y*y+z*z*z == temp {
					fmt.Printf("%d,是水仙花数", temp)
					fmt.Println()
				}
			}
		}
	}
}
func test8() {
	for i := 2; i <= 100; i++ {
		var isAnswer bool = true
		for j := 2; j <= 100; j++ {
			if j > i {
				break
			}
			if i%j == 0 && i != j {
				isAnswer = false
			}
		}
		if isAnswer == true {
			fmt.Printf("%d是素数", i)
			fmt.Println()
		}
	}
}

func test9() {
LOOP:
	fmt.Printf("1")
	fmt.Printf("2")
	fmt.Printf("3")
	fmt.Printf("4")
	goto LOOP
	fmt.Printf("5")
}
