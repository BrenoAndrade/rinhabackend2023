dev:
	go run ./cmd/rinha.go

build: clean deps
	CGO_ENABLED=0 go build -pgo=./cmd/cpu.pprof -v -o ./bin/rinha ./cmd/rinha.go

deps:
	go mod tidy

clean:
	rm -rf ./bin

docker:
	docker build -t brenoandrader/rinha-go .

docker-push:
	docker buildx build --push --platform linux/amd64 --tag brenoandrader/rinha-go .

docker-down:
	docker compose -f ./build/docker-compose.local.yml down -v --remove-orphans

docker-local: docker-down
	docker compose -f ./build/docker-compose.local.yml up --build -d