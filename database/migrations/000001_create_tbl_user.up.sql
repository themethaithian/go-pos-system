CREATE TABLE "tbl_user" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR NOT NULL UNIQUE,
  "password_hash" VARCHAR NOT NULL,
  "role" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
