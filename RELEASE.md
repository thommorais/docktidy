# Release Process

This document describes how to release new versions of docktidy.

## Overview

We use:
- **Cocogitto** for version management and changelog generation
- **GoReleaser** for building cross-platform binaries
- **GitHub Actions** for automated releases

## Version Bumping

Cocogitto automatically determines the next version based on conventional commits:

- `fix:` commits → patch version (0.0.x)
- `feat:` commits → minor version (0.x.0)
- `BREAKING CHANGE:` or `!` → major version (x.0.0)

### Creating a Release

1. Ensure all commits follow conventional commits format
2. Run the version bump command:

```bash
# Automatic version bump based on commits
cog bump --auto

# Or specify the version type explicitly
cog bump --patch   # 0.1.0 -> 0.1.1
cog bump --minor   # 0.1.0 -> 0.2.0
cog bump --major   # 0.1.0 -> 1.0.0
```

This will:
- Calculate the next version
- Update CHANGELOG.md
- Create a git tag (e.g., `v0.1.0`)
- Push the tag to GitHub (triggers release workflow)

3. GitHub Actions will automatically:
   - Build binaries for Linux, macOS, and Windows (amd64 & arm64)
   - Generate release notes from conventional commits
   - Create a GitHub release with downloadable assets
   - Generate checksums

## Manual Release (if needed)

If you need to create a release manually:

```bash
# Create and push a tag
git tag v0.1.0
git push origin v0.1.0
```

The GitHub Actions workflow will handle the rest.

## Testing Releases Locally

Test the GoReleaser configuration without publishing:

```bash
# Install goreleaser
brew install goreleaser

# Build snapshot (no tag required)
goreleaser release --snapshot --clean

# Check the dist/ directory for built binaries
ls -la dist/
```

## First Release

For the initial release:

```bash
# Create the first tag
cog bump --auto --version 0.1.0

# Or manually
git tag v0.1.0
git push origin v0.1.0
```

## Release Artifacts

Each release includes:

- **Binaries** for:
  - Linux (amd64, arm64)
  - macOS (amd64, arm64)
  - Windows (amd64, arm64)
- **Archives** (.tar.gz for Unix, .zip for Windows)
- **Checksums** (SHA256)
- **Changelog** (auto-generated from commits)

## Homebrew Tap

Releases are automatically published to a Homebrew tap (requires setup):

```bash
brew tap thommorais/tap
brew install docktidy
```

Note: This requires creating a `homebrew-tap` repository and setting the `HOMEBREW_TAP_GITHUB_TOKEN` secret.

## CI/CD Workflows

### Release Workflow
- **Trigger**: Tag push matching `v*.*.*`
- **Runs**: GoReleaser to build and publish release

### CI Workflow
- **Trigger**: Push to main or PRs
- **Runs**: Tests, linting, build verification, commit verification

## Troubleshooting

### Release failed
- Check GitHub Actions logs
- Verify tag format matches `v*.*.*`
- Ensure conventional commits are valid: `cog check`

### Version already exists
- Check existing tags: `git tag -l`
- Use `--pre-release` for pre-releases: `cog bump --auto --pre-release`

### GoReleaser errors
- Test locally: `goreleaser release --snapshot --clean`
- Check `.goreleaser.yml` syntax
