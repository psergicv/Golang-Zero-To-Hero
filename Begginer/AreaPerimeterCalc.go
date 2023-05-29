package main

import "fmt"

func main() {
	areaRect := CalculateArea(15, 23)
	fmt.Printf("The area of a 15x23 rectangle is: %v \n", areaRect)
	perimRect := CalculatePerimeter(15, 23)
	fmt.Printf("The perimeter of a 15x23 rectangle is: %v \n", perimRect)
}

func CalculateArea(w int, l int) int {
	return w * l
}

func CalculatePerimeter(w int, l int) int {
	return 2 * (l + w)
}
