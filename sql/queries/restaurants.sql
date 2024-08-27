-- name: CreateRestaurant :one
INSERT INTO restaurants (name, dietary_restrictions)
VALUES ($1, $2)
RETURNING *;


-- name: GetRestaurants :many
SELECT * FROM restaurants;

-- name: GetRestaurantsByDietaryRestrictionsAndTableCapacity :many
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
    t.capacity >= $1 AND r.dietary_restrictions && $2;