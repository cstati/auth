-- +goose Up
-- +goose StatementBegin
create table if not exists train_hyperparameters
(
    train_model_id    bigint not null references trained_models (id),
    hyperparameter_id bigint not null references hyperparameters (id),
    value             jsonb  not null,
    primary key (train_model_id, hyperparameter_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists train_hyperparameters;
-- +goose StatementEnd
