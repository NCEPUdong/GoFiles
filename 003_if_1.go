package main

import "fmt"

func main() {
	var year = 2103
	var leap = year%400 == 0 || (year == 0 && year%100 != 0)

	fmt.Print("The year is ", year, ", should you leap?\n")

	if leap {
		fmt.Println("Look befor you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}
