-- +goose Up
-- +goose StatementBegin

CREATE TYPE dataset_status AS ENUM ('', 'initializing', 'loading', 'waits_convertation', 'loading_error','convertation_in_progress','convertation_error','ready');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists dataset_status;
-- +goose StatementEnd
