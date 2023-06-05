// Employee Struct: Create an Employee struct that has fields for Name, ID, and Position. Add methods to update any of these fields.
// Also, create a method that prints all the information for an Employee in a nice format.

package main

import (
	"fmt"
)

type Employee struct {
	ID       uint
	Name     string
	Position string
}

func (p *Employee) setNewPosition(newPosition string) string {
	p.Position = "Manager"
	return p.Position
}

func main() {
	person := Employee{
		ID:       1,
		Name:     "George",
		Position: "Clerk",
	}

	fmt.Printf("%v is a %v at Go Inc.\n", person.Name, person.Position)

	person.setNewPosition("Manager")

	fmt.Printf("Now, %v position at Go Inc. is %v\n", person.Name, person.Position)
}
