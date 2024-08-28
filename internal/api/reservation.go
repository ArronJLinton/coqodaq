package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArronJLinton/coqodaq/internal/database"
)

func (config *Config) CreateReservation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := database.CreateReservationParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	defer r.Body.Close()
	reservation, err := config.DB.CreateReservation(r.Context(), params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating reservation: %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, reservation)
}

func (config *Config) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := struct{ ID int32 }{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	defer r.Body.Close()
	reservation, err := config.DB.DeleteReservation(r.Context(), params.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting reservation: %s", err))
		return
	}
	respondWithJSON(w, http.StatusOK, reservation)
}