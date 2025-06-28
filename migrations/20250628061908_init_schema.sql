-- +goose Up
-- +goose StatementBegin
CREATE TYPE token_status AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS auth (
    id SERIAL PRIMARY KEY,
    access_token VARCHAR NOT NULL,
    refresh_token VARCHAR NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    status token_status,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth;
-- +goose StatementEnd
