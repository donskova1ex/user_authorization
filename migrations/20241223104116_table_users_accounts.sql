-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_accounts (
    user_id SERIAL REFERENCES users(id),
    person_id SERIAL REFERENCES persons(id),
    role_id SERIAL REFERENCES roles(id),
    PRIMARY KEY (user_id, person_id, role_id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_accounts;
-- +goose StatementEnd
