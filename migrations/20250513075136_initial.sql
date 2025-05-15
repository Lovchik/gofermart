-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id    int NOT NULL,
    login text,
    password  text,
    PRIMARY KEY (id)
);
-- +goose StatementEnd
