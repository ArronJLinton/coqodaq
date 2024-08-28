-- name: CreateRestaurantTable :one
INSERT INTO tables (restaurant_id, capacity)
VALUES ($1, $2)
RETURNING *;