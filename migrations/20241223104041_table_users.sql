-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    user_name VARCHAR(64),
    password_hash VARCHAR(250),
    password_salt VARCHAR(100)
) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
