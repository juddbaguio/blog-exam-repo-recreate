package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   Database `json:"DB"`
	Port int
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	Driver   string `json:"driver"`
}

// ReadConfig function loads the environment variables to the system
func ReadConfig(file_path string) (*Config, error) {
	err := godotenv.Load(file_path)
	if err != nil {
		return nil, fmt.Errorf("error loading this file %s", err.Error())
	}
	return &Config{}, nil
}

func (config *Config) SetConfig() error {
	// DB Config
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_HOST := os.Getenv("DB_HOST")

	// DB Creds
	DB_PORT := os.Getenv("MYSQL_PORT")
	DB_USERNAME := os.Getenv("MYSQL_USERNAME")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_SCHEMA := os.Getenv("MYSQL_SCHEMA")
	DB_TIMEOUT := os.Getenv("MYSQL_TIMEOUT")

	dbPort, err := strconv.Atoi(DB_PORT)

	if err != nil {
		return fmt.Errorf("error: parsing string to int of DB_PORT: %v failed", dbPort)
	}

	dbTimeout, err := strconv.Atoi(DB_TIMEOUT)

	if err != nil {
		return fmt.Errorf("error: parsing string to int of DB_TIMEOUT: %v failed", dbTimeout)
	}

	if DB_USERNAME == "" {
		return fmt.Errorf("env: DB_USERNAME is missing")
	}

	if DB_PASSWORD == "" {
		return fmt.Errorf("env: DB_PASSWORD is missing")
	}

	if DB_SCHEMA == "" {
		return fmt.Errorf("env: DB_SCHEMA is missing")
	}

	if DB_PORT == "" {
		return fmt.Errorf("env: DB_PORT is missing")
	}

	if DB_DRIVER == "" {
		return fmt.Errorf("env: DB_PORT is missing")
	}

	if DB_HOST == "" {
		return fmt.Errorf("env: DB_PORT is missing")
	}

	return nil
}
