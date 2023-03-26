-- name: GetCategory :one
SELECT * FROM category
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM category
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateCategory :one
INSERT INTO category (
  title, 
  description, 
  image_url
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateCategory :one
UPDATE category
SET title = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE id = $1;