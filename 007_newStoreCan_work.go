package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		balance int = 0
		num     int = 0
	)
	fmt.Println("Current blance is $", balance)
	for balance <= 50 {
		time.Sleep(time.Second)
		num = rand.Intn(3)
		balance = addMoney(num, balance)
		fmt.Printf("Current blance is $%4.2f\n", float32(balance)/10)
	}
}

func addMoney(num int, balacnce int) int {
	const (
		Nickel  = 5
		Dime    = 10
		Quarter = 25
	)
	switch num {
	case 0:
		balacnce += Nickel
		fmt.Println("Add a Nickel $", Nickel)
	case 1:
		balacnce += Dime
		fmt.Println("Add a Dime $", Dime)
	case 2:
		balacnce += Quarter
		fmt.Println("Add a Quarter $", Quarter)
	}
	time.Sleep(time.Second)
	return balacnce
}
