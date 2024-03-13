-- +goose Up
-- +goose StatementBegin
create table if not exists users (
  id bigint primary key generated always as identity,
  google_id text not null unique,
  email text not null unique,
  created_at timestamp with time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
