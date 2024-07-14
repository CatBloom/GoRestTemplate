SET NAMES 'utf8mb4';

CREATE TABLE IF NOT EXISTS `db_name`.`users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL, 
    `email` VARCHAR(255) NOT NULL, 
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

