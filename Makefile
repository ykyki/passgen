BINDIR := .
BINNAME := passgen
BIN := $(BINDIR)/$(BINNAME)

GO_FILES := $(shell find . -type f -name '*.go' -print)

.PHONY: build
build: $(BIN)

.PHONY: clean
clean:
	@$(RM) $(BIN)

$(BIN): $(GO_FILES)
	@go build -o $@ $(GO_FILES)

.PHONY: run
run:
	@go run $(GO_FILES) $(OPTIONS)