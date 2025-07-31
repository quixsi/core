# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a monorepo with both Go backend and UI components using the following structure:

- `cmd/` - Go application entry points
  - `server/` - Backend API server
  - `cli/` - Command-line tools
- `internal/` - Private Go application code
  - `server/` - Server implementation
  - `config/` - Configuration handling
  - `models/` - Data models
  - `utils/` - Internal utilities
- `ui/` - Root directory containing all UI-related code
  - `apps/portal/` - Main Vue.js application (@quixsi/portal)
  - `packages/components/` - Shared component library (@quixsi/components)

## Development Commands

### Go Backend Commands
```bash
# From root directory
make build        # Build Go binaries (server and cli)
make dev/server   # Run development server
make dev/cli      # Run CLI tool
make test         # Run Go tests
make lint         # Run golangci-lint
make fmt          # Format Go code
make mod          # Tidy and verify Go modules
make deps         # Download Go dependencies
make clean        # Remove build artifacts
make clean/all    # Remove build artifacts and module cache
```

### UI Commands
```bash
# From root directory (using ui/ prefix)
make ui/install    # Install UI dependencies
make ui/dev        # Start UI development servers
make ui/build      # Build UI workspaces
make ui/clean      # Remove UI node_modules
make ui/clean/all  # Remove UI node_modules and dist folders
make ui/dev/container  # Run UI in Docker container

# Portal app specific commands (from ui/apps/portal/)
npm run dev         # Start development server (runs on port 3000)
npm run build       # Type-check and build for production
npm run type-check  # Run Vue TypeScript type checking
npm run preview     # Preview production build
npm run test:e2e    # Run Playwright end-to-end tests
```

### Testing
```bash
# E2E tests (from portal app directory)
npx playwright install              # Install browsers (first run only)
npm run test:e2e                   # Run all e2e tests
npm run test:e2e -- --project=chromium  # Run tests on specific browser
npm run test:e2e -- tests/example.spec.ts  # Run specific test file
npm run test:e2e -- --debug       # Run tests in debug mode
```

### Container Development
```bash
# From ui/ directory
make dev/container  # Run development server in Docker container
```

## Architecture

### Go Backend
- **Language**: Go 1.24
- **Module**: github.com/quixsi/app
- **Structure**: Standard Go project layout with cmd/ and internal/ directories
- **Server**: Backend API server for serving the frontend and API endpoints
- **CLI**: Command-line tools for administration and utilities

### Portal Application (@quixsi/portal)
- **Framework**: Vue 3 with TypeScript
- **Build Tool**: Vite with TailwindCSS v4
- **Routing**: Vue Router with single home route
- **Testing**: Playwright for E2E testing
- **Dev Tools**: Vue DevTools plugin enabled

### Components Package (@quixsi/components)
- **UI Library**: Shared Vue 3 components with TailwindCSS
- **Components**: Button, Input, Checkbox, Card (with sub-components)
- **Dependencies**: Uses reka-ui, lucide-vue-next, class-variance-authority
- **Build**: Exports individual components and styles separately
- **Utilities**: Custom utility functions in `#lib/utils`

### Key Technologies
- **Vue 3** with Composition API and `<script setup>`
- **TypeScript** with strict type checking via vue-tsc
- **TailwindCSS v4** for styling
- **Vite** for build tooling and dev server
- **npm workspaces** for monorepo management

### Development Notes
- The Go server and UI portal are separate applications that can run independently
- The portal app runs on port 3000 by default
- Both UI projects use `@` alias pointing to their respective `src/` directories
- Components package uses internal imports with `#lib/utils` mapping
- The workspace setup allows the portal to import from `@quixsi/components`
- Root Makefile provides unified access to both Go and UI commands via make targets