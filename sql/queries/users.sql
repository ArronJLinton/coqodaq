-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number)
VALUES ($1, $2, $3)
RETURNING *;