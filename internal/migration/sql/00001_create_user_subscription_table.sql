-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_subscription (
  id UUID,
  user_id UUID NOT NULL,
  service_name TEXT NOT NULL,
  price INT NOT NULL,
  start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_subscription;
-- +goose StatementEnd
