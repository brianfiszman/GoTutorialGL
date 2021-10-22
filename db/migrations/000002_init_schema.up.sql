CREATE TABLE IF NOT EXISTS users (
    "username" varchar PRIMARY KEY,
    "full_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);