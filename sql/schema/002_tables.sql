-- +goose Up

CREATE TABLE tables (
  id SERIAL PRIMARY KEY,
  capacity INT NOT NULL,
  is_reserved BOOLEAN NOT NULL DEFAULT FALSE,
  restaurant_id serial references restaurants(id) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE tables;