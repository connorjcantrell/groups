-- name: GetGroupBook :one
SELECT * FROM group_books
WHERE id = $1 LIMIT 1;

-- name: GetGroupBooksByGroup :many
SELECT * FROM group_books
WHERE group_id = $1;


-- name: GetGroupBooksByBook :many
SELECT * FROM group_books
WHERE book_id = $1;

-- name: BookExistsInGroup :one
SELECT COUNT(*) FROM group_books WHERE book_id = $1;

-- name: CreateGroupBook :one
INSERT INTO group_books (
  group_id,
  book_id,
  completion,
  last_modified
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateGroupBook :one
UPDATE group_books
SET
  group_id = $2,
  book_id = $3,
  completion = $4,
  last_modified = $5
WHERE id = $1
RETURNING *;

-- name: DeleteGroupBook :exec
DELETE FROM group_books
WHERE id = $1;

-- it may not correct, I will check and send you later 
-- name: GetGroupWithBooks :one
-- SELECT groups.*, groups.books 
-- FROM groups
-- INNER JOIN books ON groups.book_id = books.id