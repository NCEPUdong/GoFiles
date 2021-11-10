package main

import (
	"fmt"
)

type planets []string

func main() {
	planetList := planets{"Mars", "Uranus", "Neptune"}
	fmt.Println(planetList.terraform("new"))
}

// 函数内对切片的修改会导致外部切片的变化， 因此建议新建一个切片存放修改后的切片
func (p planets) terraform(prefix string) planets {
	// make函数可以用于预分配slice， 第一个参数为切片类型， 第二个为长度， 第三个为容量
	// 初始长度会占位， 因此空切片使用0作为初始长度。 不写容量时会默认为10
	newPlanetList := make(planets, 0, 5)
	for _, planet := range p {
		planet = prefix + " " + planet
		newPlanetList = append(newPlanetList, planet)
	}
	return newPlanetList
}
