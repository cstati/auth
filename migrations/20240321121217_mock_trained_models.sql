-- +goose Up
-- +goose StatementBegin
insert into launches (id, launch_type, name, description, launch_error)
values (1, 'train', 'Linear regression train', 'Some description', null),
       (2, 'train', 'Logistic regression train', 'Some description 2', 'launch_error'),
       (3, 'predict', 'Linear regression predict', 'Some description 3', null);

insert into trained_models (id, launch_id, name, description, model_id, model_training_status, training_dataset_id,
                            target_column)
values (1, 1, 'Linear regression train', 'desc1', 1, 'in_progress', 2, 'target'),
       (2, 2, 'Logistic regression train', 'desc12', 2, 'error', 2, 'abc');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table trained_models restart identity;
truncate table launches restart identity;
-- +goose StatementEnd
