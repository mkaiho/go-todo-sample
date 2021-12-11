BIN_DIR:=bin
ROOT_PACKAGE:=github.com/mkaiho/go-todo-sample
COMMAND_PACKAGES:=$(shell go list ./cmd/...)
BINARIES := $(COMMAND_PACKAGES:$(ROOT_PACKAGE)/cmd/%=$(BIN_DIR)/%)

.PHONY: build
build: $(BINARIES)

$(BINARIES): $(GO_FILES)
	@go build -o $@ $(@:$(BIN_DIR)/%=$(ROOT_PACKAGE)/cmd/%)

.PHONY: dev-deps
dev-deps:
	go get gotest.tools/gotestsum@v1.7.0
	go get github.com/vektra/mockery/v2/.../
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	gotestsum

.PHONY: clean
clean:
	@rm -rf ./bin
