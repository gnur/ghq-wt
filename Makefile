VERSION = $(shell grep 'const version' main.go | sed 's/.*"\(.*\)"/\1/')
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X main.revision=$(CURRENT_REVISION)"
VERBOSE_FLAG = $(if $(VERBOSE),-v)
u := $(if $(update),-u)

.PHONY: deps
deps:
	go get ${u} $(VERBOSE_FLAG)
	go mod tidy

.PHONY: devel-deps
devel-deps: deps
	go install github.com/tcnksm/ghr@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: test
test: deps
	go test $(VERBOSE_FLAG) ./...

.PHONY: lint
lint: devel-deps
	staticcheck ./...

.PHONY: build
build: deps
	go build $(VERBOSE_FLAG) -ldflags=$(BUILD_LDFLAGS) -o ghq

.PHONY: install
install: deps
	mkdir -p $(HOME)/bin
	go build $(VERBOSE_FLAG) -ldflags=$(BUILD_LDFLAGS) -o $(HOME)/bin/ghq

.PHONY: release
release: crossbuild upload

CREDITS: go.sum
	@echo "CREDITS generation skipped (godzil removed)"

DIST_DIR = dist/v$(VERSION)
.PHONY: crossbuild
crossbuild: CREDITS
	rm -rf $(DIST_DIR)
	@for os_arch in linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 windows_amd64 windows_arm64; do \
		os=$${os_arch%%_*}; arch=$${os_arch##*_}; \
		ext=""; [ "$$os" = "windows" ] && ext=".exe"; \
		dir=$(DIST_DIR)/ghq_$${os}_$${arch}; \
		mkdir -p $$dir; \
		echo "Building $$os/$$arch..."; \
		env CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -ldflags=$(BUILD_LDFLAGS) -o $$dir/ghq$$ext . ; \
		cp misc/bash/_ghq misc/zsh/_ghq $$dir/ 2>/dev/null || true; \
		(cd $(DIST_DIR) && zip -r ghq_$${os}_$${arch}.zip ghq_$${os}_$${arch}/ && rm -rf ghq_$${os}_$${arch}); \
	done
	cd $(DIST_DIR) && shasum $$(find * -type f -maxdepth 0) > SHASUMS

.PHONY: upload
upload:
	ghr v$(VERSION) $(DIST_DIR)
