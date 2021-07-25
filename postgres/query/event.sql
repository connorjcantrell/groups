-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: GetEvents :many
SELECT * FROM events;

-- name: CreateEvent :one
INSERT INTO events (
  group_id,
  book_id,
  chapter_id,
  start_time,
  duration,
  description
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateEvent :one
UPDATE events
SET 
  group_id = $2,
  book_id = $3,
  chapter_id = $4,
  start_time = $5,
  duration = $6,
  description = $7
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;