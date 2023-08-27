package config

import "os"

type Config struct {
	Port string
	Db   DbConfig
}

type DbConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Port: getEnv("LISTEN_PORT", "8080"),
		Db: DbConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			Database: getEnv("DB_NAME", ""),
			Username: getEnv("DB_USER", ""),
			Password: getEnv("DB_PASS", ""),
		},
	}
}

func getEnv(key, defaultVal string) string {
	val, flag := os.LookupEnv(key)
	if !flag {
		return defaultVal
	}
	return val
}
