package main

import (
	"fmt"
	"math/big"
	// big.Int 超大整数 10e^18
	// big.Float 任意精度浮点类型
	// big.Rat 分数(小数末尾无限长)
)

func main() {

	// 一旦计算数值时使用了big数值, 后续所有变量都必须使用big数值
	lightSpeed := big.NewInt(299792) // km/s
	secondsPerDay := big.NewInt(86400)
	daysPerYear := big.NewInt(365)

	distance := new(big.Int)
	distance.SetString("236000000000000000", 10) // km

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	lightyear := new(big.Int)
	lightyear.Div(days, daysPerYear)

	fmt.Printf("The distance between Andromeda and Earth is\040%v\040lightyears\n", lightyear)
	fmt.Println("Andromeda Galaxy is", 236000000000000000/299792/86400, "light days away")
}
