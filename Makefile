.PHONY: install-hooks lint

# Install Lefthook and configure it
install-hooks:
	go install github.com/evilmartians/lefthook@latest
	lefthook install

# Run GolangCI-Lint
lint:
	golangci-lint run ./...
