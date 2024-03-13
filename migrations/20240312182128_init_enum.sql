-- +goose Up
-- +goose StatementBegin

CREATE TYPE dataset_status AS ENUM ('initializing', 'loading', 'ready', 'error');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists dataset_status;
-- +goose StatementEnd
