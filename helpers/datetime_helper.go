package helpers

import "time"

func ConvertStringToDatetime(d string) (time.Time, error) {
	layout := "2006-01-02"

	return time.Parse(layout, d)
}
