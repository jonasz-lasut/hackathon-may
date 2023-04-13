CREATE TABLE monolith.articles (
    `id` UNSIGNED NOT NULL AUTO_INCREMENT,
    PRIMARY KEY(id),
    `title` VARCHAR(255),
    `author_id` INT UNSIGNED
);
