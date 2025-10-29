-- Active: 1760184803924@@127.0.0.1@5432
-- name: CreateUser :one
INSERT INTO
    users (
        email,
        phone_number,
        password_hash
    )
VALUES ($1, $2, $3)
RETURNING
    id,
    email,
    phone_number,
    role,
    deleted,
    created_at;

-- name: GetUserByEmail :one
SELECT
    id,
    email,
    phone_number,
    role,
    deleted,
    created_at,
    updated_at
FROM users
WHERE
    email = $1
    AND deleted = FALSE;

-- name: GetAllUsers :many
SELECT
    id,
    email,
    phone_number,
    role,
    deleted,
    created_at,
    updated_at
FROM users
WHERE
    deleted = FALSE
ORDER BY created_at DESC
LIMIT $1
OFFSET
    $2;

-- name: DeleteUser :exec
-- UPDATE users

-- -- name UpdateUser :one
-- -- UPDATE user WHERE email = $1 SET