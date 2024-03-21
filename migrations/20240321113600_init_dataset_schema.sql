-- +goose Up
-- +goose StatementBegin
create table if not exists dataset_schemas (
    dataset_id bigint references datasets(id),
    column_number int not null,
    column_name text not null,
    column_type text not null,
    primary key (dataset_id, column_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists dataset_schemas;
-- +goose StatementEnd
