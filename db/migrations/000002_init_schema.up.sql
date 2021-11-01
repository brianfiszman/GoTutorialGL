CREATE TABLE IF NOT EXISTS users (
    "username" varchar PRIMARY KEY,
    "fullname" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL
);