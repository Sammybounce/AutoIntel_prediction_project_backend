-- +goose Up
-- predictions start
CREATE TABLE
  IF NOT EXISTS predictions (
    id UUID NOT NULL,
    group_id UUID NOT NULL,
    brand VARCHAR(255) NOT NULL DEFAULT 'N/A',
    model VARCHAR(255) NOT NULL DEFAULT 'N/A',
    year BIGINT NOT NULL,
    future_year BIGINT NOT NULL,
    prediction_model VARCHAR(255) NOT NULL DEFAULT 'N/A',
    predicted_price DOUBLE PRECISION NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
  );

CREATE INDEX IF NOT EXISTS idx_predictions_brand ON predictions (brand);

CREATE INDEX IF NOT EXISTS idx_predictions_model ON predictions (model);

CREATE INDEX IF NOT EXISTS idx_predictions_year ON predictions (year);

CREATE INDEX IF NOT EXISTS idx_predictions_future_year ON predictions (future_year);

CREATE INDEX IF NOT EXISTS idx_predictions_predicted_price ON predictions (predicted_price);

CREATE INDEX IF NOT EXISTS idx_predictions_created_at ON predictions (created_at);

CREATE
OR REPLACE TRIGGER trigger_update_last_modified_on_predictions BEFORE
UPDATE ON predictions FOR EACH ROW EXECUTE PROCEDURE update_last_modified ();

-- predictions end
-- +goose Down
DROP TABLE IF EXISTS predictions CASCADE;