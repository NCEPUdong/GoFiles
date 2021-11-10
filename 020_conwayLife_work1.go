package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Universe [][]bool
	status   map[bool]int
	cellList []bool
)

func NewUniverse(height, width int) Universe {
	newU := make(Universe, 15)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newU[i] = append(newU[i], false)
		}
	}
	return newU
}

func (u Universe) Show() {
	fmt.Print(func() string {
		var str string
		for i := 0; i < len(u); i++ {
			for j := 0; j < len(u[0]); j++ {
				if u[i][j] == false {
					str += " "
				} else {
					str += "*"
				}
			}
			str += "\n"
		}
		return str
	}())
}

func (u Universe) Seed() Universe {
	var (
		randX int
		randY int
	)
	for i := 0; i < len(u)*len(u[0])/4; i++ {
		randX = rand.Intn(len(u))
		randY = rand.Intn(len(u[0]))
		u[randX][randY] = true
	}
	return u
}

func (u Universe) Neighbors(x, y int) int {
	var (
		cellBlock  cellList
		cellStatus = make(status, 2)
	)
	if x == 0 {
		if y == 0 {
			cellBlock = cellList{
				u[x-1+len(u)][y-1+len(u[1])], u[x][y-1+len(u[1])], u[x+1][y-1+len(u[1])],
				u[x-1+len(u)][y], u[x+1][y],
				u[x-1+len(u)][y+1], u[x][y+1], u[x+1][y+1],
			}
		} else if y == len(u[1]) {
			cellBlock = cellList{
				u[x-1+len(u)][y-1], u[x][y-1], u[x+1][y-1],
				u[x-1+len(u)][y], u[x+1][y],
				u[x-1+len(u)][y+1-len(u[1])], u[x][y+1-len(u[1])], u[x+1][y+1-len(u[1])],
			}
		} else {
			cellBlock = cellList{
				u[x-1+len(u)][y-1], u[x][y-1], u[x+1][y-1],
				u[x-1+len(u)][y], u[x+1][y],
				u[x-1+len(u)][y+1], u[x][y+1], u[x+1][y+1],
			}
		}
	} else if x == len(u) {
		if y == 0 {
			cellBlock = cellList{
				u[x-1][y-1+len(u[1])], u[x][y-1+len(u[1])], u[x+1-len(u)][y-1+len(u[1])],
				u[x-1][y], u[x+1-len(u)][y],
				u[x-1][y+1], u[x][y+1], u[x+1-len(u)][y+1],
			}
		} else if y == len(u[1]) {
			cellBlock = cellList{
				u[x-1][y-1], u[x][y-1], u[x+1-len(u)][y-1],
				u[x-1][y], u[x+1-len(u)][y],
				u[x-1][y+1-len(u[1])], u[x][y+1-len(u[1])], u[x+1-len(u)][y+1-len(u[1])],
			}
		} else {
			cellBlock = cellList{
				u[x-1][y-1], u[x][y-1], u[x+1-len(u)][y-1],
				u[x-1][y], u[x+1-len(u)][y],
				u[x-1][y+1], u[x][y+1], u[x+1-len(u)][y+1],
			}
		}
	} else {
		if y == 0 {
			cellBlock = cellList{
				u[x-1][y-1+len(u[1])], u[x][y-1+len(u[1])], u[x+1][y-1+len(u[1])],
				u[x-1][y], u[x+1][y],
				u[x-1][y+1], u[x][y+1], u[x+1][y+1],
			}
		} else if y == len(u[1]) {
			cellBlock = cellList{
				u[x-1][y-1], u[x][y-1], u[x+1][y-1],
				u[x-1][y], u[x+1][y],
				u[x-1][y+1-len(u[1])], u[x][y+1-len(u[1])], u[x+1][y+1-len(u[1])],
			}
		} else {
			cellBlock = cellList{
				u[x-1][y-1], u[x][y-1], u[x+1][y-1],
				u[x-1][y], u[x+1][y],
				u[x-1][y+1], u[x][y+1], u[x+1][y+1],
			}
		}
	}
	// cellBlock = cellList{
	// 	u[x-1][y-1], u[x][y-1], u[x+1][y],
	// 	u[x-1][y], u[x+1][y],
	// 	u[x-1][y+1], u[x][y+1], u[x+1][y+1],
	// }

	for _, cell := range cellBlock {
		cellStatus[cell]++
	}
	return cellStatus[true]
}

func (u Universe) Alive(x, y int) bool {
	if u[x][y] == false {
		if u.Neighbors(x, y) == 3 {
			return true
		}
		return false
	} else {
		if u.Neighbors(x, y) < 2 {
			return false
		} else if u.Neighbors(x, y) == 2 || u.Neighbors(x, y) == 3 {
			return true
		} else {
			return false
		}
	}
}

func (u Universe) Next(x, y int) bool {
	if u[x][y] == true && (u.Neighbors(x, y) == 2 || u.Neighbors(x, y) == 3) {
		return true
	}
	return false
}

func Step(a, b Universe) Universe {
	for x := 0; x < len(a)-1; x++ {
		for y := 0; y < len(a[0])-1; y++ {
			b[x][y] = a.Next(x, y)
		}
	}
	a, b = b, a
	// fmt.Println("\x0c")
	// "\x0c" equals to "\f"
	fmt.Println("\f")
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const (
		width  = 80
		height = 15
	)

	newWorld1 := NewUniverse(height, width)
	newWorld2 := NewUniverse(height, width)
	newWorld1.Seed().Show()

	for i := 1; i < 5; i++ {
		fmt.Println("The ", i, " step")
		newWorld1 = Step(newWorld1, newWorld2) // a, b = b, a
		time.Sleep(time.Second * 2)
		newWorld1.Show()
	}

}
