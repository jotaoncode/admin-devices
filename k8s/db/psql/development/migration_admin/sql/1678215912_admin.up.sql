DROP TABLE IF EXISTS admin;

CREATE TABLE IF NOT EXISTS devices(
    id SERIAL PRIMARY KEY,
    ip VARCHAR NOT NULL UNIQUE,
    mac VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);