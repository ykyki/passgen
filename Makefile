BINDIR := .
BINNAME := passgen
BIN := $(BINDIR)/$(BINNAME)

GO_FILES := $(shell find . -type f -name '*.go' -print) # TODO test ファイルをどうするか

.PHONY: build
build: $(BIN)

.PHONY: clean
clean:
	@$(RM) $(BIN)

$(BIN): $(GO_FILES)
	@go build -o $@ ./...

.PHONY: run
run:
	@go run ./... $(OPTIONS)