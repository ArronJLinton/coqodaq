-- +goose Up

CREATE TABLE reservations (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  party_size INT NOT NULL,
  time TIMESTAMPTZ NOT NULL,
  restaurant_id serial references restaurants(id) NOT NULL,
  table_id serial references tables(id) NOT NULL,
  user_id serial references users(id) NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE reservations;