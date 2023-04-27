-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM product
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateProduct :one
INSERT INTO product (
  title, 
  description, 
  category_id,
  city_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateProduct :one
UPDATE product
SET title = $2, description = $3, category_id=$4, city_id=$5
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;