package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 不能给文件起名为*_test.go, 这是golang专属的文件名, 用于测试
func main() {
	var num = rand.Intn(100)
	for {
		rand.Seed(time.Now().UnixNano())
		guess := rand.Intn(100)
		fmt.Print("origin: ", num, "\nguess: ", guess)
		if guess >= num {
			fmt.Print("\nYou Win!")
			break
		} else {
			fmt.Println("\nThe", guess, "is smaller than", num, ".\nWaiting for next guess...")
			time.Sleep(time.Second * 2)
		}
	}
}
