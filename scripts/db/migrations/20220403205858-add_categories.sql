
-- +migrate Up

CREATE TABLE categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    discount TINYINT NOT NULL DEFAULT 0
) ENGINE=INNODB;

-- +migrate Down

DROP TABLE categories
