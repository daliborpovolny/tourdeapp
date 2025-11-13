--* USER

-- name: GetUser :one
SELECT * FROM user WHERE user.id = ?;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE user.email = ?;

-- name: GetUserBySessionToken :one
SELECT * FROM user
JOIN session on user.id = session.user_id
WHERE session.token = ?;

-- name: CreateUser :one
INSERT INTO user (first_name, last_name, hash, email) VALUES (?, ?, ?, ?) RETURNING *;

--* Session

-- name: CreateSession :one
INSERT INTO session (
    user_id, token, created_at, expires_at
) VALUES (
    ?, ?, ?, ?
) RETURNING *;