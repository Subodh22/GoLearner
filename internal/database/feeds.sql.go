// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeeds = `-- name: CreateFeeds :one
INSERT INTO feeds(id,createdAt,updatedAt,name,url,user_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING id, createdat, updatedat, name, url, user_id
`

type CreateFeedsParams struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeeds(ctx context.Context, arg CreateFeedsParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeeds,
		arg.ID,
		arg.Createdat,
		arg.Updatedat,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Updatedat,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, createdat, updatedat, name, url, user_id FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Createdat,
			&i.Updatedat,
			&i.Name,
			&i.Url,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
