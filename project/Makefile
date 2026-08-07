.PHONY: all build proto clean test run-control-plane run-agent

PROTO_DIR := proto
PROTO_OUT := internal/proto

all: proto build

proto:
	protoc \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		-I $(PROTO_DIR) \
		$(PROTO_DIR)/*.proto

build:
	go build -o bin/control-plane ./cmd/control-plane
	go build -o bin/agent ./cmd/agent
	go build -o bin/orchctl ./cmd/orchctl

clean:
	rm -rf bin/ data/

test:
	go test ./...

run-control-plane:
	go run ./cmd/control-plane

run-agent:
	go run ./cmd/agent
