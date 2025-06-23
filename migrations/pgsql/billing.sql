CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS wallets
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id       UUID NOT NULL,
    currency_code VARCHAR(10) NOT NULL,
    balance       DOUBLE PRECISION DEFAULT 0.0,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS wallets;