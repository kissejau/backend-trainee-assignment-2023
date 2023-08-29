package log

import (
	"database/sql"
	"fmt"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetLogs(leftTime, rightTime time.Time) ([]Log, error) {
	var logs []Log

	lt := leftTime.Format("2006-01-02T15:04:05Z")
	rt := rightTime.Format("2006-01-02T15:04:05Z")
	fmt.Printf("lt: %v\n rt: %v\n", lt, rt)
	query := `SELECT id, user_id, slug, type, created_at FROM logs WHERE created_at > $1 AND created_at < $2`

	rows, err := r.db.Query(query, lt, rt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log Log
		err := rows.Scan(&log.Id, &log.UserId, &log.Slug, &log.Type, &log.CreatedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
