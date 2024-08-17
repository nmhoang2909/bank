CREATE TABLE IF NOT EXISTS `users` (
  `username` varchar(255) PRIMARY KEY,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `hash_password` varchar(255) NOT NULL,
  `password_changed_at` timestamp NOT NULL DEFAULT (now()),
  `created_at` timestamp DEFAULT (now())
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`username`);

CREATE UNIQUE INDEX `accounts_index_1` ON `accounts` (`owner`, `currency`);
