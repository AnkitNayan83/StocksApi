-- name: CreateStocks :one
INSERT INTO stocks (id,companyName, valuePerStock, quantity, ownerId,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetAllStocks :many
SELECT * FROM stocks
WHERE quantity > 0;