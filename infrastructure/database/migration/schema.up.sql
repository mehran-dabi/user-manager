CREATE TABLE IF NOT EXISTS users (
    id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    nick_name VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL,
    country VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
