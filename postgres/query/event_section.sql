-- name: GetEventSection :one
SELECT * FROM event_sections
WHERE id = $1 LIMIT 1;

-- name: GetEventSectionsByEvent :many
SELECT * FROM event_sections
WHERE event_id = $1;

-- name: CreateEventSection :one
INSERT INTO event_sections (
  event_id,
  section_id,
  presenter,
  complete
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateEventSection :one
UPDATE event_sections
SET 
  event_id = $2,
  section_id = $3,
  presenter = $4,
  complete = $5
WHERE id = $1
RETURNING *;

-- name: DeleteEventSection :exec
DELETE FROM event_sections
WHERE id = $1;