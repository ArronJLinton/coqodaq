package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ArronJLinton/coqodaq/internal/database"
)
type RequestBody struct {
	Capacity int32
	DietaryRestrictions []string
	UserId int32
	Time time.Time
}

func (config *Config) GetRestaurantsByDietaryRestrictionsAndTableCapacity(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := RequestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	defer r.Body.Close()
	restaurants, err := config.DB.GetRestaurantsByDietaryRestrictionsAndTableCapacity(r.Context(), 
	database.GetRestaurantsByDietaryRestrictionsAndTableCapacityParams{
		Capacity:              params.Capacity,
		DietaryRestrictions:   params.DietaryRestrictions,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting restaurants: %s", err))
		return
	}

	// Fetch existing reservations for the user
	userReservations, err := config.DB.GetReservationsByUserId(r.Context(), params.UserId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting user reservations: %s", err))
		return
	}
	// TODO: Check if there are any reservations that conflict with the requested time
		// Need a query that returns available tables at restaurants that are not reserved at the requested time
	// Check if the user has any reservations that conflict with the requested time within 2 hours(before/after)
	if len(userReservations) > 0 {
		for _, resy := range userReservations {
			if isWithinTwoHours(resy.Time, params.Time) {
				respondWithJSON(w, http.StatusOK, struct{Message string}{ Message: "You have an existing reservation within 2 hours of this time. Please select another time or cancel your exisiting reservation."})
				return
			}
		}
		respondWithJSON(w, http.StatusOK, restaurants)
	} else {
		respondWithJSON(w, http.StatusOK, restaurants)
	}
}

func isWithinTwoHours(referenceTime time.Time, checkTime time.Time) bool {
	// Calculate the difference between the two times
	diff := checkTime.Sub(referenceTime)
	// Check if the difference is within the range of -2 hours to +2 hours
	return diff >= -2*time.Hour && diff <= 2*time.Hour
}