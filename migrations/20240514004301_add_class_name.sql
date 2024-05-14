-- +goose Up
-- +goose StatementBegin
ALTER TABLE models
    ADD COLUMN IF NOT EXISTS class_name TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE models
    DROP COLUMN IF EXISTS class_name;
-- +goose StatementEnd
