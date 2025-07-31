# Quixsi Application

A monorepo containing a Go backend API server and Vue.js frontend portal application.

## Project Structure

```
.
├── cmd/                    # Go application entry points
│   ├── server/            # Backend API server
│   └── cli/               # Command-line tools
├── internal/              # Private Go application code
│   ├── server/            # Server implementation
│   ├── config/            # Configuration handling
│   ├── models/            # Data models
│   └── utils/             # Internal utilities
└── ui/                    # UI-related code
    ├── apps/portal/       # Main Vue.js application
    └── packages/components/ # Shared component library
```

## Prerequisites

- **Go**: Version 1.24 or later
- **Node.js**: Version 18 or later
- **npm**: Latest version

## Quick Start

### 1. Install Dependencies

#### Go Dependencies
```bash
go mod download
```

#### UI Dependencies
```bash
make ui/install
# or manually:
cd ui && npm install
```

### 2. Run Development Servers

#### Option A: Run Both (Recommended)
```bash
# Terminal 1: Start backend server
make dev/server

# Terminal 2: Start frontend portal
make ui/dev
```

#### Option B: Run Individually

**Backend Only:**
```bash
make dev/server
# or: go run ./cmd/server
```

**Frontend Only:**
```bash
make ui/dev
# or: cd ui/apps/portal && npm run dev
```

### 3. Access Applications

- **Frontend Portal**: http://localhost:3000
- **Backend API**: http://localhost:8080 (adjust port as configured)

## Development Commands

### Go Backend

| Command | Description |
|---------|-------------|
| `make build` | Build Go binaries (server and cli) |
| `make dev/server` | Run development server |
| `make dev/cli` | Run CLI tool |
| `make test` | Run Go tests |
| `make lint` | Run golangci-lint |
| `make fmt` | Format Go code |
| `make mod` | Tidy and verify Go modules |
| `make deps` | Download Go dependencies |
| `make clean` | Remove build artifacts |

### Frontend UI

| Command | Description |
|---------|-------------|
| `make ui/install` | Install UI dependencies |
| `make ui/dev` | Start UI development servers |
| `make ui/build` | Build UI workspaces |
| `make ui/clean` | Remove UI node_modules |
| `make ui/dev/container` | Run UI in Docker container |

### Portal App Specific

From `ui/apps/portal/` directory:

| Command | Description |
|---------|-------------|
| `npm run dev` | Start development server (port 3000) |
| `npm run build` | Type-check and build for production |
| `npm run type-check` | Run Vue TypeScript type checking |
| `npm run preview` | Preview production build |
| `npm run test:e2e` | Run Playwright end-to-end tests |

## Testing

### End-to-End Tests

```bash
# First-time setup (install browsers)
cd ui/apps/portal
npx playwright install

# Run all E2E tests
npm run test:e2e

# Run tests on specific browser
npm run test:e2e -- --project=chromium

# Run specific test file
npm run test:e2e -- tests/example.spec.ts

# Debug mode
npm run test:e2e -- --debug
```

### Go Tests

```bash
make test
# or: go test ./...
```

## Production Build

### Backend

```bash
make build
# Binaries will be created in bin/ directory
./bin/server  # Run the server
./bin/cli     # Run CLI tools
```

### Frontend

```bash
make ui/build
# or: cd ui/apps/portal && npm run build
```

## Container Deployment

### Individual Services

**Build and run server container:**
```bash
podman build -f Containerfile.server -t quixsi-server .
podman run -p 8080:8080 quixsi-server
```

**Build and run UI container:**
```bash
podman build -f Containerfile.ui -t quixsi-ui .
podman run -p 8080:8080 quixsi-ui
```

### Kubernetes Deployment

**Deploy to Kubernetes cluster:**
```bash
# Apply all manifests
kubectl apply -f deploy/

# Or apply individually
kubectl apply -f deploy/server-configmap.yaml
kubectl apply -f deploy/server-deployment.yaml
kubectl apply -f deploy/server-service.yaml
kubectl apply -f deploy/ui-deployment.yaml
kubectl apply -f deploy/ui-service.yaml
kubectl apply -f deploy/network-policies.yaml
```

**Check deployment status:**
```bash
kubectl get pods -l app=quixsi-server
kubectl get pods -l app=quixsi-ui
kubectl get services
```

**Access services (port-forward for testing):**
```bash
# Access UI
kubectl port-forward service/quixsi-ui 8080:8080

# Access server API
kubectl port-forward service/quixsi-server 8081:8080
```

### Docker Development (Legacy)

Run the UI in a Docker container:

```bash
cd ui
make dev/container
```

## Technology Stack

### Backend
- **Go 1.24**
- **Module**: github.com/quixsi/app
- Standard Go project layout

### Frontend
- **Vue 3** with Composition API and TypeScript
- **Vite** build tool
- **TailwindCSS v4** for styling
- **Vue Router** for routing
- **Playwright** for E2E testing

### Components Library
- **Vue 3** components with TailwindCSS
- **reka-ui**, **lucide-vue-next**, **class-variance-authority**
- Shared across applications via npm workspaces

## Project Architecture

- **Monorepo**: Uses npm workspaces for UI package management
- **Independent Services**: Backend and frontend can run separately
- **Shared Components**: UI components library shared between applications
- **TypeScript**: Strict type checking enabled
- **Modern Tooling**: Vite for fast development and builds

## Troubleshooting

### Common Issues

1. **Port Conflicts**: Ensure ports 3000 (frontend) and 8080 (backend) are available
2. **Node Version**: Use Node.js 18+ for compatibility
3. **Go Version**: Ensure Go 1.24+ is installed
4. **Dependencies**: Run `go mod download` and `make ui/install` if packages are missing

### Build Issues

If builds fail:
```bash
# Clean and reinstall
make clean/all
make ui/clean/all
go mod download
make ui/install
```

## Contributing

1. Follow existing code conventions
2. Run tests before submitting changes
3. Use the provided linting tools (`make lint`, `npm run type-check`)
4. Keep commits focused and descriptive

For more detailed development guidance, see `CLAUDE.md`.