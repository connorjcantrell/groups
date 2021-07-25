-- name: GetChapter :one
SELECT * FROM chapters
WHERE id = $1 LIMIT 1;

-- name: GetChapters :many
SELECT * FROM chapters;

-- name: GetChaptersByBook :many
SELECT * FROM chapters
WHERE book_id = $1;

-- name: CreateChapter :one
INSERT INTO chapters (
    book_id,
    title,
    number
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateChapter :one
UPDATE chapters
SET
  book_id = $2,
  title = $3,
  number = $4
WHERE id = $1
RETURNING *;

-- name: DeleteChapter :exec
DELETE FROM chapters
WHERE id = $1;