CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "city" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "characteristic" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "image" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "alt_text" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "policy" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "category_id" bigint,
  "city_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "image" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

CREATE TABLE "product_characteristic" (
  "product_id" bigserial,
  "characteristic_id" bigserial,
  PRIMARY KEY ("product_id", "characteristic_id")
);

ALTER TABLE "product_characteristic" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_characteristic" ADD FOREIGN KEY ("characteristic_id") REFERENCES "characteristic" ("id");


CREATE TABLE "product_policy" (
  "product_id" bigserial,
  "policy_id" bigserial,
  PRIMARY KEY ("product_id", "policy_id")
);

ALTER TABLE "product_policy" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_policy" ADD FOREIGN KEY ("policy_id") REFERENCES "policy" ("id");
