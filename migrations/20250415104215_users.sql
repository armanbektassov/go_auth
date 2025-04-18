-- +goose Up
CREATE TABLE users
(
    id         serial primary key,
    name       varchar,
    email      varchar,
    role       int,
    password   text,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose Down
DROP TABLE users;