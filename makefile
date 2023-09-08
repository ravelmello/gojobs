.PHONY: default run build test docs clean

APP_NAME=gojobs
#Tasks
default: run
run-with-docs:
	@swag init
	@go run main.go
run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm rf ./docs
