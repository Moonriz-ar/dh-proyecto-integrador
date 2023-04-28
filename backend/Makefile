build:
	if [ ! -d "bin" ]; then mkdir -p bin; fi && go build -o bin main.go
start:
	./bin/main
test: # ./... argument to run unit test in all packages
	go test -v -cover ./...

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root rentcar
dropdb:
	docker exec -it postgres12 dropdb rentcar
accessdb:
	docker exec -it postgres12 psql -U root rentcar
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rentcar?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/rentcar?sslmode=disable" -verbose down
sqlc:
	sqlc generate
mockgen:
	mockgen -package mockdb -destination db/mock/store.go proyecto-integrador/db/sqlc Store

.PHONY: build start test postgres createdb dropdb accessdb migrateup migratedown sqlc mockgen