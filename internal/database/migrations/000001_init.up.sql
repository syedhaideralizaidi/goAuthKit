CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       username VARCHAR(255) NOT NULL,
                       phone_number VARCHAR(20),
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(50) NOT NULL,
                       is_verified BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Ensure the email and username are unique
CREATE UNIQUE INDEX ON "users" ("email");
CREATE UNIQUE INDEX ON "users" ("username");

-- Add check constraint to ensure role is one of admin, seller, or buyer
ALTER TABLE "users"
    ADD CONSTRAINT "check_role"
        CHECK ("role" IN ('admin', 'seller', 'buyer'));

-- Indexes for common lookup fields
CREATE INDEX ON "users" ("phone_number");
CREATE INDEX ON "users" ("role");

-- Comments for clarity
COMMENT ON COLUMN "users"."password" IS 'hashed password';
COMMENT ON COLUMN "users"."role" IS 'role-based access: admin, seller, or buyer';

-- source: user.sql

ALTER TABLE users ADD COLUMN reset_token TEXT;
ALTER TABLE users ADD COLUMN reset_token_expiry TIMESTAMPTZ;