postgres:
	docker run --name bankdb -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine

createdb:
	docker exec -it bankdb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bankdb dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
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
proto:
	rm -f pb/*.go
	rm -f swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=swagger --openapiv2_opt=allow_merge=true,merge_file_name=bankapi \
    proto/*.proto
	statik -src=./swagger -dest=./swagger
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/PetarGeorgiev-hash/bankapi/db/sqlc Store
evans:
	evans --host localhost --port 9090 -r repl
redis:
	docker run --name redis -p 6379:6379 -d redis:8-alpine

.PHONY:postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock migrateuplocal proto evans redis
 