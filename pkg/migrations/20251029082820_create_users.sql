-- migrate:up
-- uuid-ossp extension is required for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DATABASE sqlc_go;

CREATE TYPE user_role AS ENUM ('admin', 'user', 'guest');

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(25) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    role user_role NOT NULL DEFAULT 'user',
    -- country_code CHAR(2) NOT NULL,
    properties JSONB,
    password_hash TEXT NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_role;