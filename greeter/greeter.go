package greeter

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error) {
	if name == "" {
		return name, errors.New("name not specified")
	}
	message := fmt.Sprintf("Hello %v", name)
	return message, nil
}

func main() {
	fmt.Println(Greet("Josip"))
}
