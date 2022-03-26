CREATE DATABASE notes;

CREATE TABLE `note` (
    `id` bigint auto_increment primary key,
    `user_id` bigint not null,
    `title` varchar(120) not null,
    `description` varchar(255) not null,
    `archived` tinyint(1) not null default '0',
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
)