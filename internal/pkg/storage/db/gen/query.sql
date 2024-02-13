-- name: GetUserByExternalID :one
SELECT * FROM users
WHERE google_id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    google_id, email
) VALUES (
             $1, $2
         )
RETURNING id;
