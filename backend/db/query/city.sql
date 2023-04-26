-- name: GetCity :one
SELECT * FROM city
WHERE id = $1 LIMIT 1;

-- name: ListCities :many
SELECT * FROM city
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateCity :one
INSERT INTO city (name) VALUES ($1) RETURNING *;