-- +goose Up
-- installing uuid extension start
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- installing uuid extension end
-- updating update_at fields for all tables start
CREATE
OR REPLACE FUNCTION update_last_modified () RETURNS TRIGGER AS 'BEGIN NEW.updated_at = NOW(); RETURN NEW; END;' LANGUAGE plpgsql;

-- +goose Down
DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;

-- updating update_at fields for all tables end
DROP FUNCTION IF EXISTS update_last_modified CASCADE;