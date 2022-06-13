package main

import (
	"fmt"

	"github.com/szhao15/greetings/greetings"
)

func main() {
	message := greetings.Hello("Stephen", "Zhao")
	fmt.Println(message)
}
