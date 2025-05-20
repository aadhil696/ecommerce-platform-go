package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string
}

func SetupEnv() (cfg AppConfig, err error) {

	if err := godotenv.Load(); err != nil {
		log.Println("env file not found or couldn't be loaded")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable not found")
	}

	Dsn := os.Getenv("DSN")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variable not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("app Secret not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret}, nil

}
