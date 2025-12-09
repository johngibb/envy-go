
install: ## Builds using go install
    go install ./...

build: ## Builds the binary into out/johngibb/envy-go
    go build ./... -o out/johngibb/envy-go

test: ## Runs all tests
    go test ./...

run: ## Runs the application with go run
    go run .

help:
    @grep -E '^[0-9a-zA-Z_-]+:.*?## .*36325'  | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", 363251, 363252}'
