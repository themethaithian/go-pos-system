CREATE TABLE "tbl_product" (
    "id" varchar NOT NULL,
    "name" varchar NOT NULL,
    "description" varchar,
    "price" numeric NOT NULL DEFAULT 0.00,
    "quantity" int4 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);
