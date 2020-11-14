CREATE TABLE users (id serial primary key, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, hashed_password CHAR(60) NOT NULL, created timestamp(0) NOT NULL, active BOOLEAN NOT NULL DEFAULT TRUE);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
