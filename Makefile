postgres:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -p 5432:5432 -v ~/postgres_data:/data/db -d postgres:14-alpine
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres patient_tracker
startdb:
	docker start postgres
accessdb:
	docker exec -it postgres psql -U postgres patient_tracker
dropdb:
	docker exec -it postgres dropdb patient_tracker
migrate:
	docker pull migrate/migrate
migrateup:
	migrate -path internal/db/migrations -database "postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/db/migrations -database "postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable" -verbose down
migrateforce1:
	migrate -path internal/db/migrations -database "postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable" force 1
migrateforce2:
	migrate -path internal/db/migrations -database "postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable" force 2
migrateforce3:
	migrate -path internal/db/migrations -database "postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable" force 3
test:
	go test -v -cover ./...
server:
	go run ./cmd/patient_tracker
.PHONY: postgres startdb accessdb dropdb migrate migrateup migratedown migrateforce test
