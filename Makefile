postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root todo

dropdb:
	docker exec -it postgres12 dropdb todo

migrateup:
	migrate -path db/migration/ -database "postgresql://root:admin@localhost:5432/todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:admin@localhost:5432/todo?sslmode=disable" -verbose down

sqlc:
	sqlc generate

sqlcDocker:
	docker run --rm -v D:\henryS\go-todo-system:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdxb dropdb migrateup migratedown sqlc sqlcDocker test