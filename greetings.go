package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %s. welcome!", name)
	fmt.Println("Printing from the function:")
	fmt.Println(message)
	return message
}
