CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS admin;

CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    abbr VARCHAR NOT NULL,
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);

CREATE TABLE IF NOT EXISTS user_devices(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id VARCHAR NOT NULL,
    device_id VARCHAR NOT NULL,
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);