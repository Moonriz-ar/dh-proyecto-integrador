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
  "category_id" bigint NOT NULL,
  "city_id" bigint NOT NULL,
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

-- insert cities of argentina
INSERT INTO "city" (name) VALUES
  ('Buenos Aires'),
  ('Córdoba'),
  ('Rosario'),
  ('La Plata'),
  ('Mar del Plata'),
  ('San Miguel de Tucumán'),
  ('Ciudad de Salta'),
  ('Santa Fe de la Vera Cruz'),
  ('San Juan'),
  ('Resistencia'),
  ('Ciudad de Corrientes'),
  ('Posadas'),
  ('Santiago del Estero'),
  ('San Salvador de Jujuy'),
  ('Bahía Blanca'),
  ('Paraná'),
  ('Neuquén'),
  ('Formosa'),
  ('San Luis'),
  ('La Rioja'),
  ('Catamarca'),
  ('Río Gallegos'),
  ('Ushuaia'),
  ('Viedma'),
  ('Rawson'),
  ('Santa Rosa'),
  ('San Carlos de Bariloche'),
  ('Comodoro Rivadavia'),
  ('Junín'),
  ('Venado Tuerto');

-- insert car categories in spanish
INSERT INTO category (title, description, image_url) VALUES
  ('Compacto', 'Automóviles pequeños y eficientes en combustible.', 'https://asset.cloudinary.com/andrealinar/3275e07b31b7bbf4f46cb1380ed91a38'),
  ('Sedán', 'Automóviles de tamaño mediano con 4 puertas.', 'https://asset.cloudinary.com/andrealinar/57aacd17fb9c5353677b206f0fb5d173'),
  ('SUV', 'Vehículos utilitarios deportivos con tracción en las cuatro ruedas.', 'https://asset.cloudinary.com/andrealinar/5e0df7c75ec1883533753dd04f81baf7'),
  ('Deportivo', 'Automóviles de alto rendimiento y velocidad.', 'https://asset.cloudinary.com/andrealinar/57b60d4fd925cec7a7c985cd0d845615'),
  ('Camioneta', 'Vehículos utilitarios con espacio de carga en la parte trasera.', 'https://asset.cloudinary.com/andrealinar/3819c197fda7272a5cf65968404f04b3'),
  ('Furgoneta', 'Vehículos utilitarios con espacio interior amplio para pasajeros o carga.', 'https://asset.cloudinary.com/andrealinar/06d929eb17bef042d1a6eeef16763395');