-- +goose Up
create table posts (
  id uuid not null default gen_random_uuid(),
  created_at timestamp not null,
  updated_ad timestamp not null,
  title varchar not null,
  url varchar not null,
  description varchar,
  published_at timestamp,
  feed_id uuid not null references feeds(id) on delete cascade,
  primary key(id)
);

-- +goose Down
drop table posts;
