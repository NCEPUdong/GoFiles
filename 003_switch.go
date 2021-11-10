package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("There is a cavern entrance here and a path to the ease.")

	var commandList = [5]string{
		"go inside",
		"go east",
		"enter cave",
		"read sign",
		"others",
	}
	rand.Seed(time.Now().UnixNano())
	var command = commandList[rand.Intn(5)]
	fmt.Print("The command chosen this time is ", command, ".\n")

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}
