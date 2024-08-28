package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ArronJLinton/coqodaq/internal/config"
	"github.com/ArronJLinton/coqodaq/internal/database"
)

func main() {
	// Initialize app configuration with env variables
	c := config.InitConfig()
	// Connection to DB
	conn, err := sql.Open("postgres", c.DB_URL)
	if err != nil {
		log.Fatal("Failed to connect to Database - ", err)
	}

	Queries := database.New(conn)
	bulkInsert(Queries)
}	

func bulkInsert(db *database.Queries) {
	// Create Restaurants
	restaurants := []database.CreateRestaurantParams{
		{
			Name:               "Restaurant A",
			DietaryRestrictions: []string{"Vegan", "Vegetarian"},
		},
		{
			Name:               "Restaurant B",
			DietaryRestrictions: []string{"Gluten-Free", "Dairy-Free"},
		},
		{
			Name:               "Restaurant C",
			DietaryRestrictions: []string{"Keto", "Paleo"},
		},
	}
	for _, restaurant := range restaurants {
		_, err := db.CreateRestaurant(context.Background(), restaurant)
		if err != nil {
			log.Fatal("Failed to create restaurant - ", err)
		}
	}

	// Create Tables
	tables := []database.CreateRestaurantTableParams{
		{
			RestaurantID: 1,
			Capacity:     4,
		},
		{
			RestaurantID: 1,
			Capacity:     6,
		},
		{
			RestaurantID: 2,
			Capacity:     2,
		},
		{
			RestaurantID: 2,
			Capacity:     4,
		},
		{
			RestaurantID: 3,
			Capacity: 2,
		},
		{
			RestaurantID: 3,
			Capacity: 4,
		},
	}
	for _, table := range tables {
		_, err := db.CreateRestaurantTable(context.Background(), table)
		if err != nil {
			log.Fatal("Failed to create table - ", err)
		}
	}

	// Create Users
	users := []database.CreateUserParams{	
		{
			FirstName: "Fred",
			LastName:  "Flintstone",
			PhoneNumber: "123-456-7890",
		},
		{
			FirstName: "George",
			LastName:  "Jetson",
			PhoneNumber: "098-765-4321",
		},
	}
	for _, user := range users {
		_, err := db.CreateUser(context.Background(), user)
		if err != nil {
			log.Fatal("Failed to create user - ", err)
		}
	}

	// Create Reservations
	reservations := []database.CreateReservationParams{
		{
			TableID:    1,
			RestaurantID: 1,
			UserID:    2,
			Name:       "Jetson Family",
			Time:       time.Now().AddDate(0, 0, 1),
			PartySize:  4,
		},
		{
			TableID:    2,
			RestaurantID: 1,
			UserID:    1,
			Name:       "Flintstone Family",
			Time:       time.Now().AddDate(0, 0, 2),
			PartySize:  4,
		},
	}

	for _, reservation := range reservations {
		_, err := db.CreateReservation(context.Background(), reservation)
		if err != nil {
			log.Fatal("Failed to create reservation - ", err)
		}
	}
}