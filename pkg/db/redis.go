package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"strconv"
)

var (
	//REDIS_HOST     = os.Getenv("REDIS_HOST")
	//REDIS_PORT     = os.Getenv("REDIS_PORT")
	//REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	//REDIS_DB       = os.Getenv("REDIS_DB")

	REDIS_HOST     = "141.11.25.60"
	REDIS_PORT     = "6379"
	REDIS_PASSWORD = ""
	REDIS_DB       = "0"
)

func ConnectToRedis() *redis.Client {
	// address
	address := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)

	// db
	db, err := strconv.Atoi(REDIS_DB)
	if err != nil {
		logrus.Fatal("error convert string to int")
		return nil
	}

	// create new client
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: REDIS_PASSWORD,
		DB:       db,
	})

	logrus.Println("Connected to Redis")

	return client

}
