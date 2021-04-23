package faker

import (
	"github.com/aknfujii/faker/data"
	"math"
	"math/rand"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Height float64 `json:"height"`
	Country string `json:"country,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func GetPerson(gender string) *Person {
	person := Person{Age: Age(), CreatedAt: time.Now()}
	if gender == ""{
		gender = randomChoice(data.Gender)
	}
	if gender == "Female" {
		person.Name = FemaleName()
		person.Gender = "Female"
		person.Height= FemaleHeight()
	} else if gender == "Male" {
		person.Name = MaleName()
		person.Gender = "Male"
		person.Height= MaleHeight()
	}
	return &person
}

func GetPersons(length int) []Person{
	var persons = make([]Person, 0)
	for i := 0; i < length; i++ {
		persons = append(persons, *GetPerson(""))
	}
	return persons
}

func Age() int {
	return rand.Intn(60) + 10
}

func MaleHeight() float64 {
	return math.Round((rand.NormFloat64() * 6.6 + 171.7) *10) / 10
}

func FemaleHeight() float64 {
	return math.Round((rand.NormFloat64() * 5.7 + 158.3) *10) / 10
}


func PostCode() string {
	// TODO
	return ""
}

func Province() string {
	// TODO
	return ""
}

func Address() string {
	// TODO
	// buildingNumber streetName streetAddress postcode City
	return ""
}
