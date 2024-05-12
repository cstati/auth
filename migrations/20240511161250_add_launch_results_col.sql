-- +goose Up
-- +goose StatementBegin
alter table launches
    add column if not exists input  json,
    add column if not exists output json;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table launches
    drop column if exists input,
    drop column if exists output;
-- +goose StatementEnd
