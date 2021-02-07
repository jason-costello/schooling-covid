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
	"time"
)

func main(){

	logger := logrus.New()

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err!=nil{
		panic("no port provided")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == ""{
		panic("no db url provided")
	}
	dbc := server.DBConfig{
		URL:  dbUrl  ,
		Logger: nil,
	}

	app := server.NewServer(port, dbc, logger)


	r := mux.NewRouter()

	r.HandleFunc("/api/v1/districts", app.GetDistricts).Methods("GET")
	r.HandleFunc("/api/v1/district/{id}", app.GetDistrict).Methods("GET")
	r.HandleFunc("/api/v1/district/{id}/schools", app.GetDistrictSchools).Methods("GET")
	r.HandleFunc("/api/v1/district/{id}/school/{id}", app.GetDistrictSchool).Methods("GET")
	r.HandleFunc("/", app.GetDistricts).Methods("GET")


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
