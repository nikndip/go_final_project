package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseMonthDays(value string) (map[int]bool, error) {
	days := make(map[int]bool)

	for _, part := range strings.Split(value, ",") {
		day, err := strconv.Atoi(part)
		if err != nil || day == 0 || day < -2 || day > 31 {
			return nil, fmt.Errorf(
				"день месяца должен быть от 1 до 31 либо равен -1 или -2",
			)
		}

		days[day] = true
	}

	return days, nil
}

func parseMonths(value string) ([13]bool, error) {
	var months [13]bool

	if value == "" {
		for month := 1; month <= 12; month++ {
			months[month] = true
		}

		return months, nil
	}

	for _, part := range strings.Split(value, ",") {
		month, err := strconv.Atoi(part)
		if err != nil || month < 1 || month > 12 {
			return months, fmt.Errorf("месяц должен быть числом от 1 до 12")
		}

		months[month] = true
	}

	return months, nil
}

func nextMonthDate(now, date time.Time, daysValue, monthsValue string) (string, error) {
	days, err := parseMonthDays(daysValue)
	if err != nil {
		return "", err
	}

	months, err := parseMonths(monthsValue)
	if err != nil {
		return "", err
	}

	if !afterNow(date, now) {
		year, month, day := now.Date()
		date = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	}

	limit := date.AddDate(400, 0, 0)

	for date.Before(limit) {
		date = date.AddDate(0, 0, 1)

		if !months[date.Month()] {
			continue
		}

		lastDay := time.Date(
			date.Year(), date.Month()+1, 0,
			0, 0, 0, 0, time.UTC,
		).Day()

		day := date.Day()
		if days[day] ||
			(days[-1] && day == lastDay) ||
			(days[-2] && day == lastDay-1) {
			return date.Format(dateFormat), nil
		}
	}

	return "", fmt.Errorf("правило не задаёт существующую дату")
}
