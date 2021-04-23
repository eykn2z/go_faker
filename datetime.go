package faker

import "time"

type CustomTime struct {
	time.Time
}

func (t CustomTime) format() string {
	return t.Time.Format("15:04")
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

func (t CustomTime) Parse(strTime string) CustomTime {
	time, _ := time.Parse("15:04", strTime)
	return CustomTime{time}
}

type CustomDate struct {
	time.Time
}

func (t CustomDate) format() string {
	return t.Time.Format("2006-01-02")
}

func (t CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

func (t CustomDate) Parse(strDate string) CustomDate {
	date, _ := time.Parse("2006-01-02", strDate)
	return CustomDate{date}
}

type CustomDateTime struct {
	time.Time
}

func (t CustomDateTime) format() string {
	return t.Time.Format("2006-01-02 15:04")
}

func (t CustomDateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.format() + `"`), nil
}

func (t CustomDateTime) Parse(strDateTime string) CustomDateTime {
	datetime, _ := time.Parse("2006-01-02 15:04", strDateTime)
	return CustomDateTime{datetime}
}
