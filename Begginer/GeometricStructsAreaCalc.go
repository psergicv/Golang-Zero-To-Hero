package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func main() {
	rect := Rectangle{
		Width:  10.0,
		Height: 7.5,
	}

	rectArea := rect.Area()
	fmt.Printf("The area of a rectangle with Width = %v and Height = %v is: %v\n", rect.Width, rect.Height, rectArea)

	circ := Circle{
		Radius: 11.4,
	}

	circleArea := circ.Area()
	fmt.Printf("The area of a circle with radius = %v is: %v\n", circ.Radius, circleArea)

}
