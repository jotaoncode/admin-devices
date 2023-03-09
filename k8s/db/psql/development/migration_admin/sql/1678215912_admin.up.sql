CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS devices(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    ip VARCHAR NOT NULL,
    mac VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at DATE,
    updated_at DATE,
    deleted_at DATE
);