package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ArronJLinton/coqodaq/internal/api"
	"github.com/ArronJLinton/coqodaq/internal/config"
	"github.com/ArronJLinton/coqodaq/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

/*

1) Connect Database
2) Create Tables
	a) Restaurants
		- name
		- dietary_restrictions
		Relationships - Restaurant has many Tables
	b) Tables
		- restaurant_id
		- capacity
		- is_reserved
	c) Reservations
		- table_id
		- restaurant_id
		- name
		- time
		- party_size

4) Create Queries
	a) Create Restaurant
	b) Create Table
	c) Create Reservation
	d) Get Restaurants and filter by dietary_restrictions, party_size (party_size <= capacity), and time
	e) Get Tables and filter by restaurant_id, capacity, is_reserved
	f) Get Reservations and filter by table_id, restaurant_id, time
	g) Update Table
	h) Update Reservation
	i) Delete Reservation

5) Create Endpoints
	a) Get Restaurants with available table at specified time
	b) Create Reservation
	c) Delete Reservation
*/

type Config struct {
	DB *database.Queries
}

func main() {
	// Initialize app configuration with env variables
	c := config.InitConfig()
	// Connection to DB
	conn, err := sql.Open("postgres", c.DB_URL)
	if err != nil {
		log.Fatal("Failed to connect to Database - ", err)
	}
	// Create router
	router := chi.NewRouter()
	// Setup cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", ";http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"string"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// Create API router
	apiRouter := api.New(api.Config{
		DB:             database.New(conn),
	})
	router.Mount("/api", apiRouter)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + c.PORT,
	}
	fmt.Printf("Server starting on port %v", c.PORT)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}