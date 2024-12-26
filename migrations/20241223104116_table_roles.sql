-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists roles (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    name VARCHAR(50)
    ) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
