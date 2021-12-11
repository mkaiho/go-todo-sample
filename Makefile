BIN_DIR:=bin
ROOT_PACKAGE:=github.com/mkaiho/go-todo-sample
COMMAND_PACKAGES:=$(shell go list ./cmd/...)
BINARIES := $(COMMAND_PACKAGES:$(ROOT_PACKAGE)/cmd/%=$(BIN_DIR)/%)

.PHONY: build
build: $(BINARIES)

$(BINARIES): $(GO_FILES)
	@go build -o $@ $(@:$(BIN_DIR)/%=$(ROOT_PACKAGE)/cmd/%)

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	@rm -rf ./bin
