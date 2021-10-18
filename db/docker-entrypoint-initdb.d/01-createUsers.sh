#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE EXTENSION "pgcrypto";
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    DROP TABLE IF EXISTS Users;

    CREATE TABLE Users (
        UserID UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
        Account VARCHAR(50) NOT NULL,
        Password BYTEA NOT NULL,
        Name VARCHAR(20) NOT NULL,
        Disable BOOLEAN DEFAULT FALSE NOT NULL,
        CreateAt    TIMESTAMP DEFAULT NOW() NOT NULL,
        UpdateAt    TIMESTAMP DEFAULT NOW() NOT NULL,
        DeleteAt    TIMESTAMP
    );
EOSQL