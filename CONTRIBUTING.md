# Contributing to docktidy

## Commit Message Convention

This project follows [Conventional Commits](https://www.conventionalcommits.org/).

### Format

```md
<type>(<scope>): <subject>

<body>
```

### Types

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that don't affect code meaning (formatting, etc.)
- **refactor**: Code change that neither fixes a bug nor adds a feature
- **perf**: Performance improvement
- **test**: Adding or updating tests
- **build**: Changes to build system or dependencies
- **ci**: Changes to CI configuration
- **chore**: Other changes that don't modify src or test files
- **revert**: Reverts a previous commit

### Examples

```bash
feat(tui): add resource listing screen
fix(docker): handle connection timeout gracefully
docs: update installation instructions
refactor(domain): simplify pruning logic
test(storage): add SQLite adapter tests
```

### Scope (optional)

Common scopes in this project:

- `tui` - Terminal UI
- `docker` - Docker adapter
- `storage` - Storage adapter
- `domain` - Core domain logic
- `cli` - CLI commands

### Rules

- Use lowercase for type
- Don't end subject with period
- Keep header under 72 characters
- Use imperative mood ("add" not "added")
- Body and footer are optional

## Setup

To enable commit linting locally, install [cocogitto](https://docs.cocogitto.io/):

**macOS:**

```bash
brew install cocogitto
```

**Linux/Other:**

```bash
cargo install cocogitto
```

Then install the git hooks:

```bash
cog install-hook --all
```

### Using cocogitto

```bash
# Verify a commit message
cog verify "feat: add new feature"

# Create a conventional commit interactively
cog commit

# Check commit history
cog check

# Generate changelog
cog changelog
```
