# rssagg

RSS aggregator for multiple feeds.

## Important Tools

- `sqlc` - for generating type-safe SQL queries

    ``` bash
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```

- `goose` - for database migrations

    ``` bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```

## Important Commands

- To up the database

    ``` bash
    # MOVE TO directory 
    cd sql/schema/

    # UP the database
    goose postgres <connection-string> up
    ```

- To down the database

    ``` bash
    # MOVE TO directory 
    cd sql/schema/

    # DOWN the database
    goose postgres <connection-string> down
    ```

## Technologies

- Golang
- chi
- godotenv
