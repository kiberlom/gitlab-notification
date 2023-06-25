swagger:
	find $(PWD) \( -path "*/docs" \) -exec rm -rf {} +
	# go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/main.go -o ./docs --parseInternal

linter:
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.53.3 golangci-lint run -v