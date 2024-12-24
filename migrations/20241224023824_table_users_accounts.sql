-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users_accounts (
    user_id SERIAL REFERENCES users(id),
    person_id SERIAL REFERENCES persons(id),
    PRIMARY KEY (user_id, person_id)
    ) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users_accounts;
-- +goose StatementEnd
