package date

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow is
func GetNow() time.Time {
	now := time.Now().UTC()
	return now
}

//GetNowString is
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}
