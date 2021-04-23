package faker

import (
	"github.com/aknfujii/faker/data"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age int
	Gender string
	Height float64
	Country string
	CreatedAt time.Time
}

func (f *faker) GetPerson(gender string) *Person {
	person := Person{Age: f.Age(), CreatedAt: time.Now()}
	if gender == ""{
		gender = randomChoice(data.Gender)
	}
	if gender == "Female" {
		person.Name = f.FemaleName()
		person.Gender = "Female"
		person.Height= f.FemaleHeight()
	} else if gender == "Male" {
		person.Name = f.MaleName()
		person.Gender = "Male"
		person.Height= f.MaleHeight()
	}
	return &person
}

func (f *faker) GetPersons(length int) []Person{
	var persons = make([]Person, 0)
	for i := 0; i < length; i++ {
		persons = append(persons, *f.GetPerson(""))
	}
	return persons
}

func (f *faker) Age() int {
	return rand.Intn(60) + 10
}

func (f *faker) MaleHeight() float64 {
	return rand.NormFloat64() * 6.6 + 171.7
}

func (f *faker) FemaleHeight() float64 {
	return rand.NormFloat64() * 5.7 + 158.3
}


func (f *faker) PostCode() string {
	// TODO
	return ""
}

func (f *faker) Province() string {
	// TODO
	return ""
}

func (f *faker) Address() string {
	// TODO
	// buildingNumber streetName streetAddress postcode City
	return ""
}
