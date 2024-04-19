CREATE TABLE invoices (
    id INT PRIMARY KEY AUTO_INCREMENT,
    company_id INT NOT NULL,
    customer_id INT NOT NULL,
    issue_date DATE NOT NULL,
    payment DECIMAL(15, 2),
    fee DECIMAL(12, 2),
    fee_rate DECIMAL(5, 2),
    tax DECIMAL(12, 2),
    tax_rate DECIMAL(5, 2),
    billing_amount DECIMAL(15, 2),
    deadline DATE,
    status INT NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    created_at DATETIME default current_timestamp,
    updated_at DATETIME default current_timestamp
);
