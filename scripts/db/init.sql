-- Create the user schema
CREATE SCHEMA IF NOT EXISTS users;

CREATE EXTENSION pgcrypto;

-- Create the users table
CREATE TABLE IF NOT EXISTS users.users
(
    id         UUID        DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT users_pk PRIMARY KEY,
    first_name TEXT                                  NOT NULL,
    last_name  TEXT                                  NOT NULL,
    nickname   TEXT                                  NOT NULL,
    password   TEXT                                  NOT NULL,
    email      TEXT UNIQUE                           NOT NULL,
    country    CHAR(2)                               NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()             NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW()             NOT NULL
);

-- Create unique index on ID
CREATE UNIQUE INDEX users_id_uindex
    on users.users (id);

-- Create function to automatically set updated at when a row is changed
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger for the above rule
CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON users.users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- Seed database
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ('jack', 'hughes', 'jack-hughes', crypt('jack-test-pw', gen_salt('bf', 8)), 'jack@test.com', 'GB');
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ('jane', 'doe', 'jane-doe', crypt('jane-test-pw', gen_salt('bf', 8)), 'jane@test.com', 'US');
INSERT INTO users.users (first_name, last_name, nickname, password, email, country)
VALUES ('john', 'smith', 'john-smith', crypt('john-test-pw', gen_salt('bf', 8)), 'john@test.com', 'UA');
