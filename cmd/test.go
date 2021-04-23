package main

import (
	"fmt"

	"github.com/aknfujii/faker"
)

func main() {
	faker := faker.GetFaker()
	fmt.Println(faker.GetPersons(10))
}
