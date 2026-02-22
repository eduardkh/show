# Session Notes - 2026-02-22

## Purpose

Refactor project layout to be more idiomatic for Go, update the Go version, and document all changes for both humans and agents.

## Changes made

1. Restructured entrypoint and CLI package
- Moved `main.go` to `cmd/show/main.go`.
- Moved command files from `cmd/*.go` to `internal/cli/*.go`.
- Renamed package declarations from `package cmd` to `package cli`.
- Updated import in entrypoint to `github.com/eduardkh/show/internal/cli`.

2. Cleaned root command boilerplate
- File: `internal/cli/root.go`.
- Replaced manual `os.Exit(1)` handling with `cobra.CheckErr(rootCmd.Execute())`.
- Removed unused default `toggle` flag.

3. Updated Go module target
- File: `go.mod`.
- Changed `go 1.19` -> `go 1.23`.

4. Updated helper scripts
- File: `build.bat`: now builds explicit entrypoint with `go build -o show.exe ./cmd/show`.
- File: `install.bat`: now uses local install target `go install ./cmd/show`.

5. Rewrote top-level docs
- File: `README.md` refreshed with:
  - requirements
  - build/install commands
  - new project structure
  - common usage examples
  - development commands

6. Added persistent architecture documentation
- File: `docs/ARCHITECTURE.md` with structure rationale, boundaries, and future refactor guidance.

## Behavior compatibility

- Existing CLI command names and subcommand hierarchy were preserved.
- Refactor was layout/package-focused; command behavior was not intentionally changed.

## Follow-up suggestions

- Add tests for helper functions currently embedded in command files.
- Extract reusable logic from `internal/cli` into focused internal packages.
- Add CI workflow for `go fmt`, `go vet`, and build checks.
