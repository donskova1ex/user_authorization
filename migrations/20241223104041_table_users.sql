-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    name VARCHAR(64),
    password_hash VARCHAR(250)
) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
