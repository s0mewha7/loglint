LINTER=loglint
TESTDATA=testdata/src

.PHONY: all build test run clean

all: build test run

build:
	@echo "==> Building loglint..."
	go build -o $(LINTER) ./cmd/loglint

test:
	@echo "==> Running tests..."
	go test ./internal/rules

run:
	@echo "==> Running loglint on testdata..."
	-go vet -vettool=./$(LINTER) $(TESTDATA)/slog/slog.go
	-go vet -vettool=./$(LINTER) $(TESTDATA)/uberzap/zap.go

clean:
	@echo "==> Cleaning..."
	rm -f $(LINTER)