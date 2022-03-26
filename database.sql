CREATE DATABASE notes;

DROP TABLE IF EXISTS `note`;
CREATE TABLE `note` (
    `id` bigint unsigned NOT NULL auto_increment primary key,
    `user_id` bigint unsigned NOT NULL,
    `title` varchar(120) NOT NULL,
    `description` varchar(255) NOT NULL,
    `archived` tinyint(1) NOT NULL DEFAULT "0",
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);