.PHONY: lint
lint: tools-install
	golangci-lint run ./...

.PHONY: tools-install
tools-install:
	./_tools/install.sh
