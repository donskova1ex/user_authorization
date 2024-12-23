-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXITS permissions (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    permissions_name VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists permissions;
-- +goose StatementEnd
