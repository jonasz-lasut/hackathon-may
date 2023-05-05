-- Create "articles" table
CREATE TABLE `articles` (`id` int unsigned NOT NULL AUTO_INCREMENT, `title` varchar(255) NULL, `author_id` int unsigned NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
