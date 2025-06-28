package main

import (
	"net/http"
	"time"

	"github.com/xufor/golang-learning/internal/app"
	"github.com/xufor/golang-learning/internal/routes"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our app!")

	router := routes.SetupRoutes(app)
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
