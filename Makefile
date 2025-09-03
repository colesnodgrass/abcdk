.PHONY: gen
gen:
	go tool github.com/atombender/go-jsonschema -p protocol https://raw.githubusercontent.com/airbytehq/airbyte-protocol/refs/heads/main/protocol-models/src/main/resources/airbyte_protocol/v0/airbyte_protocol.yaml > protocol/protocol.go

.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w"

.PHONY: test
test:
	go test ./...

.PHONY: docker
docker: build
	docker buildx build --platform linux/amd64,linux/arm64 .
