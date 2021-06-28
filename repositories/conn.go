package repository

import (
	"database/sql"
	"fmt"
	"templ/env"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = createConnection()
	if err != nil {
		log.Error(err)
	}

	go startDBStatistics()
}

func createConnection() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.DbHost, env.DbPort, env.DbUser, env.DbPassword, env.DbName)
	log.Info(fmt.Sprintf("Connecting to DB %s", psqlInfo))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	if env.DbMaxOpenConnections != 0 {
		db.SetMaxOpenConns(env.DbMaxOpenConnections)
		log.Info("Set Db max open connections to %d", env.DbMaxOpenConnections)
	}
	if env.DbMaxIdleConnections != 0 {
		db.SetMaxIdleConns(env.DbMaxIdleConnections)
		log.Info("Set Db max idle connections to %d", env.DbMaxIdleConnections)
	}
	if env.DbMaxConnectionLifetimeInMinutes != 0 {
		db.SetConnMaxLifetime(time.Duration(env.DbMaxConnectionLifetimeInMinutes) * time.Minute)
		log.Info("Set Db max connection lifetime in minutes to %d", env.DbMaxConnectionLifetimeInMinutes)
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	log.Debug("Successfully connected to the database. ")
	return db, nil
}

func startDBStatistics() {
	err := gocron.Every(10).Minutes().Do(stats)
	if err != nil {
		log.Error(err)
		return
	}
	<-gocron.Start()
}

func stats() {
	log.Info(Db.Stats())
}
