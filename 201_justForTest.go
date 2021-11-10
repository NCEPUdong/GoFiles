package main

import (
	"fmt"
)

func main() {
	var array = [][]int{
		{1, 2, 3, 4, 5, 6},
		{4, 5, 6, 7, 8, 9},
		{7, 8, 9, 1, 2, 3},
	}
	// 其实从array结构可以看出来，array里面直接是有3个元素（切片） len(array)
	// 每个切片里又有6个元素（int） len(array[0])
	fmt.Println(len(array))
	fmt.Println(len(array[1]))
}
