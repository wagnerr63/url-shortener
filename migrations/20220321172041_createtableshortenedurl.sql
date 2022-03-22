-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shortened_urls (
    id VARCHAR NOT NULL,
    url VARCHAR,
    shortened_id VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_in TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shortened_urls;
-- +goose StatementEnd
