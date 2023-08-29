package log

import "time"

type Repository interface {
	GetLogs(leftTime, rightTime time.Time) ([]Log, error)
}
