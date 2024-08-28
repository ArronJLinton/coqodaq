package main

import (
	"database/sql"
	"log"

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
}

func BulkInsert(unsavedRows []*struct{}, db *database.Queries) error {
    valueStrings := make([]string, 0, len(unsavedRows))
    valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
    for _, post := range unsavedRows {
        valueStrings = append(valueStrings, "(?, ?, ?)")
        valueArgs = append(valueArgs, post.Column1)
        valueArgs = append(valueArgs, post.Column2)
        valueArgs = append(valueArgs, post.Column3)
    }
    // stmt := fmt.Sprintf("INSERT INTO %s (column1, column2, column3) VALUES %s", strings.Join(valueStrings, ","))
    // _, err := db.Exec(stmt, valueArgs...)
    return err
}