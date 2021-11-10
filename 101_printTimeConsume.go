package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const (
		width  = 80
		height = 15
	)
	show1(height, width)

	show2(height, width)

	show3(height, width)
}

func show1(height, width int) {
	start := time.Now() // 获取当前时间
	fmt.Print(func() string {
		var str string
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				str += "*"
			}
			str += "\n"
		}
		return str
	}())
	elapsed := time.Since(start) // 获取结束时间
	fmt.Println(elapsed)
}

func show2(height, width int) {
	start := time.Now()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print("-")
		}
		fmt.Print("\n")
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func show3(height, width int) {
	start := time.Now()
	for i := 0; i < height; i++ {
		fmt.Println(strings.Repeat("+", width))
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
