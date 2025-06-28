package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xufor/golang-learning/internal/api"
)

type Application struct {
	Logger         *log.Logger
	Workouthandler *api.Workouthandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		Workouthandler: workoutHandler,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available.")
}
