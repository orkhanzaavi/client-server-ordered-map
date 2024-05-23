.DEFAULT_GOAL := help
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

####################################################################################################
## MAIN COMMANDS
####################################################################################################
help: ## Commands list
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'


install: ## Make a binary to ./bin folder
	go build -o ./bin/server  ./cmd/server/main.go
	go build -o ./bin/client  ./cmd/client/main.go

analyze: ## Run static analyzer
	test -s ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.58.1
	./bin/golangci-lint run -c ./.golangci.yaml ./...

run: ## Run server
	./bin/server

send: ## Run a client to send messages from the commands.example file to the server
	./bin/client --file=commands.example

test: ## Run tests
	test -s ./bin/gotest || go build -o ./bin/gotest github.com/rakyll/gotest
	./bin/gotest -failfast  ./internal/...

.PHONY: mocks
generate-mocks:
	go run github.com/vektra/mockery/v2/
