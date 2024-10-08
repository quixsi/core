include ./Makefile.Common

# Rules
$(TOOLS_DIR):
	mkdir -p $@

# Funcs
.PHONY: check-fmt
check-fmt: fmt
	@git diff -s --exit-code *.go || (echo "Build failed: a go file is not formated correctly. Run 'make fmt' and update your PR." && exit 1)

.PHONY: fmt
fmt:
	go fmt ./...
	$(TEMPL) fmt .

.PHONY: govet
govet:
	go vet ./...

.PHONY: build
build: generate
	go build -v -o $(LETS_PARTY) $(ROOT_DIR)/cmd/server/main.go

.PHONY: generate
generate:
	$(TEMPL) generate

.PHONY: compilecheck
compilecheck:
	$(GO_ENV)
	go build -v ./...

.PHONY: run
run: fmt build
	$(LETS_PARTY)

.PHONY: gotest
gotest: 
	$(GO_ENV)
	go test -v ./... -failfast

.PHONY: localtest
localtest: fmt govet check-fmt
	$(GO_ENV)
	go test -v ./... -failfast

.PHONY: gomoddownload
gomoddownload:
	go mod download -x

.PHONY: tools
tools: $(TOOLS_DIR)
	GOBIN=$(TOOLS_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VERSION)
	GOBIN=$(TOOLS_DIR) go install github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION)

.PHONY: golint
golint:
	$(LINT) run --verbose --allow-parallel-runners --timeout=10m 

.PHONY: gotidy
gotidy:
	go mod tidy -compat=$(GO_VERSION)

.PHONY: check-licensehead
check-licensehead:
	@is_successful=true; \
	for f in $(LICENSED_FILES); do \
		first_line=$$(sed -n '1p' "$$f"); \
		second_line=$$(sed -n '2p' "$$f"); \
		if [ "$$first_line" != "$(LICENSEHEAD_FIRST_LINE)" ] || [ "$$second_line" != "$(LICENSEHEAD_SECOND_LINE)" ]; then \
			echo "\x1b[31mError: License header mismatch in $$f\x1b[0m"; \
			is_successful=false; \
		else \
			echo "Check: License header checked in $$f"; \
		fi; \
	done; \
	if [ "$$is_successful" != true ]; then \
		exit 1; \
	fi

.PHONY: licensehead
licensehead: deletehead
	@for f in $(LICENSED_FILES); do \
			first_line=$$(sed -n '1p' "$$f"); \
			second_line=$$(sed -n '2p' "$$f"); \
			if [ "$$first_line" != "$(LICENSEHEAD_FIRST_LINE)" ] || [ "$$second_line" != "$(LICENSEHEAD_SECOND_LINE)" ]; then \
				echo "Found: License header mismatch in $$f"; \
				sed -i "1i$(LICENSEHEAD_FIRST_LINE)" "$$f"; \
				sed -i "2i$(LICENSEHEAD_SECOND_LINE)\n" "$$f"; \
				echo "Written: License header written in $$f"; \
			else \
				echo "Check: License header checked in $$f"; \
			fi \
		done; \
		echo "License headers written successfully."

.PHONY: deletehead
deletehead:
	@for f in $(LICENSED_FILES); do \
			if [ "$$(sed -n '1p' "$$f" | grep '^//')" ] && [ "$$(sed -n '2p' "$$f" | grep '^//')" ]; then \
				if [ -z "$$(sed -n '3p' "$$f")" ]; then \
					sed -i '1,3d' "$$f"; \
				else \
					sed -i '1,2d' "$$f"; \
				fi; \
				echo "Deleted: License headers deleted in $$f"; \
			else \
				echo "Exit: There were no license headers found in $$f"; \
			fi; \
		done; \
		echo "License headers deleted successfully."
