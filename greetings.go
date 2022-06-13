package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Printf("Hi, %v. welcome!", name)
	return message
}
