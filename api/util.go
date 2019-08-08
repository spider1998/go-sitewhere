package api

import (
	"strings"
	"time"
)

type Time time.Time

func (t Time) Format() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.Format()), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format() + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	v, err := time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(string(b), "\""), time.Local)
	if err != nil {
		return err
	}
	*t = Time(v)
	return nil
}

func ParseTime(timeString string) (t Time, err error) {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return
	}

	t = Time(parsedTime)

	return
}
