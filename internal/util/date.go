package util

import (
	"fmt"
	"log"
	"time"
)

func formatDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d", t.Year(), t.Month())
}

func GetYearMonth(monthsAgo int) string {
	date := time.Now().AddDate(0, -monthsAgo, 0)
	return formatDate(date)
}

func MustBeValid(raw string) time.Time {
	t, err := time.Parse("2006-01-02", raw)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
