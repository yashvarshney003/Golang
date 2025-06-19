-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
DROP COLUMN IF EXISTS product_id,
DROP COLUMN IF EXISTS variant_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
ADD COLUMN product_id TEXT,
ADD COLUMN variant_id TEXT;
-- +goose StatementEnd
