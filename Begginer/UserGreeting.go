package main

import "fmt"

func main() {

	var userName string
	fmt.Scan(&userName)
	fmt.Println(Greeting(userName))

}

func Greeting(name string) string {
	greet_message := fmt.Sprintf("Welcome to Golang, %s!", name)

	return greet_message
}
