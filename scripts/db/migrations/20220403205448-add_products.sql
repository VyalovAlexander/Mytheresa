
-- +migrate Up

CREATE TABLE products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    price BIGINT NOT NULL DEFAULT 0,
    category_id BIGINT NOT NULL DEFAULT 0,
    discount TINYINT NOT NULL DEFAULT 0
) ENGINE=INNODB;

-- +migrate Down

DROP TABLE products;