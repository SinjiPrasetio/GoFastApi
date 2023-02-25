start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

run:
	go run main.go

build:
	go build