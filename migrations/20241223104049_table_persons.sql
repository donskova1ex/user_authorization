-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    first_name VARCHAR(50),
    second_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(50),
    address VARCHAR(150),
    age INTEGER,
    gender VARCHAR(15),
    date_of_birth DATE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd
