package time

import "time"

func NewTimeZone(name string) {
	ict, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
