-- name: GetSection :one
SELECT * FROM sections
WHERE id = $1 LIMIT 1;

-- name: GetSectionsByChapter :many
SELECT * FROM sections
WHERE chapter_id = $1;

-- name: CreateSection :one
INSERT INTO sections (
  chapter_id,
  title,
  number
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateSection :one
UPDATE sections
SET
  chapter_id = $2,
  title = $3,
  number = $4
WHERE id = $1
RETURNING *;

-- name: DeleteSection :exec
DELETE FROM sections
WHERE id = $1;