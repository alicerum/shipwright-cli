APP = shp
OUTPUT_DIR ?= _output

CMD = ./cmd/$(APP)/...
PKG = ./pkg/...
BIN = $(OUTPUT_DIR)/$(APP)

GO_FLAGS ?= -v -mod=vendor
GO_TEST_FLAGS ?= -race -cover

ARGS ?=

.PHONY: $(BIN)
$(BIN):
	go build $(GO_FLAGS) -o $(BIN) $(CMD)

build: $(BIN)

# creates a kubectl prefixed shp binary, "kubectl-shp". When placed on user's PATH, works as a
# kubectl plugin, therefore "kubectl shp" becomes available
.PHONY: kubectl
kubectl: BIN = $(OUTPUT_DIR)/kubectl-$(APP)
kubectl: $(BIN)

clean:
	rm -rf "$(OUTPUT_DIR)"

run:
	go run $(GO_FLAGS) $(CMD) $(ARGS)

test: test-unit

.PHONY: test-unit
test-unit:
	go test $(GO_FLAGS) $(GO_TEST_FLAGS) $(CMD) $(PKG) $(ARGS)
