all:
	go run cmd/main.go
start:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
stop:
	docker stop todo-db
migrate:
	migrate -path schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up