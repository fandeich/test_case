-- +goose Up
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    status     INT NOT NULL,
    note       TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE INDEX idx_transactions_status ON transactions(status);

-- +goose Down
DROP INDEX idx_transactions_status;
DROP TABLE transactions;
