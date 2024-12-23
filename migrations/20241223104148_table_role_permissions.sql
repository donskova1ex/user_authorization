-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id SERIAL REFERENCES user_roles(id),
    permissions_id SERIAL REFERENCES permissions(id),
    PRIMARY KEY (role_id, permissions_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists role_permissions;
-- +goose StatementEnd
