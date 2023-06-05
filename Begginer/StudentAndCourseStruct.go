// Student and Course Structs: Create two structs, one for a Student and one for a Course. Each student should have fields for
// Name and a slice of Courses they are enrolled in. Each course should have a field for Name and a slice of Students enrolled.
// Implement methods for enrolling a student in a course, and for listing all courses a student is enrolled in, and all students
// enrolled in a course.

package main

import (
	"fmt"
)

type Student struct {
	Name    string
	Courses []*Course
}

type Course struct {
	Name     string
	Students []*Student
}

func (c *Course) addStudentToCourse(student *Student) {
	c.Students = append(c.Students, student)
}

func (s *Student) enrollStudentToCourse(course *Course) {
	s.Courses = append(s.Courses, course)
}

func main() {
	course := &Course{
		Name: "Informatics",
	}
	student := &Student{
		Name: "John",
	}

	course.addStudentToCourse(student)
	student.enrollStudentToCourse(course)

	fmt.Println("Students in Math course:")
	for _, student := range course.Students {
		fmt.Println(student.Name)
	}

	fmt.Println("Courses John enrolled:")
	for _, course := range student.Courses {
		fmt.Println(course.Name)
	}

}
