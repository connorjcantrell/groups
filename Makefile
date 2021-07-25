postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root groups

dropdb:
	docker exec -it postgres13 dropdb groups

migrateup:
	migrate -path postgres/migration -database "postgresql://root:secret@localhost:5432/groups?sslmode=disable" -verbose up
    
migratedown:
	migrate -path postgres/migration -database "postgresql://root:secret@localhost:5432/groups?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown