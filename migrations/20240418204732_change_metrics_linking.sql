-- +goose Up
-- +goose StatementBegin
drop table if exists problem_metrics;
drop table if exists metrics;

create table if not exists metrics
(
    id           bigint primary key generated always as identity,
    model_id     bigint   not null references models (id),
    metric_name  text      not null,
    created_at   timestamp not null default now(),
    updated_at   timestamp not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists metrics;
-- +goose StatementEnd
