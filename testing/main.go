package main

import (
	"fmt"

	"github.com/szhao15/greetings"
)


func main() {
	//message := "Stephen"
	// fmt.Printf("Hi there %v. We hope you have a great day.\n", message)
	// fmt.Sprintf("Hi there %v. We also hope you have a good time.", message)
	// fmt.Println("Good bye.", message)

	welcome := greetings.Hello("Stephen")
	fmt.Println(welcome)
}