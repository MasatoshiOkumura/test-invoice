CREATE TABLE users (
    id INT PRIMARY KEY,
    company_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp
);
