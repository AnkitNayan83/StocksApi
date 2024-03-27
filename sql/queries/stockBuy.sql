-- name: CreateStockBuy :one
INSERT INTO stockBuy(id,userId,stockId,quantity,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: UpdateStockQuantity :one
UPDATE stockBuy
SET quantity = quantity - $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;