package faker

import (
	"github.com/aknfujii/faker/data"
)

func (f *faker) Name() string {
	firstName, _ := f.FirstName()
	return firstName + " " + f.LastName()
}

func (f *faker) FemaleName() string {
	return f.FemaleFirstName() + " " + f.LastName()
}

func (f *faker) MaleName() string {
	return f.MaleFirstName() + " " + f.LastName()
}


func (f *faker) FirstName() (string, int) {
	firstNameArray := [][]string{
		data.FirstNameFemaleArray,
		data.FirstNameMaleArray,
	}
	choicedIndex := randomChoiceIndexNested(firstNameArray)
	gender := choicedIndex
	return randomChoice(firstNameArray[choicedIndex]), gender
}

func (f *faker) FemaleFirstName() string {
	return randomChoice(data.FirstNameFemaleArray)
}

func (f *faker) MaleFirstName() string {
	return randomChoice(data.FirstNameMaleArray)
}


func (f *faker) LastName() string {
	return randomChoice(data.LastNameArray)
}

func (f *faker) Gender() string {
	return randomChoice(data.Gender)
}
