-- name: CreateLog :one
INSERT INTO logs (
    model,
    user_uuid,
    user_name,
    action,
    action_at,
    description
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetLogByID :one
SELECT * FROM logs
WHERE id = $1;

-- name: ListLogs :many
SELECT * FROM logs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListLogsByUserUUID :many
SELECT * FROM logs
WHERE user_uuid = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListLogsByModel :many
SELECT * FROM logs
WHERE model = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3; 