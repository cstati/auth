-- +goose Up
-- +goose StatementBegin

create table if not exists datasets
(
    id           bigint primary key generated always as identity,
    name         text                     not null,
    version      text                     not null default '0.0',
    status       dataset_status           not null default 'initializing',
    rows_count   bigint                   not null default 0,
    creator_id   bigint                   not null references users (id),
    created_at   timestamp with time zone not null default now(),
    updated_at   timestamp with time zone not null default now(),
    upload_error text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists datasets;
-- +goose StatementEnd
