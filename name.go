package faker

import (
	"github.com/aknfujii/go_exercise/faker/data"
)

func (f *Faker) Name() string {
	firstName, _ := f.FirstName()
	return firstName + " " + f.LastName()
}

func (f *Faker) FemaleName() string {
	return f.FemaleFirstName() + " " + f.LastName()
}

func (f *Faker) MaleName() string {
	return f.MaleFirstName() + " " + f.LastName()
}


func (f *Faker) FirstName() (string, int) {
	firstNameArray := [][]string{
		data.FirstNameFemaleArray,
		data.FirstNameMaleArray,
	}
	choicedIndex := randomChoiceIndexNested(firstNameArray)
	gender := choicedIndex
	return randomChoice(firstNameArray[choicedIndex]), gender
}

func (f *Faker) FemaleFirstName() string {
	return randomChoice(data.FirstNameFemaleArray)
}

func (f *Faker) MaleFirstName() string {
	return randomChoice(data.FirstNameMaleArray)
}


func (f *Faker) LastName() string {
	return randomChoice(data.LastNameArray)
}

func (f *Faker) Gender() string {
	return randomChoice(data.Gender)
}
