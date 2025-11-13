# docktidy

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

```bash
go install github.com/thommorais/docktidy/cmd/docktidy@latest
```

Or build from source:

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

### Prerequisites

- Go 1.21 or later
- Docker running locally
- [Air](https://github.com/air-verse/air) for hot reload (optional)
- [Cocogitto](https://docs.cocogitto.io/) for conventional commits (optional)

### Running with hot reload

```bash
go install github.com/air-verse/air@latest
air
```

### Contributing

This project follows [Conventional Commits](https://www.conventionalcommits.org/). To set up commit hooks:

```bash
brew install cocogitto
cog install-hook --all
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.

### Building

```bash
go build -o docktidy ./cmd/docktidy
```

### Project Status

This project is in initial setup phase. Core features are being developed incrementally.

## License

MIT
