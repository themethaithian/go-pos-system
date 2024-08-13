CREATE TABLE "tbl_product" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" VARCHAR,
  "price" NUMERIC(10, 2) NOT NULL,
  "category_id" INTEGER NOT NULL,
  "stock_level" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);