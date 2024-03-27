-- name: CreateStocks :one
INSERT INTO stocks (id,companyName, valuePerStock, quantity, ownerId,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetAllStocks :many
SELECT * FROM stocks
WHERE quantity > 0;


-- name: UpdateStock :one
UPDATE stocks
SET companyName = $1, valuePerStock = $2, quantity = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4 AND ownerId = $5
RETURNING *;