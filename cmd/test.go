package main

import (
	"fmt"

	faker "github.com/eykn2z/go_faker"
)

func main() {
	fmt.Println(faker.GetPersons(10))
}
