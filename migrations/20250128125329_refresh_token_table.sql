-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id          SERIAL PRIMARY KEY,
    user_uuid   UUID      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at  TIMESTAMP NOT NULL,
    uuid        UUID      NOT NULL UNIQUE,
    active      BOOLEAN   NOT NULL DEFAULT TRUE,
    device_uuid UUID      NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_tokens;
-- +goose StatementEnd
