package db

import "time"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	query := `
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (?, ?, ?, ?)
	`

	result, err := db.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func Tasks(limit int, search string) ([]*Task, error) {
	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler
	`

	var args []any

	if search != "" {
		if date, err := time.Parse("02.01.2006", search); err == nil {
			query += " WHERE date = ?"
			args = append(args, date.Format("20060102"))
		} else {
			query += " WHERE title LIKE ? OR comment LIKE ?"
			pattern := "%" + search + "%"
			args = append(args, pattern, pattern)
		}
	}

	query += " ORDER BY date LIMIT ?"
	args = append(args, limit)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)

	for rows.Next() {
		var task Task

		if err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
