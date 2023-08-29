package log

import (
	"fmt"
	"time"
)

type Log struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Slug      string    `json:"slug"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

func (log *Log) ToArray() []string {
	arr := []string{}
	arr = append(arr, fmt.Sprintf("идентификатор пользователя %v", log.Id))
	arr = append(arr, log.Slug)
	arr = append(arr, fmt.Sprintf("операция %v", log.Type))
	arr = append(arr, log.CreatedAt.String())
	return arr
}
