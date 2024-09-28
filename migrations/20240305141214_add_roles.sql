-- +goose Up
-- +goose StatementBegin
create table if not exists user_roles
(
    email      text                     not null,
    role       text                     not null,
    created_at timestamp with time zone not null default now(),
    primary key (email, role)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists user_roles;
-- +goose StatementEnd
