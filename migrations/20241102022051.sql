-- Create "accounts" table
CREATE TABLE "public"."accounts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "uuid" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX "idx_accounts_deleted_at" ON "public"."accounts" ("deleted_at");
