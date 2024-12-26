-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    firstname VARCHAR(50),
    middlename VARCHAR(50),
    lastname VARCHAR(50),
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
