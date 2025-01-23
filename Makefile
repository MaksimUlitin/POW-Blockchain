
run:
	go run cmd/main.go

docker-build:
	docker build -t blockchain-app .

docker-run:
	docker-compose up

all: run
