package config

import (
	"database/sql"
	"fmt"
	"os"
)

var JWT_SECRET string

type Config struct {
	Db *sql.DB
}

func LoadConfig() (*Config, error) {
	// if err := godotenv.Load("../.env"); err != nil {
	// 	return nil, fmt.Errorf("error in loading env: %s", err.Error())
	// }

	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgDbName := os.Getenv("PG_DBNAME")

	db, err := loadDb(pgUser, pgPass, pgHost, pgPort, pgDbName)
	if err != nil {
		return nil, fmt.Errorf("error in loading db: %s", err.Error())
	}

	return &Config{
		Db: db,
	}, nil
}
