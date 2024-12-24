-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists role_permissions (
    role_id SERIAL REFERENCES roles(id),
    permissions_id SERIAL REFERENCES permissions(id),
    PRIMARY KEY (role_id, permissions_id)
    ) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
-- +goose StatementEnd
