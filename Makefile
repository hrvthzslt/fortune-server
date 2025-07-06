.DEFAULT_GOAL := help

help:
	@grep -h -E '^[a-zA-Z0-9_-]+:.*?# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: # Run with go
	FORTUNE_PORT=8080 go run main.go

build: # Build and run
	go build

start: # Start the compiled binary
	FORTUNE_PORT=8080 ./fortune-server
