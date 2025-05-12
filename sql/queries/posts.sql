-- name: CreatePost :one
INSERT INTO posts(
  id,
  created_at,
  updated_at,
  title,
  url,
  description,
  published_at,
  feed_id
)
VALUES(
$1,
$2,
$3,
$4,
$5,
$6,
$7,
$8
)
returning *;

-- name: GetPostsForUser :many
select * from posts
inner join feeds on posts.feed_id = feeds.id
where feeds.user_id = $1
order by feeds.published_at desc
limit $2;
