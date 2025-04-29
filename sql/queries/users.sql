-- name: ListUsersAsc :many
SELECT * FROM users
ORDER BY name ASC
LIMIT ? OFFSET ?;

-- name: ListUsersDesc :many
SELECT * FROM users
ORDER BY name DESC
LIMIT ? OFFSET ?;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;
 
-- name: CreateUser :exec
INSERT INTO users (id,name,email,password)
VALUES (?,?,?,?);

-- name: UpdateUser :exec
UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
