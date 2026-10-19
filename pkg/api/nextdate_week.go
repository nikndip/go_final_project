package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func nextWeekDate(now, date time.Time, days string) (string, error) {
	var allowedDays [8]bool

	for _, value := range strings.Split(days, ",") {
		day, err := strconv.Atoi(value)
		if err != nil || day < 1 || day > 7 {
			return "", fmt.Errorf("день недели должен быть числом от 1 до 7")
		}

		allowedDays[day] = true
	}

	for {
		date = date.AddDate(0, 0, 1)

		weekday := int(date.Weekday())
		if weekday == 0 {
			weekday = 7
		}

		if allowedDays[weekday] && afterNow(date, now) {
			return date.Format(dateFormat), nil
		}
	}
}
