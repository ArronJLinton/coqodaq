package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ArronJLinton/coqodaq/internal/database"
)

func (config *Config) GetRestaurantsByDietaryRestrictionsAndTableCapacity(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := database.GetRestaurantsByDietaryRestrictionsAndTableCapacityParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	defer r.Body.Close()
	restaurants, err := config.DB.GetRestaurantsByDietaryRestrictionsAndTableCapacity(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting restaurants: %s", err))
		return
	}

	// check to see if the user has any current reservations
	reservations, err := config.DB.GetReservationsByUserId(r.Context(), params.UserId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user reservations: %s", err))
		return
	}

	// if reservations exist, check to see if the user has a reservation at any of the restaurants within 2hours before/after the requested time
		// if there is a conflict, respond to the user with a message - "You already have a reservation at this time"
		// else, return the list of restaurants
	// else, return the list of restaurants

	if len(reservations) > 0 {
		// every reservation is a reservation for a restaurant
		for _, reservation := range reservations {
			// for each restaurant in the list of restaurants
			for _, restaurant := range restaurants {
				// if the user has a reservation at the restaurant
				if reservation.RestaurantID == restaurant.ID {
					// check to see if the reservation time is within 2 hours of the requested time
					if reservation.Time.Add(-2 * time.Hour).Before(params.Time) && reservation.Time.Add(2 * time.Hour).After(params.Time) {
						respondWithError(w, http.StatusBadRequest, "You already have a reservation at this time")
						return
					}
				}
			}
		}

		respondWithJSON(w, http.StatusOK, restaurants)
	} else {
		respondWithJSON(w, http.StatusOK, restaurants)
	}
}