package main

import (
	"fmt"
	"math"
)

func main() {
	//Basic output
	fmt.Println("Hello World")

	//Variables in Go
	variables()

	//Types in Go
	types()
}

func variables() {
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

func types() {
	//Booleans
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)

	//%T shows the type
	f1, f2 := 2.45, 23.3
	fmt.Println("Type of f1 %T f2 %T\n", f1, f2)

	//Strings
	first := "Santiago"
	last := "Cortez"
	name := first + " " + last
	fmt.Println("My name is", name)

	//Type conversion
	i := 22
	j := 92.2
	sum := i + int(j)
	fmt.Println("Sum is:", sum)
}
