package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.pi) // wont work pi x 
	fmt.Println(math.Pi) // Pi  have to be capitalized letter 
}

/*
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}



var c, python, java bool

func main() {
	var i int =1
	var x float64 = 1.5
	python = true
	fmt.Println(i, c,x, python, java)
}
