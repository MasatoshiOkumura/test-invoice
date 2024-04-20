INSERT INTO companies (
    name,
    repesentative,
    tel,
    post_code,
    address
) VALUES (
    'test-company',
    'test-repesentative',
    '0312345678',
    '151-0001',
    '東京都千代田区千代田9-9'
);

INSERT INTO customers (
    company_id,
    name,
    repesentative,
    tel,
    post_code,
    address
) VALUES (
    1,
    'test-customer',
    'test-customer-repesentative',
    '090-1234-5678',
    '111-2222',
    'カスタマーテスト住所'
);

INSERT INTO banks (
    customer_id,
    name,
    account_number,
    account_name
) VALUES (
    1,
    'test-bank',
    '0987654',
    'テスト口座名'
);
