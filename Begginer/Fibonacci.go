package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Please enter a number: ")
	userInput := 0
	fmt.Scan(&userInput)
	sequence, err := Fibo(userInput)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The Fibonacci sequence of %v is: %v", userInput, sequence)
	}
}

func Fibo(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Please, enter a positive number")
	} else if n == 1 {
		return 0, nil
	} else if n == 2 {
		return 1, nil
	} else {
		val1, err1 := Fibo(n - 1)
		val2, err2 := Fibo(n - 2)
		if err1 != nil {
			return 0, err1
		}
		if err2 != nil {
			return 0, err2
		}
		return val1 + val2, nil
	}
}
