CREATE TABLE `category` (
  `id` int PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `city` (
  `id` int PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `characteristic` (
  `id` int PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `image` (
  `id` int PRIMARY KEY,
  `product_id` int,
  `alt_text` varchar(255) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `policy` (
  `id` int PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `product` (
  `id` int PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  `city_id` int NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

ALTER TABLE `product` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

ALTER TABLE `product` ADD FOREIGN KEY (`city_id`) REFERENCES `city` (`id`);

ALTER TABLE `image` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

CREATE TABLE `product_characteristic` (
  `product_id` int,
  `characteristic_id` int,
  PRIMARY KEY (`product_id`, `characteristic_id`)
);

ALTER TABLE `product_characteristic` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `product_characteristic` ADD FOREIGN KEY (`characteristic_id`) REFERENCES `characteristic` (`id`);


CREATE TABLE `product_policy` (
  `product_id` int,
  `policy_id` int,
  PRIMARY KEY (`product_id`, `policy_id`)
);

ALTER TABLE `product_policy` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `product_policy` ADD FOREIGN KEY (`policy_id`) REFERENCES `policy` (`id`);
