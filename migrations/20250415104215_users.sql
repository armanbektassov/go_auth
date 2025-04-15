-- +goose Up
CREATE TABLE users (
                      id int NOT NULL,
                      username text,
                      email text,
                      PRIMARY KEY(id)
);
-- +goose Down
DROP TABLE users;
