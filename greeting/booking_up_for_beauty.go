package greeting

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	if date == "" {
		panic("Please implement the Schedule function")
	}

	// 2006 年 1 月 2 日 下午 3 點 4 分 5 秒
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Please implement the Schedule function")
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	if date == "" {
		panic("Please implement the HasPassed function")
	}

	layout := "January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		panic("Please implement the HasPassed function")
	}

	currentTime := time.Now()
	return parsedTime.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	if date == "" {
		panic("Please implement the IsAfternoonAppointment function")
	}

	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		panic("Please implement the IsAfternoonAppointment function")
	}

	hour := parsedTime.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	if date == "" {
		panic("Please implement the Description function")
	}

	parsedTime := Schedule(date)
	outputLayout := "Monday, January 2, 2006, at 15:04"
	return "You have an appointment on " + parsedTime.Format(outputLayout) + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
