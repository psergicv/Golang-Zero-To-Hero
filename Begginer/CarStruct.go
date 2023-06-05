// Car Struct: Create a Car struct with fields for Make, Model, and Year. Create a method called honk that prints
// "Beep beep! I am a [Year] [Make] [Model]!" when called. Create several car instances and call the honk method to test your code.

package main

import (
	"fmt"
)

type Car struct {
	Make  string
	Model string
	Year  uint
}

func (c Car) Honk() string {
	return fmt.Sprint("Beep beep! I am a ", c.Year, " ", c.Make, " ", c.Model)
}

func main() {

	renault := Car{
		Make:  "Renault",
		Model: "Megane",
		Year:  2013,
	}

	peugeot := Car{
		Make:  "Peugeot",
		Model: "3008",
		Year:  2018,
	}

	volvo := Car{
		Make:  "Volvo",
		Model: "XC40",
		Year:  2021,
	}

	carList := []Car{renault, peugeot, volvo}

	for _, car := range carList {
		fmt.Println(car.Honk())
	}
}
