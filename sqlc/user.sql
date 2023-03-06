-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: GetUserByAccount :one
SELECT * FROM user
WHERE account = ?;

-- name: CreateUser :execresult
INSERT INTO user (
  account, password
) VALUES (
  ?, ?
);

-- name: UpdateUser :exec
UPDATE user SET account = ?, password = ?
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM user
ORDER BY account;