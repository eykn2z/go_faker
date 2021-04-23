package faker

import (
	"math/rand"
	"time"

	"github.com/aknfujii/faker/data"
)

type Person struct {
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	Gender    string     `json:"gender"`
	BirthDay  CustomDate `json:"birthday,omitempty"`
	Height    float64    `json:"height"`
	Country   string     `json:"country,omitempty"`
	Address   string     `json:"address,omitempty"`
	Duration  `json:"duration,omitempty"`
	CreatedAt CustomDateTime `json:"created_at"`
}

type Duration struct {
	CanContact bool        `json:"can_contact"`
	Weekday    []string    `json:"weekday,omitempty"`
	StartTime  *CustomTime `json:"start_time,omitempty"`
	EndTime    *CustomTime `json:"end_time,omitempty"`
}

func getWeekdayStrArray(weekdayArray []time.Weekday) []string {
	var weekdayStrArray []string
	for _, weekday := range weekdayArray {
		weekdayStrArray = append(weekdayStrArray, weekday.String())
	}
	return weekdayStrArray
}

func getDuration() Duration {
	if rand.Intn(2) == 0 {
		return Duration{CanContact: false}
	} else {
		weekdayStrArray := getWeekdayStrArray([]time.Weekday{1, 2, 3, 4, 5})
		return Duration{CanContact: true, Weekday: weekdayStrArray, StartTime: CustomTime{}.Parse("09:00"), EndTime: CustomTime{}.Parse("17:00")}
	}
}

func GetPerson(gender string) *Person {
	age, birthday := getAgeAndBirthday()
	person := Person{Age: age, BirthDay: birthday, CreatedAt: CustomDateTime{time.Now()}, Duration: getDuration()}
	if gender == "" {
		gender = randomChoice(data.Gender)
	}
	if gender == "Female" {
		person.Name = FemaleName()
		person.Gender = "Female"
		person.Height = FemaleHeight()
	} else if gender == "Male" {
		person.Name = MaleName()
		person.Gender = "Male"
		person.Height = MaleHeight()
	}
	return &person
}

func GetPersons(length int) []Person {
	var persons = make([]Person, 0)
	for i := 0; i < length; i++ {
		persons = append(persons, *GetPerson(""))
	}
	return persons
}

func Age() int {
	return rand.Intn(60) + 10
}

func getAgeAndBirthday() (int, CustomDate) {
	age := Age()
	now := time.Now()
	birthday := CustomDate{now.AddDate(-1*age, 0, 0)}
	return age, birthday
}

func MaleHeight() float64 {
	return round(rand.NormFloat64()*6.6+171.7, 10)
}

func FemaleHeight() float64 {
	return round(rand.NormFloat64()*5.7+158.3, 10)
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
