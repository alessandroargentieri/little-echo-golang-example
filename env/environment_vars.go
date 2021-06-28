package env

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

var Port int

var DbHost string
var DbPort int
var DbUser string
var DbPassword string
var DbName string

var LogLevel string
var DbMaxOpenConnections int
var DbMaxIdleConnections int
var DbMaxConnectionLifetimeInMinutes int

func init() {
	LoadConfigs()
}

func LoadConfigs() {

	var err error

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Error(err)
		Port = 8081
	}

	DbHost = os.Getenv("DB_HOST")

	DbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error(err)
		DbPort = 5432
	}

	DbUser = os.Getenv("DB_USER")

	DbPassword = os.Getenv("DB_PASSWORD")

	DbName = os.Getenv("DB_NAME")

	LogLevel = os.Getenv("LOG_LEVEL")
	if LogLevel == "" {
		LogLevel = "error"
	}

	DbMaxOpenConnections, err = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err != nil {
		log.Info("No $DB_MAX_OPEN_CONN value found: set to default.")
		DbMaxOpenConnections = 0 //this won't be applied and will leave the value to its default
	}

	DbMaxIdleConnections, err = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		log.Info("No $DB_MAX_IDLE_CONN value found: set to default.")
		DbMaxIdleConnections = 0 //this won't be applied and will leave the value to its default
	}

	DbMaxConnectionLifetimeInMinutes, err = strconv.Atoi(os.Getenv("DB_MAX_CONN_LIFETIME"))
	if err != nil {
		log.Info("No $DB_MAX_CONN_LIFETIME value found: set to default.")
		DbMaxConnectionLifetimeInMinutes = 0 //this won't be applied and will leave the value to its default
	}
}
