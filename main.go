package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/josipklaric/go-playground/greeter"
)

func main() {
	fmt.Println("Hello there!")

	fmt.Println(quote.Hello())

	message, err := greeter.Greet("")

	if err != nil {
		message = fmt.Sprintf("Greeter module returned and error: %v", err.Error())
	}

	fmt.Println(message)
}
