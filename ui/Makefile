.PHONY: install
install:
	@npm install --ignore-scripts

.PHONY: dev
dev:
	@npm run -ws --if-present dev

.PHONY: build
build:
	@npm run -ws --if-present build

.PHONY: clean
clean:
	@find . -name 'node_modules' -type d -prune -exec rm -rf '{}' +

.PHONY: clean/all
clean/all:
	@find . -name 'node_modules' -o -name 'dist' -type d -prune -exec rm -rf '{}' +

.PHONY: dev/container
dev/container:
	@docker run -itp 3000:3000 --rm \
    -w /app --mount type=bind,src=$(shell pwd),target=/app \
    node:22.15.1-alpine \
    sh -c "npm ci --include optional && npm -w @quixsi/portal run dev"