package greetings

import "fmt"

func Hello(name, last string) string {
	message := fmt.Sprintf("Hi, %s %s. welcome!", name, last)
	fmt.Println("Printing from the function:")
	fmt.Println(message)
	return message
}
