package api

import (
	"net/http"

	"github.com/ArronJLinton/coqodaq/internal/database"
	"github.com/go-chi/chi"
)

type Config struct {
	DB *database.Queries
}

func New(config Config) http.Handler {
	router := chi.NewRouter()
	// restaurant routes
	restaurantRouter := chi.NewRouter()	
	restaurantRouter.Get("/", config.GetRestaurantsByDietaryRestrictionsAndTableCapacity)

	// reservation routes
	reservationRouter := chi.NewRouter()
	reservationRouter.Post("/create", config.CreateReservation)
	reservationRouter.Delete("/delete", config.CreateReservation)

	router.Mount("/reservation", reservationRouter)
	router.Mount("/restaurants", restaurantRouter)

	return router
}
