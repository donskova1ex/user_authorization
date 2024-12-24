-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists permissions (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(50) unique,
    permissions_name VARCHAR(50)
    ) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
