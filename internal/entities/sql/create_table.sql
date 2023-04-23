create Table IF NOT EXISTS `user_tab` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "id",
  `name` VARCHAR(128) NOT NULL COMMENT "user name",
  `email` VARCHAR(256) UNIQUE NOT NULL COMMENT "user email",
  `avatar_url` VARCHAR(256) COMMENT "user avatar",
  `password` VARCHAR(128) NOT NULL COMMENT "user login password",
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "create time",
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "update time",
  INDEX `name_email` (`name`, `email`)
) ENGINE=InnoDB CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
ALTER TABLE `user_tab` AUTO_INCREMENT = 100000000;