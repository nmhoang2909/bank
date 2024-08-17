ALTER TABLE `accounts` DROP INDEX `accounts_index_1`;

ALTER TABLE `accounts` DROP CONSTRAINT `accounts_ibfk_1`;

DROP TABLE IF EXISTS `users`;
