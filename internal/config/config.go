package config

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	SelfAddress    string
	Database       string
	AccrualAddress string
}

func (c Config) String() string {
	return fmt.Sprintf("server: %s, db: %s, accrual: %s", c.SelfAddress, c.Database, c.AccrualAddress)
}

var AppConfig Config

func InitConfig() {

	conf := Config{}
	envAddress := os.Getenv("RUN_ADDRESS")
	if envAddress == "" {
		envAddress = "localhost:8080"
	}
	envDSN := os.Getenv("DATABASE_URI")
	if envDSN == "" {
		envDSN = "postgres://gophermart:gophermart@localhost:5432/gophermart?sslmode=disable"
	}

	flag.StringVar(&conf.SelfAddress, "a", envAddress, "Эндпоинт сервера HOST:PORT")
	flag.StringVar(&conf.Database, "d", envDSN, "Адрес подключения к базе данных.")

	log.Println("Config:", conf)
	AppConfig = conf
}
