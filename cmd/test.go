package main

import (
	"fmt"

	"github.com/aknfujii/faker"
)

func main() {
	fmt.Println(faker.GetPersons(10))
}
