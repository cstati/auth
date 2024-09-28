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


-- name: GetUserRolesByEmail :many
select role
from user_roles
where email = $1;

-- name: GetUserRolesByID :many
select role
from users u
join user_roles ur on u.email = ur.email
where id = $1;
