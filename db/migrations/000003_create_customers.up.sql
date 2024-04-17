CREATE TABLE customers (
    id INT PRIMARY KEY,
    company_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    repesentative VARCHAR(255),
    tel VARCHAR(255),
    post_code VARCHAR(255),
    address VARCHAR(255),
    FOREIGN KEY (company_id) REFERENCES companies(id),
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp
);

CREATE TABLE banks (
    id INT PRIMARY KEY,
    customer_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    account_number VARCHAR(20) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp
)
