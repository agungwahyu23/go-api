CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    email VARCHAR(200) UNIQUE NOT NULL,
    username VARCHAR(200) NULL,
    password VARCHAR(200) NULL,
    email_verified_at TIMESTAMP,
    address TEXT,
    phone VARCHAR(20),
    date_of_birth DATE,
    gender INTEGER NULL,
    is_active INTEGER NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
