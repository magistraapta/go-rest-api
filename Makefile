migrate-up:
	migrate -path=internal/db/migrations -database "postgresql://postgres:root@localhost:5432/golang-test?sslmode=disable" -verbose up

migrate-down:
	migrate -path=internal/db/migrations -database "postgresql://postgres:root@localhost:5432/golang-test?sslmode=disable" -verbose down

.PHONY: migrate-up migrate-down