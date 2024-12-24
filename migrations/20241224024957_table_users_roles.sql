-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users_roles (
    user_id SERIAL REFERENCES users(id),
    role_id SERIAL REFERENCES roles(id),
    PRIMARY KEY (user_id, role_id)
    ) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_roles;
-- +goose StatementEnd
