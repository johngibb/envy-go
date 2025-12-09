build: ## Builds the binary into out/envy-go
	go build -o bin/envy-go ./...

test: ## Runs all tests
	go test ./...

modd: ## Runs modd
	modd

help: ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*##' Makefile \
		| awk 'BEGIN {FS = ":.*##"} {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'