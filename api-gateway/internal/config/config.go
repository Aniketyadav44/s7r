package config

import (
	"database/sql"
	"fmt"
	"os"

	gpb "github.com/Aniketyadav44/s7r/api-gateway/proto/getter"
	spb "github.com/Aniketyadav44/s7r/api-gateway/proto/shortener"
	"github.com/redis/go-redis/v9"
)

var JWT_SECRET string

type Config struct {
	Port            string
	Db              *sql.DB
	Redis           *redis.Client
	ShortenerClient *spb.ShortenerClient
	GetterClient    *gpb.GetterClient
}

func LoadConfig() (*Config, error) {
	// if err := godotenv.Load("../.env"); err != nil {
	// 	return nil, fmt.Errorf("error in loading env: %s", err.Error())
	// }

	port := os.Getenv("API_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgDbName := os.Getenv("PG_DBNAME")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	shortenerGrpcUrl := os.Getenv("SHORTENER_URL")
	getterGrpcUrl := os.Getenv("GETTER_URL")

	db, err := loadDb(pgUser, pgPass, pgHost, pgPort, pgDbName)
	if err != nil {
		return nil, fmt.Errorf("error in loading db: %s", err.Error())
	}

	redis, err := loadRedis(redisHost, redisPort)
	if err != nil {
		return nil, fmt.Errorf("error in loading redis client: %s", err.Error())
	}

	shortenerClient, err := loadShortenerGrpcClient(shortenerGrpcUrl)
	if err != nil {
		return nil, fmt.Errorf("error in loading shortener grpc client: %s", err.Error())
	}

	getterClient, err := loadGetterGrpcClient(getterGrpcUrl)
	if err != nil {
		return nil, fmt.Errorf("error in loading getter grpc client: %s", err.Error())
	}

	return &Config{
		Port:            port,
		Db:              db,
		Redis:           redis,
		ShortenerClient: shortenerClient,
		GetterClient:    getterClient,
	}, nil
}
