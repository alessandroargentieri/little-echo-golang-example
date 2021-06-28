package utils

import "time"

func PointerOfString(s string) *string {
	return &s
}

func PointerOfTime(t time.Time) *time.Time {
	return &t
}
