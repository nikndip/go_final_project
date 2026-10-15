package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.Format(dateFormat) > now.Format(dateFormat)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("некорректная дата: %w", err)
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", fmt.Errorf("правило повторения не задано")
	}

	var years, days int

	switch parts[0] {
	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("ожидается правило d с интервалом")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval < 1 || interval > 400 {
			return "", fmt.Errorf("интервал должен быть числом от 1 до 400")
		}

		days = interval

	case "y":
		if len(parts) != 1 {
			return "", fmt.Errorf("правило y не принимает параметры")
		}

		years = 1

	case "w":
		if len(parts) != 2 {
			return "", fmt.Errorf("ожидается правило w со списком дней недели")
		}

		return nextWeekDate(now, date, parts[1])

	case "m":
		if len(parts) < 2 || len(parts) > 3 {
			return "", fmt.Errorf("ожидается правило m со списком дней и необязательным списком месяцев")
		}

		months := ""
		if len(parts) == 3 {
			months = parts[2]
		}

		return nextMonthDate(now, date, parts[1], months)

	default:
		return "", fmt.Errorf("неподдерживаемое правило: %s", parts[0])
	}

	for {
		date = date.AddDate(years, 0, days)
		if afterNow(date, now) {
			return date.Format(dateFormat), nil
		}
	}
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	if value := r.FormValue("now"); value != "" {
		parsed, err := time.Parse(dateFormat, value)
		if err != nil {
			http.Error(w, "некорректный параметр now", http.StatusBadRequest)
			return
		}

		now = parsed
	}

	nextDate, err := NextDate(
		now,
		r.FormValue("date"),
		r.FormValue("repeat"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, nextDate)
}
