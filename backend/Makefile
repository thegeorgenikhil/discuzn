# default args if not given
user = postgres
pass = postgres
db = discuzn-db
container_name = discuzn-postgres
db_uri = postgres://$(user):$(pass)@localhost:5432/$(db)?sslmode=disable

migrate = $(HOME)/go/bin/migrate

hello:
	echo "hello"
run:
	go run main.go
fmt:
	gofmt -s -w .
psql-start:
	docker run --name $(container_name) -e POSTGRES_PASSWORD=$(pass) -e POSTGRES_DB=$(db) -p 5432:5432 -d postgres
psql-stop:
	docker stop $(container_name)
psql-rm:
	docker rm $(container_name)
psql-kill:
	docker stop $(container_name)
	docker rm $(container_name)
migrate-up:
	$(migrate) -database "$(db_uri)" -path db/migrations up
