-- name: GetEventAttendee :one
SELECT * FROM event_attendees
WHERE id = $1 LIMIT 1;

-- name: GetEventAttendeesByEvent :many
SELECT * FROM event_attendees
WHERE event_id = $1;

-- name: GetEventAttendeesByUser :many
SELECT * FROM event_attendees
WHERE user_id = $1;

-- name: CreateEventAttendee :one
INSERT INTO event_attendees (
  event_id,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateEventAttendee :one
UPDATE event_attendees
SET
  event_id = $2,
  user_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEventAttendee :exec
DELETE FROM event_attendees
WHERE id = $1;