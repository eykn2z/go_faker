package faker

import (
	"github.com/eykn2z/go_faker/data"
)

func Name() string {
	firstName, _ := FirstName()
	return firstName + " " + LastName()
}

func FemaleName() string {
	return FemaleFirstName() + " " + LastName()
}

func MaleName() string {
	return MaleFirstName() + " " + LastName()
}

func FirstName() (string, string) {
	firstNameArray := [][]string{
		data.FirstNameFemaleArray,
		data.FirstNameMaleArray,
	}
	choicedIndex := randomIndex(len(firstNameArray))
	gender := choicedIndex
	return randomChoice(firstNameArray[choicedIndex]), data.Gender[gender]
}

func FemaleFirstName() string {
	return randomChoice(data.FirstNameFemaleArray)
}

func MaleFirstName() string {
	return randomChoice(data.FirstNameMaleArray)
}

func LastName() string {
	return randomChoice(data.LastNameArray)
}

func Gender() string {
	return randomChoice(data.Gender)
}
