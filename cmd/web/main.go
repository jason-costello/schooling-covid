package main

import (
	"github.com/jason-costello/schooling-covid/internal/server"
	"github.com/sirupsen/logrus"
)

func main(){

	logger := logrus.New()
	dbConfig := server.DBConfig{
		Host:               "",
		Port:               0,
		User:               "",
		Pass:               "",
		DBName:             "",
		SSLMode:            "",
		SSLCert:            "",
		SSLKey:             "",
		SSLServerCA:        "",
		LoggingEnabled:     false,
		Dialect:            "",
		MaxOpenConnections: 0,
		MaxIdleConnections: 0,
		ConnMaxLifetime:    0,
		Logger:             nil,
	}
	app := server.NewApplication(dbConfig, logger)

	app.Serve()


}
