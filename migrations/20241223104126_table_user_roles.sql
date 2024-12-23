-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    role_name VARCHAR(50));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
