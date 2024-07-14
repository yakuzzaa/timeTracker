-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    passport_number VARCHAR(50) NOT NULL,
    passport_series VARCHAR(50) NOT NULL,
    surname VARCHAR(50),
    name VARCHAR(50),
    patronymic VARCHAR(50),
    address TEXT

);


-- +goose Down
DROP TABLE users;
