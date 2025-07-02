package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func loadDb(pgUser, pgPass, pgHost, pgPort, pgDbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, pgHost, pgPort, pgDbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error in connecting to db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error in pinging db: %s", err.Error())
	}

	return db, nil
}
