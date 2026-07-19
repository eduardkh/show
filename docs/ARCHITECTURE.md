# Architecture and Conventions

## Overview

This repository is a Go CLI application using Cobra.

- The primary binary entrypoint lives in `main.go` at the module root.
- CLI implementation lives in `internal/cli`.
- Internal packages are intentionally non-importable from external modules.

## Why this structure

The CLI implementation uses an internal package with a thin root entrypoint:

- The module-root entrypoint makes `go install github.com/eduardkh/show@latest`
  work and gives pkg.go.dev a documented root package to display.
- `internal/<pkg>` keeps implementation private and clearer for larger growth.
- This improves discoverability for contributors and automation agents.

## Command wiring model

- `internal/cli/root.go` defines `rootCmd` and `Execute()`.
- Each command file (`ip.go`, `external.go`, etc.) registers itself in `init()`.
- Parent-child relationships are defined through `AddCommand(...)` calls.

## Current package boundaries

- `internal/cli`: command definitions + command-level helpers.
- `internal/whois`: WHOIS transport, IANA referral discovery, registry response
  parsing, and normalized summaries.

Potential future refinement:
- Move other reusable logic (HTTP clients and WMI readers) into focused
  `internal/...` packages as those boundaries become useful.
- Keep `internal/cli` focused on argument parsing and output formatting.

## Build and install

- Build: `go build -o show.exe .`
- Local install: `go install .`
- Remote install: `go install github.com/eduardkh/show@latest`

## Go version policy

- Module target: Go `1.26` (`go.mod`).
- Keep this in sync with CI/dev tooling when upgrading.

## Agent working notes

When making structural changes in future sessions:

1. Prefer preserving command behavior while moving files.
2. Update scripts (`build.bat`, `install.bat`) together with layout changes.
3. Run `go fmt ./...`, `go test ./...`, and `go build .` before concluding.
4. Log migration details in a dated docs file under `docs/`.
