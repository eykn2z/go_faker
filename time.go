package faker

import "time"

func (f *Faker) Datetime() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}

func (f *Faker) Time() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}

func (f *Faker) Date() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}
