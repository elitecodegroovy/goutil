package goutil

import (
	"fmt"
	"math"
	"regexp"
	"time"
)

const (
	LayoutDateISO       = "2006-01-02"
	LayoutTimeISO       = "2006-01-02 15:04:05"
	LayoutTimeNumberISO = "20060102150405"
)

// StringsFallback2 returns the first of two not empty strings.
func StringsFallback2(val1 string, val2 string) string {
	return stringsFallback(val1, val2)
}

// StringsFallback3 returns the first of three not empty strings.
func StringsFallback3(val1 string, val2 string, val3 string) string {
	return stringsFallback(val1, val2, val3)
}

func stringsFallback(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// SplitString splits a string by commas or empty spaces.
func SplitString(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	return regexp.MustCompile("[, ]+").Split(str, -1)
}

// GetAgeString returns a string representing certain time from years to minutes.
func GetAgeString(t time.Time) string {
	if t.IsZero() {
		return "?"
	}

	sinceNow := time.Since(t)
	minutes := sinceNow.Minutes()
	years := int(math.Floor(minutes / 525600))
	months := int(math.Floor(minutes / 43800))
	days := int(math.Floor(minutes / 1440))
	hours := int(math.Floor(minutes / 60))

	if years > 0 {
		return fmt.Sprintf("%dy", years)
	}
	if months > 0 {
		return fmt.Sprintf("%dM", months)
	}
	if days > 0 {
		return fmt.Sprintf("%dd", days)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	}
	if int(minutes) > 0 {
		return fmt.Sprintf("%dm", int(minutes))
	}

	return "< 1m"
}

func GetCurrentDateISOStrTime() string {
	t := time.Now()
	return t.Format(LayoutDateISO)
}

func GetISOStrDateTime(t time.Time) string {
	return t.Format(LayoutDateISO)
}

func GetCurrentTimeISOStrTime() string {
	t := time.Now()
	return t.Format(LayoutTimeISO)
}

func GetISOStrTime(t time.Time) string {
	return t.Format(LayoutTimeISO)
}

func GetCurrentTimeNumberISOStrTime() string {
	t := time.Now()
	return t.Format(LayoutTimeNumberISO)
}

func GetISOStrTimeNumber(t time.Time) string {
	return t.Format(LayoutTimeNumberISO)
}
