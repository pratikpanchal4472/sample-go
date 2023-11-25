BINDIR := $(CURDIR)/bin
GOLANGCI_LINT_VERSION := v1.55.2

all: bin

.PHONY: bin
bin:
	GOBIN=$(BINDIR) go install ./...

# linting
.golangci-bin:
	@echo "===> Installing golangci-lint <==="
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $@ $(GOLANGCI_LINT_VERSION)

.PHONY: golangci
golangci: .golangci-bin
	@echo "===> Running golangci <==="
	.golangci-bin/golangci-lint run -c .golangci.yml


.PHONY: golangci-fix
golangci-fix: .golangci-bin
	@echo "===> Running golangci-fix <==="
	.golangci-bin/golangci-lint run -c .golangci.yml --fix

.PHONY: unit-tests
unit-tests:
	@echo "===> Running unit tests <==="
	go test ./...

.PHONY: check
check: golangci-fix unit-tests

.PHONY: clean
clean:
	rm -rf .golangci-bin
