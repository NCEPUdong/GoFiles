package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	pig struct {
		name string
	}

	dog struct {
		name string
	}

	duck struct {
		name string
	}

	parrot struct {
		name string
	}

	gopher struct {
		name string
	}

	mover interface {
		move()
	}

	eater interface {
		eat()
	}
)

var food = []string{"cabbage", "carrot", "nut", "peanut", "watermelon"}

func (p pig) String() string {
	return fmt.Sprint(p.name)
}

func (p pig) move() string {
	return fmt.Sprintln(p.name, "is running!")
}

func (p pig) eat(foodRand int) string {
	return food[foodRand]
}

func (do dog) String() string {
	return fmt.Sprint(do.name)
}

func (d dog) move() string {
	return fmt.Sprintln(d.name, "is barking!")
}

func (d dog) eat(foodRand int) string {
	return food[foodRand]
}

func (du duck) String() string {
	return fmt.Sprint(du.name)
}

func (du duck) move() string {
	return fmt.Sprintln(du.name, "is paddding!")
}

func (du duck) eat(foodRand int) string {
	return food[foodRand]
}

func (p parrot) String() string {
	return fmt.Sprint(p.name)
}

func (p parrot) move() string {
	return fmt.Sprintln(p.name, "is flying")
}

func (p parrot) eat(foodRand int) string {
	return food[foodRand]
}

func (g gopher) String() string {
	return fmt.Sprint(g.name)
}

func (g gopher) move() string {
	return fmt.Sprintln(g.name, "is digging")
}

func (g gopher) eat(foodRand int) string {
	return food[foodRand]
}

func main() {
	animals := []Struct{
		pig{name:boxer}
	}
	for i:=0; i<24; i++ {
		rand.Seed(time.Now().UnixNano())
		foodRand := rand.Intn(6)
		fmt.Println()
	}
}
