package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Intn(10)生成[0, 10)之间的整数, 包括0但不包括10
	var num = rand.Intn(10) + 1
	fmt.Print(num, "\n")

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
