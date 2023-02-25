start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

run:
	go run main.go

build:
	@if [ -z "$(APP_NAME)" ]; then \
		export $(grep -v '^#' .env | xargs) && \
		go build -ldflags="-X 'main.AppName=$${APP_NAME}'" -o build/$${APP_NAME}; \
	else \
		go build -ldflags="-X 'main.AppName=$(APP_NAME)'" -o build/$(APP_NAME); \
	fi
