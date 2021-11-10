package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
fmt包中已经声明过关于String的接口

type Stringer interface (
	String() string
)

type Json.Marshaler interface {
	MarshalJSON() ([]byte, error)
}
因此在本函数中可以直接使用String方法
*/

type (
	coordinate struct {
		d float64
		m float64
		s float64
		h rune
	}

	jsonCo struct {
		Decimal    float64 `json: "decimal"`
		Dms        string  `json: "dms"`
		Degrees    float64 `json: "degrees"`
		Minutes    float64 `json: "minutes"`
		Seconds    float64 `json: "second"`
		Hemisphere rune    `json: "hemisphere"`
	}
)

// 为coordinate定义了一个Stringer接口的String方法，因此在fmt.Print时，会直接使用定义的String方法进行输出
func (c coordinate) String() string {
	return fmt.Sprint(c.d, "° ", c.m, "' ", c.s, `" `, string(c.h))
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	jsonCo_1 := jsonCo{
		Decimal:    c.decimal(),
		Dms:        c.String(),
		Degrees:    c.d,
		Minutes:    c.m,
		Seconds:    c.s,
		Hemisphere: c.h,
	}
	jsonCo_1_Output, err := json.MarshalIndent(jsonCo_1, "", "\t")
	return jsonCo_1_Output, err
}

// transfer rad to decimal
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	co1 := coordinate{d: 4, m: 30, s: 0.0, h: 'N'}
	co2 := coordinate{d: 135, m: 54, s: 0.0, h: 'E'}

	fmt.Println("Elysium Planitia is at", co1, co2)

	co1_bytes, err1 := co1.MarshalJSON()
	co2_bytes, err2 := co2.MarshalJSON()

	exitOnError(err1)
	fmt.Println(string(co1_bytes))

	exitOnError(err2)
	fmt.Println(string(co2_bytes))
}
