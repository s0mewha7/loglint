# Makefile для loglint

LINTER=loglint
TESTDATA=testdata/src

.PHONY: all build run clean

all: build run

build:
	@echo "==> Building loglint..."
	go build -o $(LINTER) ./cmd/loglint

run:
	@echo "==> Running loglint on testdata..."
	-go vet -vettool=./$(LINTER) $(TESTDATA)/slog/slog.go
	-go vet -vettool=./$(LINTER) $(TESTDATA)/uberzap/zap.go

clean:
	@echo "==> Cleaning..."
	rm -f $(LINTER)