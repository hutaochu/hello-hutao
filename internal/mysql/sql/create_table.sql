create Table IF NOT EXISTS `user_tab` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "id",
  `name` VARCHAR(128) NOT NULL COMMENT "user name",
  `email` VARCHAR(256) NOT NULL COMMENT "user email",
  `password` VARCHAR(128) NOT NULL COMMENT "user login password",
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "create time",
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT "update time",
  INDEX `name_email` (`name`, `email`)
) ENGINE=InnoDB CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;