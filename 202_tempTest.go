package main

import (
	"fmt"
)

type person struct {
	int age
}

func main() {
	var p *person = nil
	fmt.Println(p.age)
}
