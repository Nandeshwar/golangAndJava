#Create database

CREATE DATABASE if NOT EXISTS quotedb;
GRANT ALL PRIVILEGES ON quotedb.* TO 'user1'@'%';
FLUSH PRIVILEGES;

use quotedb;

CREATE TABLE IF NOT EXISTS quotedb.`event`  (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `day` INTEGER NOT NULL,
    `month` INTEGER NOT NULL,
    `year` INTEGER NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `description` VARCHAR(1000) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS quotedb.`event_link` (
	`id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `link_id` INTEGER,
    `link` VARCHAR(100),
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(link_id) REFERENCES event(ID)
);


