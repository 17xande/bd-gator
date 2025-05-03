-- +goose Up
create table users (
  id uuid DEFAULT gen_random_uuid(),
  created_at timestamp not null,
  updated_at timestamp not null,
  name varchar not null unique,
  primary key(id)
);

-- +goose Down
drop table users;
