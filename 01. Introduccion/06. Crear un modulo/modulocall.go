package main

import (
	"fmt"
	"github.com/alexroel/greetings"
)
// para que sirva debemos instalarlo, y si esta en local go mod edit -replace github.com/alexroel/greetings=../greetings
// para que se aplique hacemos go mod tidy

func main() {
	message := greetings.Hello("Alex")
	fmt.Println(message)
}