# Todo list server in GO

## Check list

### 1. Branch naming

`<Trello ticket number>-<brief-description>`

For example: `TD0001-db-crud-unit-test`

### 2. Tidy go mod

Run `go mod tidy`

### 3. create PR to `dev` branch

---

## `sqlc` for Windows

1. Run `docker pull kjconroy/sqlc` to pull `sqlc` docker image

2. Replace the `<repo_path>` with **your absolute repo path** in `Makefile`:

    ```makefile
    sqlcDocker:
     docker run --rm -v <repo_path>:/src -w /src kjconroy/sqlc generate

    ```

3. Use `make sqlcDocker` to generate sqlc files
