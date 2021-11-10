package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		balance float32 = 0
		num     int     = 0
	)
	fmt.Println("Current blance is", balance)
	for balance <= 20 {
		time.Sleep(time.Second)
		num = rand.Intn(3)
		balance = addMoney(num, balance)
		fmt.Printf("Current blance is %4.2f\n", balance)
	}
}

func addMoney(num int, balacnce float32) float32 {
	const (
		Nickel  = 0.05
		Dime    = 0.10
		Quarter = 0.25
	)
	switch num {
	case 0:
		balacnce += Nickel
		fmt.Println("Add a Nickel", Nickel)
	case 1:
		balacnce += Dime
		fmt.Println("Add a Dime", Dime)
	case 2:
		balacnce += Quarter
		fmt.Println("Add a Quarter", Quarter)
	}
	time.Sleep(time.Second)
	return balacnce
}
