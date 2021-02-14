package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jason-costello/schooling-covid/internal/server"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func main() {

	logger := logrus.New()

	logLevel := os.Getenv("LogLevel")

	switch strings.ToLower(logLevel){
	case "debug":
		logger.Level = logrus.DebugLevel
	case "info":
		logger.Level = logrus.InfoLevel
	case "warning":
		logger.Level = logrus.WarnLevel
	case "error":
		logger.Level = logrus.ErrorLevel
	default:
		logger.Level = logrus.ErrorLevel
	}


	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("no port provided")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("no db url provided")
	}
	dbc := server.DBConfig{
		URL:    dbURL,
	}

	app := server.NewServer(port, dbc, logger)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/districts", app.GetDistricts).Methods("GET")
	r.HandleFunc("/api/v1/district/{did}", app.GetDistrictSchools).Methods("GET")

	app.SetRouter(r)
	err = r.Walk(gorillaWalkFn)
	if err != nil {
		log.Fatal(err)
	}

	app.Serve()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	app.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
func gorillaWalkFn(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, err := route.GetPathTemplate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	return nil
}
