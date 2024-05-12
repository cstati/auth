-- +goose Up
-- +goose StatementBegin

insert into problems (name, description)
VALUES ('regression', 'Задача, в которой модель будет побирать вещественное число в качестве ответа'),
       ('classification', 'Задача, в которой модель будет распределять объекты по заданным классам'),
       ('clusterization', 'Задача, в которой модель будет объединять объекты в группы по их сходству');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table problems restart identity;
-- +goose StatementEnd
