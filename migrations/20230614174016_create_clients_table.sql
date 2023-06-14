-- +goose Up
CREATE TABLE clients (
    id UUID PRIMARY KEY,
    balance NUMERIC(12, 2) NOT NULL DEFAULT 0.0,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE clients;
