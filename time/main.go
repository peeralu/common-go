package time

import "time"

func InitTimeZone(name string) {
	ict, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
