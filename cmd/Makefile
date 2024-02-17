build:
	@go build -o bin/go-jwt

run: build
	@./bin/go-jwt

test:
	@go test -v ./...

postgresinit:
	docker run --name jwt-postgres -e POSTGRES_PASSWORD=test -p 8001:5432 -d postgres

dockerstart:
	docker start jwt-postgres