package main

import (
	"day1/mylib"
	"fmt"
)

func main() {
	var arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	mylib.BubbleSort_(arr, len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
