package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var JWT_SECRET string

type Config struct {
	Db    *sql.DB
	Redis *redis.Client
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
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	db, err := loadDb(pgUser, pgPass, pgHost, pgPort, pgDbName)
	if err != nil {
		return nil, fmt.Errorf("error in loading db: %s", err.Error())
	}

	redis, err := loadRedis(redisHost, redisPort)
	if err != nil {
		return nil, fmt.Errorf("error in loading redis client: %s", err.Error())
	}

	return &Config{
		Db:    db,
		Redis: redis,
	}, nil
}
