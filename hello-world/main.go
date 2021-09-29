package main

import (
	"fmt"
	"math"
)

func main() {
	//Basic output
	fmt.Println("Hello World")

	//Initialize variables
	var age int = 28
	fmt.Println("My age is", age)

	var width, heigth int = 100, 50
	fmt.Println("My width is", width, "and my height is", heigth)

	var (
		otherAge    = 29
		otherName   = "Sam"
		otherHeight = 190
	)
	fmt.Println("My other age is", otherAge, "and other name is", otherName, "and other height is", otherHeight)

	//Short hand declaration
	count := 10
	fmt.Println("Count = ", count)

	otherCount, anotherAge := 10, 25
	fmt.Println("Another count =", otherCount, "and Another age:", anotherAge)

	//Calculate variables
	a, b := 10.4, 20.5
	c := math.Min(a, b)
	fmt.Println("The minimum value is:", c)
}
