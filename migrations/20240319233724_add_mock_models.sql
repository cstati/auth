-- +goose Up
-- +goose StatementBegin

insert into problems (name, description)
VALUES ('Regression', 'Задача, в которой модель будет побирать вещественное число в качестве ответа'),
       ('Classification', 'Задача, в которой модель будет распределять объекты по заданным классам');

insert into metrics (name, description)
values ('MSE', 'Среднее разницы квадратов расхождений ожидаемого результата и эталонного'),
       ('MAE', 'Среднее разницы модулей расхождений ожидаемого результата и эталонного'),
       ('ErrorCount', 'Количество расхождений ожидаемого результата и эталонного');

insert into problem_metrics (problem_id, metric_id)
VALUES (1, 1),
       (1, 2),
       (1, 3),
       (2, 3);

insert into models (name, description, problem_id)
values ('Linear regression',
        'Модель которая подбирает коэффициенты к фичам и пытается выстроить линейную зависимость для таргета', 1),
       ('Logisitic regression',
        'Модель которая подбирает коэффициенты к фичам и пытается выстроить линейную зависимость для таргета и привести ее к диапазону [0; 1]',
        2);

insert into hyperparameters (name, description, type, default_value, model_id)
VALUES ('learning_rate', 'коэффициент скорости обучения модели', 'float', '0.1', 1),
       ('iterations_count', 'количество итераций градиентного спуска', 'int', '1000', 1),
       ('learning_rate', 'коэффициент скорости обучения модели', 'float', '0.1', 2),
       ('iterations_count', 'количество итераций градиентного спуска', 'int', '1000', 2),
       ('train_test_split', 'пропорция в которой делим на тестовую и тренировочную выборки', 'float', '0.2', 2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table metrics restart identity;
truncate table problems restart identity;
truncate table problem_metrics restart identity;
truncate table models restart identity;
truncate table hyperparameters restart identity;
-- +goose StatementEnd
