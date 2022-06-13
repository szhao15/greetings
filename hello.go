package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprint("Hi, %v. welcome!", name)
	return message
}
