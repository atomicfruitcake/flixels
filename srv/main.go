package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/atomicfruitcake/flixels/constants"
	"github.com/atomicfruitcake/flixels/handlers/createjob"
	"github.com/atomicfruitcake/flixels/handlers/getjob"
	"github.com/atomicfruitcake/flixels/handlers/health"
	"github.com/atomicfruitcake/flixels/handlers/root"
	"github.com/atomicfruitcake/flixels/middleware/auth"
	"github.com/atomicfruitcake/flixels/middleware/logging"

	"github.com/gorilla/mux"
)

func main() {
	err := redis.Ping()
	if err != nil {
		log.Fatal("Could not connect to Redis, cannot boot flixels")

	}
	log.Println("Starting a new flixels HTTP Server")
	log.Println("Building the Gorilla MUX Router")
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", root.Handler).Methods("GET")
	r.HandleFunc("/health", health.Handler).Methods("GET")
	r.Use(logging.Middleware)

	jr := r.PathPrefix("/job").Subrouter()
	jr.HandleFunc("/createJob", createjob.Handler).Methods("POST")
	jr.HandleFunc("/getJob", getjob.Handler).Methods("GET", "POST")

	am := auth.Middleware{}
	am.Populate()
	jr.Use(am.Middleware)
	http.Handle("/", r)

	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"Graceful Shutdown time is 15 seconds",
	)
	flag.Parse()
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", constants.AppPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 120,
		Handler:      r,
	}
	log.Printf("flixels HTTP Webserver is running on port %s\n", constants.AppPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting Down flixels Server")
	os.Exit(0)
}
