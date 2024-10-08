// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reservations.sql

package database

import (
	"context"
	"time"
)

const createReservation = `-- name: CreateReservation :one
INSERT INTO reservations (name, party_size, time, restaurant_id, table_id, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, party_size, time, restaurant_id, table_id, user_id, is_active, created_at, updated_at
`

type CreateReservationParams struct {
	Name         string
	PartySize    int32
	Time         time.Time
	RestaurantID int32
	TableID      int32
	UserID       int32
}

func (q *Queries) CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, createReservation,
		arg.Name,
		arg.PartySize,
		arg.Time,
		arg.RestaurantID,
		arg.TableID,
		arg.UserID,
	)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PartySize,
		&i.Time,
		&i.RestaurantID,
		&i.TableID,
		&i.UserID,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReservation = `-- name: DeleteReservation :one
DELETE FROM reservations
	WHERE id=$1
RETURNING id, name, party_size, time, restaurant_id, table_id, user_id, is_active, created_at, updated_at
`

func (q *Queries) DeleteReservation(ctx context.Context, id int32) (Reservation, error) {
	row := q.db.QueryRowContext(ctx, deleteReservation, id)
	var i Reservation
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PartySize,
		&i.Time,
		&i.RestaurantID,
		&i.TableID,
		&i.UserID,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getReservationsByUserId = `-- name: GetReservationsByUserId :many
SELECT
    res.user_id,
    res.restaurant_id,
    res.time
FROM 
    reservations res 
LEFT JOIN 
    users u ON res.user_id = u.id
WHERE 
    res.is_active = true AND res.user_id = $1
`

type GetReservationsByUserIdRow struct {
	UserID       int32
	RestaurantID int32
	Time         time.Time
}

func (q *Queries) GetReservationsByUserId(ctx context.Context, userID int32) ([]GetReservationsByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getReservationsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetReservationsByUserIdRow
	for rows.Next() {
		var i GetReservationsByUserIdRow
		if err := rows.Scan(&i.UserID, &i.RestaurantID, &i.Time); err != nil {
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
