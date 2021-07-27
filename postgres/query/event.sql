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
  video_link,
  start_time,
  duration,
  description
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: UpdateEvent :one
UPDATE events
SET 
  group_id = $2,
  book_id = $3,
  chapter_id = $4,
  video_link = $5,
  start_time = $6,
  duration = $7,
  description = $8
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;