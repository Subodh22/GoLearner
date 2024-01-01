-- name: CreateFeeds :one
INSERT INTO feeds(id,createdAt,updatedAt,name,url,user_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING * ;

-- name: GetFeeds :many 
SELECT * FROM feeds;
 