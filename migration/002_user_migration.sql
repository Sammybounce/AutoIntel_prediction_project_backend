-- +goose Up
-- user start
CREATE TABLE
  IF NOT EXISTS users (
    id UUID NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(60) NOT NULL,
    u_password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    UNIQUE (email)
  );

CREATE INDEX IF NOT EXISTS idx_user_first_name ON users (first_name);

CREATE INDEX IF NOT EXISTS idx_user_last_name ON users (last_name);

CREATE INDEX IF NOT EXISTS idx_user_email ON users (email);

CREATE INDEX IF NOT EXISTS idx_user_password ON users (u_password);

CREATE INDEX IF NOT EXISTS idx_User_created_at ON users (created_at);

CREATE
OR REPLACE TRIGGER trigger_update_last_modified_on_users BEFORE
UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_last_modified ();

-- user end
-- +goose Down
DROP TABLE IF EXISTS users CASCADE;