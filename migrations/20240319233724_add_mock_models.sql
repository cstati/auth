-- +goose Up
-- +goose StatementBegin

insert into problems (id, name, description)
VALUES (1, 'Регрессия', 'Задача, в которой модель будет побирать вещественное число в качестве ответа'),
       (2, 'Классификация', 'Задача, в которой модель будет распределять объекты по заданным классам');

insert into metrics (id, name, description)
values (1, 'MSE', 'Среднее разницы квадратов расхождений ожидаемого результата и эталонного'),
       (2, 'MAE', 'Среднее разницы модулей расхождений ожидаемого результата и эталонного'),
       (3, 'ErrorCount', 'Количество расхождений ожидаемого результата и эталонного');

insert into problem_metrics (problem_id, metric_id)
VALUES (1, 1),(1, 2), (1, 3), (2, 3);

insert into models (id, name, description, problem_id)
values (1, 'Linear regression',
        'Модель которая подбирает коэффициенты к фичам и пытается выстроить линейную зависимость для таргета', 1),
       (2, 'Logisitic regression',
        'Модель которая подбирает коэффициенты к фичам и пытается выстроить линейную зависимость для таргета и привести ее к диапазону [0; 1]',
        2);

insert into hyperparameters (id, name, description, type, default_value, model_id)
VALUES (1, 'learning_rate', 'коэффициент скорости обучения модели', 'float', '0.1', 1),
       (2, 'iterations_count', 'количество итераций градиентного спуска', 'int', '1000', 1),
       (3, 'learning_rate', 'коэффициент скорости обучения модели', 'float', '0.1', 2),
       (4, 'iterations_count', 'количество итераций градиентного спуска', 'int', '1000', 2),
       (5, 'train_test_split', 'пропорция в которой делим на тестовую и тренировочную выборки', 'float', '0.2', 2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table metrics restart identity;
truncate table problems restart identity;
truncate table problem_metrics restart identity;
truncate table models restart identity;
truncate table hyperparameters restart identity;
-- +goose StatementEnd
