-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS peoples (
    id serial primary key,
    surname varchar(255) not null,
    name varchar(255) not null,
    patronymic varchar(255) default null,
    passport_series int not null unique,
    passport_number int not null unique,
    task_id int default null
);

CREATE TABLE IF NOT EXISTS tasks (
    id serial primary key,
    people_id int not null,
    start_task bigint default extract(epoch from now())::bigint,
    end_task bigint default null,
    labor integer default null,
    foreign key (people_id) references peoples(id) on delete cascade
);

CREATE INDEX IF NOT EXISTS peoples_id ON peoples(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS peoples;
DROP TABLE IF EXISTS tasks;
DROP INDEX IF EXISTS peoples_id
-- +goose StatementEnd
