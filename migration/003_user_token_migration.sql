-- +goose Up
-- user token start
CREATE TABLE
  IF NOT EXISTS user_tokens (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    token TEXT NOT NULL,
    expire_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
  );

CREATE INDEX IF NOT EXISTS idx_user_token ON user_tokens (token);

CREATE INDEX IF NOT EXISTS idx_user_token_created_at ON user_tokens (created_at);

CREATE
OR REPLACE TRIGGER trigger_update_last_modified_on_user_tokens BEFORE
UPDATE ON user_tokens FOR EACH ROW EXECUTE PROCEDURE update_last_modified ();

-- user token end
-- +goose Down
DROP TABLE IF EXISTS user_tokens CASCADE;