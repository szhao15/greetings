package main

import (
	"fmt"

	"github.com/szhao15/greetings/hello"
)

func main() {
	message := hello.Hello("Stephen")
	fmt.Println(message)
}
