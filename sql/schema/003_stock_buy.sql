-- +goose Up
CREATE TABLE stockBuy (
    id UUID PRIMARY KEY,

    userId UUID NOT NULL REFERENCES users(id),
    stockId UUID NOT NULL REFERENCES stocks(id),
    quantity INT NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT positive_quantity CHECK (quantity>0)
);

-- +goose Down

DROP TABLE stockBuy