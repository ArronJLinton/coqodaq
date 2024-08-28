-- name: CreateReservation :one
INSERT INTO reservations (name, party_size, time, restaurant_id, table_id, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetReservationsByUserId :many
SELECT
    res.user_id,
    res.restaurant_id,
    res.time
FROM 
    reservations res 
LEFT JOIN 
    users u ON res.user_id = u.id
WHERE 
    res.is_active = true AND res.user_id = $1;

-- name: DeleteReservation :one
DELETE FROM reservations
	WHERE id=$1
RETURNING *;


