# Architecture and Conventions

## Overview

This repository is a Go CLI application using Cobra.

- Binary entrypoint lives in `cmd/show/main.go`.
- CLI implementation lives in `internal/cli`.
- Internal packages are intentionally non-importable from external modules.

## Why this structure

The project was migrated from a flat root `main.go` + `cmd` package to a more idiomatic layout:

- `cmd/<binary>` is the conventional place for executable entrypoints.
- `internal/<pkg>` keeps implementation private and clearer for larger growth.
- This improves discoverability for contributors and automation agents.

## Command wiring model

- `internal/cli/root.go` defines `rootCmd` and `Execute()`.
- Each command file (`ip.go`, `external.go`, etc.) registers itself in `init()`.
- Parent-child relationships are defined through `AddCommand(...)` calls.

## Current package boundaries

- `internal/cli`: command definitions + command-level helpers.

Potential future refinement:
- Move reusable logic (HTTP clients, WMI readers, parsers) into `internal/netinfo`, `internal/whois`, etc.
- Keep `internal/cli` focused on argument parsing and output formatting.

## Build and install

- Build: `go build -o show.exe ./cmd/show`
- Local install: `go install ./cmd/show`
- Remote install: `go install github.com/eduardkh/show@latest`

## Go version policy

- Module target: Go `1.23` (`go.mod`).
- Keep this in sync with CI/dev tooling when upgrading.

## Agent working notes

When making structural changes in future sessions:

1. Prefer preserving command behavior while moving files.
2. Update scripts (`build.bat`, `install.bat`) together with layout changes.
3. Run `go fmt ./...` and `go build ./cmd/show` before concluding.
4. Log migration details in a dated docs file under `docs/`.
