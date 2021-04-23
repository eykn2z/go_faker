package main

import (
	"fmt"

	"github.com/aknfujii/faker"
)

func main() {
	faker := faker.Faker{}
	fmt.Println(faker.GetPersons(10))
}
