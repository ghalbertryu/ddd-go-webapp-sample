-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM user
ORDER BY account;

-- name: CreateUser :execresult
INSERT INTO user (
  account, password
) VALUES (
  ?, ?
);
