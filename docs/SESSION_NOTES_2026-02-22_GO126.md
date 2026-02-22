# Session Notes - 2026-02-22 (Go 1.26.0 Target)

## Purpose

Align module metadata and docs with the installed Go toolchain (`go1.26.0`).

## Changes made

1. Updated module Go language version
- File: `go.mod`
- Changed `go 1.23` -> `go 1.26`.

2. Pinned toolchain version
- File: `go.mod`
- Added `toolchain go1.26.0` for consistent local builds.

3. Updated documented requirement
- File: `README.md`
- Changed requirements from `Go 1.23+` to `Go 1.26.0+`.

## Verification

- `go fmt ./...`
- `go test ./...`
- `go build ./cmd/show`
