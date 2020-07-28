package main

import "fmt"

func main() {
	test()
	sort()
	slice4()
}
func test() {
	var arr = [4]int{1, 2, 3, 4}
	for i, item := range arr {
		fmt.Printf("%d是第%d个", item, i+1)
		fmt.Println()
	}
}

func sort() {
	var arr = [7]int{23, 42, 1, 25, 15, 88, 12}
	for i := 0; i < len(arr)-1; i++ {
		for k := 0; k < len(arr)-i-1; k++ {
			if arr[k] > arr[k+1] {
				var temp = arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

func arr() {
	var arr = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {2, 3, 4, 6}}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
func slice() {
	var slice = make([]int, 3, 8)
	slice[1], slice[2], slice[0] = 1, 2, 3
	fmt.Println(slice)
	fmt.Printf("长度:%d,容量:%d", len(slice), cap(slice))
}
func slice2() {
	var slice = make([]int, 3, 8)
	slice[0], slice[1], slice[2] = 1, 2, 3
	slice = append(slice, 1, 2)
	fmt.Println(slice)
	fmt.Printf("长度:%d,容量:%d", len(slice), cap(slice))
}

func slice3() {
	var slice = make([]int, 3, 8)
	slice[0], slice[1], slice[2] = 1, 2, 3
	var slice2 = []int{5, 6, 7, 8}
	slice = append(slice, slice2...)
	fmt.Println(slice)
	fmt.Printf("长度:%d,容量:%d", len(slice), cap(slice))
}

func slice4() {
	var slice = []int{1, 2, 3, 4}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for i, value := range slice {
		fmt.Printf("序号:%d,值:%d", i, value)
	}
}
