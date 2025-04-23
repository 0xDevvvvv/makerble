-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('doctor', 'receptionist')),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_passwords (
    id INT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    password_hash TEXT NOT NULL
);

CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    gender VARCHAR(10),
    address TEXT,
    phone VARCHAR(15),
    illness TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_passwords;
DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
