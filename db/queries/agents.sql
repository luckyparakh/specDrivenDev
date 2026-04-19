-- name: CreateAgent :one
INSERT INTO agents (name, email, password_hash)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetAgentByEmail :one
SELECT * FROM agents WHERE email = ? LIMIT 1;

-- name: GetAgentByID :one
SELECT * FROM agents WHERE id = ? LIMIT 1;
