CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS tables (
    id SERIAL PRIMARY KEY,
    restaurant_id INT REFERENCES restaurants(id),
    capacity INT NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    number INT NOT NULL
);

CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    table_id INT REFERENCES tables(id),
    user_id INT REFERENCES users(id),
    status VARCHAR(50) NOT NULL
);

ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE restaurants ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE tables ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE bookings ADD COLUMN deleted_at TIMESTAMP;