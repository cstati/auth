-- +goose Up
-- +goose StatementBegin
CREATE TYPE model_training_status AS ENUM ('not_started', 'in_progress', 'error', 'done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists model_training_status;
-- +goose StatementEnd
