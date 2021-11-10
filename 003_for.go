package main

import (
	"fmt"
	"time"
)

// for 后面如果没有跟循环条件, 就是无限循环
func main() {
	var count1, count2 = 3, 4
	for count1 > 0 {
		fmt.Println(count1)
		time.Sleep(time.Second)
		count1--
	}
	fmt.Println("Lift off!")

	for {
		if count2 < 0 {
			break
		}
		fmt.Println(count2)
		time.Sleep(time.Second)
		count2--
	}
	print("Lift off!")
}
