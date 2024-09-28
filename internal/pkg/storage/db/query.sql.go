// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    google_id, email
) VALUES (
             $1, $2
         )
RETURNING id
`

type CreateUserParams struct {
	GoogleID string
	Email    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, createUser, arg.GoogleID, arg.Email)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getUserByExternalID = `-- name: GetUserByExternalID :one
SELECT id, google_id, email, created_at FROM users
WHERE google_id = $1 LIMIT 1
`

func (q *Queries) GetUserByExternalID(ctx context.Context, googleID string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByExternalID, googleID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GoogleID,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getUserRolesByEmail = `-- name: GetUserRolesByEmail :many
select role
from user_roles
where email = $1
`

func (q *Queries) GetUserRolesByEmail(ctx context.Context, email string) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserRolesByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		items = append(items, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserRolesByID = `-- name: GetUserRolesByID :many
select role
from users u
join user_roles ur on u.email = ur.email
where id = $1
`

func (q *Queries) GetUserRolesByID(ctx context.Context, id int64) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserRolesByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, err
		}
		items = append(items, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
