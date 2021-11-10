package main

import (
	"fmt"
)

func main() {
	var (
		slice     = make([]int, 0, 2)
		i     int = 0
		cap1  int
		cap2  int
	)
	for i <= 128 {
		cap1 = cap(slice)
		slice = append(slice, i)
		i++
		cap2 = cap(slice)
		if cap1 != cap2 {
			fmt.Println(cap2)
		}
	}
}
