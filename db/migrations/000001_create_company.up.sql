CREATE TABLE companies (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    repesentative VARCHAR(255),
    tel VARCHAR(255),
    post_code VARCHAR(255),
    address VARCHAR(255),
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp
);
