# Development Guide

Quick reference for common development tasks.

## Getting Started

```bash
# Initial setup (first time only)
make setup

# Start development
make dev
```

## Common Commands

### Development
```bash
make dev              # Start with hot reload
make build            # Build binary
make run              # Build and run
make clean            # Clean build artifacts
```

### Testing
```bash
make test             # Run tests
make test-coverage    # Run tests with coverage report
make check            # Run all checks (fmt, vet, lint, test)
```

### Code Quality
```bash
make fmt              # Format code
make vet              # Run go vet
make lint             # Run golangci-lint
make ci               # Run all CI checks locally
```

### Dependencies
```bash
make deps             # Download dependencies
make tidy             # Tidy go.mod
make verify           # Verify dependencies
```

### Committing
```bash
make commit           # Interactive conventional commit
git commit -m "feat: add feature"  # Manual commit
```

### Releasing
```bash
make release          # Create and push new release
make release-dry      # Preview release without creating it
make release-snapshot # Build release artifacts locally
make changelog        # Generate changelog
```

### Version
```bash
make version          # Show version info
./tmp/docktidy version  # Show built binary version
```

### Utilities
```bash
make docker-check     # Verify Docker is running
make tools            # Install development tools
make help             # Show all available commands
```

## Project Structure

```
docktidy/
├── cmd/docktidy/          # Application entrypoint
│   └── main.go
├── internal/              # Private application code
│   ├── domain/           # Core business logic
│   ├── ports/            # Interface definitions
│   └── adapters/         # Implementations
│       ├── docker/       # Docker client
│       ├── sqlite/       # Database
│       └── tui/          # Terminal UI
├── pkg/                  # Public libraries (future)
├── .github/workflows/    # CI/CD pipelines
├── tmp/                  # Build output (gitignored)
└── dist/                 # Release builds (gitignored)
```

## Development Workflow

1. **Create a branch** for your feature
   ```bash
   git checkout -b feat/my-feature
   ```

2. **Make changes** with hot reload running
   ```bash
   make dev
   ```

3. **Test your changes**
   ```bash
   make test
   make check
   ```

4. **Commit with conventional commits**
   ```bash
   make commit
   # or
   git commit -m "feat(tui): add resource selection"
   ```

5. **Push and create PR**
   ```bash
   git push origin feat/my-feature
   ```

## Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting
- `refactor`: Code restructuring
- `perf`: Performance improvement
- `test`: Tests
- `build`: Build system
- `ci`: CI/CD
- `chore`: Maintenance

**Scopes:**
- `tui`: Terminal UI
- `docker`: Docker adapter
- `storage`: Storage adapter
- `domain`: Core domain logic
- `cli`: CLI commands

**Examples:**
```bash
feat(tui): add resource filtering
fix(docker): handle connection timeout
docs: update installation guide
refactor(domain): simplify pruning logic
test(storage): add SQLite adapter tests
```

## Release Process

Releases are automated based on conventional commits:

- `fix:` → Patch release (0.0.x)
- `feat:` → Minor release (0.x.0)
- `BREAKING CHANGE:` → Major release (x.0.0)

```bash
# Preview what will happen
make release-dry

# Create the release
make release
```

This will:
1. Analyze commits since last release
2. Calculate next version
3. Update CHANGELOG.md
4. Create git tag
5. Push to GitHub
6. Trigger automated build and release

## Troubleshooting

### Hot reload not working
```bash
# Reinstall Air
go install github.com/air-verse/air@latest
# Ensure it's in PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

### Tests failing
```bash
# Ensure Docker is running
make docker-check
# Run tests with verbose output
go test -v ./...
```

### Linter errors
```bash
# Auto-fix formatting
make fmt
# Check what's wrong
make lint
```

### Commit hook failing
```bash
# Reinstall hooks
cog install-hook --all
# Verify commit message format
cog verify "feat: my message"
```

## Tools

Required tools are auto-installed by `make setup`:

- **Air** - Hot reload
- **golangci-lint** - Linting
- **Cocogitto** - Conventional commits
- **GoReleaser** - Release automation

Manual installation:
```bash
# macOS
brew install air golangci-lint cocogitto goreleaser

# Other platforms
make tools
```

## CI/CD

**GitHub Actions:**
- `.github/workflows/ci.yml` - Tests, lint, build on push/PR
- `.github/workflows/release.yml` - Automated releases on tag

**Run CI checks locally:**
```bash
make ci
```

## Architecture

This project follows **Hexagonal Architecture** (Ports & Adapters):

- **Domain Layer** (`internal/domain/`) - Pure business logic, no dependencies
- **Ports Layer** (`internal/ports/`) - Interfaces defining contracts
- **Adapters Layer** (`internal/adapters/`) - Concrete implementations

This allows:
- Easy testing (mock adapters)
- Swappable implementations
- Clear separation of concerns
- Framework independence

## Resources

- [Conventional Commits](https://www.conventionalcommits.org/)
- [Cocogitto Documentation](https://docs.cocogitto.io/)
- [GoReleaser Documentation](https://goreleaser.com/)
- [Bubbletea Documentation](https://github.com/charmbracelet/bubbletea)
