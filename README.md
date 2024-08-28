# coqodaq
Restaurant Booking API

#### Tools
- Migrations - https://github.com/pressly/goose
- SQL Queries - https://docs.sqlc.dev/en/latest/

#### Migrations
    - goose postgres postgres://postgres:@localhost:5431/coqodaq up

#### Creating New Model
    - Step 1) Create New Schema in `sql/schema`
    - Step 2) Run the migration
    - Step 3) Create Queries in `sql/queries`
    - Step 4) Run `sqlc generate`

##### Routes

Create Reservation
- `/api/reservation/create`

Example Request Body
```
{
    name: "Flintstone Family",
    party_size: 2,
    time: "2024-08-31T23:00:00Z",
    restaurant_id: 1,
    table_id: 1
}
```