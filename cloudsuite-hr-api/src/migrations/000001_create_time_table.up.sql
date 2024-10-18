CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS times (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    date VARCHAR(255) NOT NULL,
    entry_time TIMESTAMP NOT NULL,
    lunch_entry_time TIMESTAMP NOT NULL,
    lunch_exit_time TIMESTAMP NOT NULL,
    exit_time TIMESTAMP NOT NULL
);