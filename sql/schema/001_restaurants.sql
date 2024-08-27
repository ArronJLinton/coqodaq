-- +goose Up

CREATE TABLE restaurants (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  dietary_restrictions TEXT[] NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE restaurants; 