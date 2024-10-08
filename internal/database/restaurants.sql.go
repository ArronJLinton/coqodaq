// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: restaurants.sql

package database

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createRestaurant = `-- name: CreateRestaurant :one
INSERT INTO restaurants (name, dietary_restrictions)
VALUES ($1, $2)
RETURNING id, name, dietary_restrictions, created_at, updated_at
`

type CreateRestaurantParams struct {
	Name                string
	DietaryRestrictions []string
}

func (q *Queries) CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (Restaurant, error) {
	row := q.db.QueryRowContext(ctx, createRestaurant, arg.Name, pq.Array(arg.DietaryRestrictions))
	var i Restaurant
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.DietaryRestrictions),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRestaurants = `-- name: GetRestaurants :many
SELECT id, name, dietary_restrictions, created_at, updated_at FROM restaurants
`

func (q *Queries) GetRestaurants(ctx context.Context) ([]Restaurant, error) {
	rows, err := q.db.QueryContext(ctx, getRestaurants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Restaurant
	for rows.Next() {
		var i Restaurant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			pq.Array(&i.DietaryRestrictions),
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRestaurantsByDietaryRestrictionsAndTableCapacity = `-- name: GetRestaurantsByDietaryRestrictionsAndTableCapacity :many
SELECT
    t.id,
    t.capacity,
    t.restaurant_id,
    r.name,
    r.dietary_restrictions
FROM 
    tables t 
LEFT JOIN 
    restaurants r ON r.id = t.restaurant_id
WHERE 
    t.capacity >= $1 AND r.dietary_restrictions && $2
`

type GetRestaurantsByDietaryRestrictionsAndTableCapacityParams struct {
	Capacity            int32
	DietaryRestrictions []string
}

type GetRestaurantsByDietaryRestrictionsAndTableCapacityRow struct {
	ID                  int32
	Capacity            int32
	RestaurantID        int32
	Name                sql.NullString
	DietaryRestrictions []string
}

func (q *Queries) GetRestaurantsByDietaryRestrictionsAndTableCapacity(ctx context.Context, arg GetRestaurantsByDietaryRestrictionsAndTableCapacityParams) ([]GetRestaurantsByDietaryRestrictionsAndTableCapacityRow, error) {
	rows, err := q.db.QueryContext(ctx, getRestaurantsByDietaryRestrictionsAndTableCapacity, arg.Capacity, pq.Array(arg.DietaryRestrictions))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRestaurantsByDietaryRestrictionsAndTableCapacityRow
	for rows.Next() {
		var i GetRestaurantsByDietaryRestrictionsAndTableCapacityRow
		if err := rows.Scan(
			&i.ID,
			&i.Capacity,
			&i.RestaurantID,
			&i.Name,
			pq.Array(&i.DietaryRestrictions),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
