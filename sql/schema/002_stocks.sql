-- +goose Up

CREATE TABLE stocks (
    id UUID PRIMARY KEY,
    companyName VARCHAR(50) NOT NULL UNIQUE,
    valuePerStock FLOAT NOT NULL,
    quantity INT NOT NULL,
    ownerId UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    
    CONSTRAINT positive_value_per_stock CHECK (valuePerStock>0),
    CONSTRAINT positive_quantity CHECK (quantity>0)
);


-- +goose Down

DROP TABLE stocks;