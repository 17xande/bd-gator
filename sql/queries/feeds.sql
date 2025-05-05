-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
returning *;

-- name: GetFeed :one
select * from feeds
where url = $1;

-- name: ResetFeed :exec
delete from feeds;

-- name: GetFeeds :many
select feeds.id, feeds.created_at, feeds.updated_at, feeds.name, feeds.url, feeds.user_id, users.name as user_name
from feeds
inner join users on users.id = feeds.user_id;
