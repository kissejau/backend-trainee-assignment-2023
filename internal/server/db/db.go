package db

import (
	"database/sql"
	"fmt"

	"github.com/kissejau/backend-trainee-assignment-2023/internal/server/config"
)

func NewDb(cfg config.DbConfig) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := CheckConn(db); err != nil {
		return nil, err
	}
	return db, nil
}

func CheckConn(db *sql.DB) error {
	return db.Ping()
}
