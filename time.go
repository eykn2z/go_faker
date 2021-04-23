package faker

import "time"

func (f *faker) Datetime() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}

func (f *faker) Time() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}

func (f *faker) Date() time.Time {
	// TODO start,end(option) =>random
	return time.Now()
}
