-- +goose Up
-- +goose StatementBegin
create table if not exists launches
(
    id            bigint primary key generated by default as identity,
    launch_type   text                     not null,
    user_id       bigint                   not null references users (id),
    name          text                     not null,
    description   text                     not null,
    created_at    timestamp with time zone not null default now(),
    updated_at    timestamp with time zone not null default now(),
    finished_at   timestamp with time zone,
    launch_status text                     not null,
    launch_error  text
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists launches;

-- +goose StatementEnd
