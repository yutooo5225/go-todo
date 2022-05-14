# Todo list server in GO

## `sqlc` for Windows

1. Run `docker pull kjconroy/sqlc` to pull `sqlc` docker image

2. Replace the `<repo_path>` with **your absolute repo path** in `Makefile`:

    ```makefile
    sqlcDocker:
     docker run --rm -v <repo_path>:/src -w /src kjconroy/sqlc generate

    ```

3. Use `make sqlcDocker` to generate sqlc files
