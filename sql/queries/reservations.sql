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