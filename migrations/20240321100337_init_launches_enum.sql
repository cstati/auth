-- +goose Up
-- +goose StatementBegin
create type launch_type as enum ('train', 'predict', 'generic');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type launch_type;
-- +goose StatementEnd
