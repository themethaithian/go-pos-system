CREATE TABLE "tbl_inventory_log" (
  "id" SERIAL PRIMARY KEY,
  "product_id" INTEGER NOT NULL,
  "change_quantity" INTEGER NOT NULL,
  "reason" VARCHAR,

);
