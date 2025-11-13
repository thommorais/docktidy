# docktidy

[![Coverage Status](https://coveralls.io/repos/github/thommorais/docktidy/badge.svg?branch=main)](https://coveralls.io/github/thommorais/docktidy?branch=main)

A TUI (Terminal User Interface) tool for safely cleaning up Docker resources.

## Features

- Interactive terminal UI for managing Docker resources
- Safe pruning with intelligent risk assessment
- Usage history tracking to avoid removing active resources
- Support for containers, images, volumes, and networks
- Dry-run mode to preview changes before applying

## Architecture

docktidy follows hexagonal/ports-and-adapters architecture:

```md
internal/
├── domain/          # Core business logic (framework-independent)
│   ├── resource.go  # Docker resource models
│   ├── prune.go     # Pruning logic
│   └── history.go   # Usage tracking
├── ports/           # Interfaces defining contracts
│   ├── docker.go    # Docker operations interface
│   ├── storage.go   # Storage interface
│   └── ui.go        # UI interface
└── adapters/        # Concrete implementations
    ├── docker/      # Docker client adapter
    ├── sqlite/      # SQLite storage adapter
    └── tui/         # Bubbletea TUI adapter
```

## Tech Stack

- **Go 1.21+** - Core language
- **Bubbletea** - Terminal UI framework
- **Lipgloss** - Terminal styling
- **Bubbles** - TUI components
- **Docker SDK** - Docker API client
- **SQLite** - Local storage
- **Cobra** - CLI framework
- **Viper** - Configuration management

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap thommorais/tap
brew install docktidy
```

### Download Binary

Download the latest release for your platform from [GitHub Releases](https://github.com/thommorais/docktidy/releases).

### Go Install

```bash
go install github.com/thommorais/docktidy/cmd/docktidy@latest
```

### Build from Source

```bash
git clone https://github.com/thommorais/docktidy.git
cd docktidy
go build -o docktidy ./cmd/docktidy
```

## Usage

Start the interactive TUI:

```bash
docktidy
```

## Development

### Quick Start

```bash
# Initial setup (installs tools and dependencies)
make setup

# Start development server with hot reload
make dev

# Run tests
make test

# Build binary
make build
```

Run `make help` to see all available commands.

### Prerequisites

- Go 1.21 or later
- Docker running locally
- Make
- [Air](https://github.com/air-verse/air) for hot reload (auto-installed by `make setup`)
- [Cocogitto](https://docs.cocogitto.io/) for conventional commits (auto-installed by `make setup`)

### Common Development Tasks

```bash
# Start development with hot reload
make dev

# Run tests with coverage
make test-coverage

# Run all checks (format, vet, lint, test)
make check

# Build and run
make run

# Create a commit interactively
make commit
```

### Building

```bash
# Quick build
make build

# Or with Go directly
go build -o docktidy ./cmd/docktidy

# Install to GOPATH/bin
make install
```

### Contributing

This project follows [Conventional Commits](https://www.conventionalcommits.org/).

```bash
# Create a commit interactively
make commit

# Or manually with git
git commit -m "feat: add new feature"
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

### Project Status

This project is in initial setup phase. Core features are being developed incrementally.

### Releases

We use semantic versioning and automated releases. See [RELEASE.md](RELEASE.md) for details.

```bash
# Create a new release (auto-detects version from commits)
make release

# Dry run (see what would happen)
make release-dry

# Build release locally without publishing
make release-snapshot
```

## License

MIT
