# coqodaq
Restaurant Booking API


#### Migrations
    - goose postgres "user=postgres dbname=coqodaq sslmode=disable" status
    - goose postgres postgres://postgres:@localhost:5431/coqodaq up


##### Creating New Model
    - Step 1) Create New Schema
    - Step 2) Create Query
    - Step 3) Run `sqlc generate`


##### DB Models

Example Request Body

```
Create Reservation

{
    name: "Richards",
    party_size: 2,
    time: "2024-08-31T23:00:00Z",
    restaurant_id: 1,
    table_id: 1
}
```