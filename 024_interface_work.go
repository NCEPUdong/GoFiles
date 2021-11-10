package main

import (
	"fmt"
)

/*
fmt包中已经声明过关于String的接口

type Stringer interface (
	String() string
)

因此在本函数中可以直接使用String方法
*/

type (
	coordinate struct {
		d, m, s float64
		h       rune
	}
)

func (c coordinate) String() string {
	return fmt.Sprint(c.d, "° ", c.m, "' ", c.s, `" `, string(c.h))
}

func main() {
	co1 := coordinate{d: 4, m: 30, s: 0.0, h: 'N'}
	co2 := coordinate{d: 135, m: 54, s: 0.0, h: 'E'}

	fmt.Println("Elysium Planitia is at", co1, co2)
}
