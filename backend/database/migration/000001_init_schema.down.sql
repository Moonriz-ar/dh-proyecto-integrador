ALTER TABLE `product_policy` DROP FOREIGN KEY `product_policy_ibfk_2`;
ALTER TABLE `product_policy` DROP FOREIGN KEY `product_policy_ibfk_1`;
DROP TABLE IF EXISTS `product_policy`;

ALTER TABLE `product_characteristic` DROP FOREIGN KEY `product_characteristic_ibfk_2`;
ALTER TABLE `product_characteristic` DROP FOREIGN KEY `product_characteristic_ibfk_1`;
DROP TABLE IF EXISTS `product_characteristic`;

ALTER TABLE `image` DROP FOREIGN KEY `image_ibfk_1`;
DROP TABLE IF EXISTS `image`;

DROP TABLE IF EXISTS `product`;

DROP TABLE IF EXISTS `policy`;

DROP TABLE IF EXISTS `characteristic`;

DROP TABLE IF EXISTS `city`;

DROP TABLE IF EXISTS `category`;
