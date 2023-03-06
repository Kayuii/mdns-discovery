GOCMD:=$(shell which go)
GOLINT:=$(shell which golint)
GOIMPORT:=$(shell which goimports)
GOFMT:=$(shell which gofmt)
GOBUILD:=$(GOCMD) build
GOINSTALL:=$(GOCMD) install
GOCLEAN:=$(GOCMD) clean
GOTEST:=$(GOCMD) test
GOGET:=$(GOCMD) get
GOLIST:=$(GOCMD) list
GOVET:=$(GOCMD) vet
GOMOD:=$(GOCMD) mod
GOPATH:=$(shell $(GOCMD) env GOPATH)
u := $(if $(update),-u)

BINARY_NAME:=mdnscli
PACKAGES:=$(shell $(GOLIST))
GOFILES:=$(shell find . -name "*.go" -type f)

export GO111MODULE := on

all: test build

mini: test build-mini

.PHONY: build
build: deps
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/

.PHONY: build-mini
build-mini: deps
	$(GOBUILD) -ldflags "-s -w" -o $(BINARY_NAME)-mini ./cmd/

.PHONY: build-static
build-static: deps
	CGO_ENABLED=0 $(GOBUILD) -ldflags '-linkmode "external" -extldflags "-static" -w -s ' -o $(BINARY_NAME)-static ./cmd/

.PHONY: install
install: deps
	$(GOINSTALL) ./cmd/

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: deps
deps:
	$(GOGET) github.com/grandcat/zeroconf
	$(GOGET) github.com/cenkalti/backoff
	$(GOGET) github.com/urfave/cli/v2
	$(GOGET) golang.org/x/net

.PHONY: devel-deps
devel-deps:
	GO111MODULE=off $(GOGET) -v -u \
		golang.org/x/lint/golint

.PHONY: lint
lint: devel-deps
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: vet
vet: deps devel-deps
	$(GOVET) $(PACKAGES)

.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
