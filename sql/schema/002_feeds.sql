-- +goose Up
create table feeds (
  id uuid DEFAULT gen_random_uuid(),
  created_at timestamp not null,
  updated_at timestamp not null,
  name varchar not null,
  url varchar not null unique,
  user_id uuid not null references users(id) on delete cascade,
  primary key(id)
);

-- +goose Down
drop table feeds;
