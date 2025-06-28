package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xufor/golang-learning/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", app.HealthCheck)
	router.Get("/workouts/{id}", app.Workouthandler.HandleGetWorkoutByID)
	router.Post("/workouts", app.Workouthandler.HandleCreateWorkout)
	return router
}
