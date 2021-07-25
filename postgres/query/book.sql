-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBooksByCategory :many
SELECT * FROM books
WHERE category = $1;

-- name: CreateBook :one
INSERT INTO books (
  title,
  author,
  category
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET 
  title = $2,
  author = $3,
  category = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;