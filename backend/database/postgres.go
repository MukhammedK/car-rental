package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"car-rental-system/backend/config"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_SSLMODE"),
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {

		db, err = sql.Open("postgres", dsn)
		if err == nil {

			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}

		time.Sleep(3 * time.Second)
	}

	return nil, err
}
