postgres:
	docker run --name bankdb -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

createdb:
	docker exec -it bankdb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bankdb dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:eR9rI95v.2MxU%2Aq%7CQSY%24w%3E%23%5B4mk%3A@simple-bank.ch6co8eigbcb.eu-central-1.rds.amazonaws.com:5432/simple_bank" -verbose up
migrateuplocal:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/PetarGeorgiev-hash/bankapi/db/sqlc Store
.PHONY:postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock migrateuplocal
 