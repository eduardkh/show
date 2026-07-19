# Module-root install and package documentation — 2026-07-19

## What changed

- Added a thin `main` package at the module root so users can install the CLI
  with `go install github.com/eduardkh/show@latest`.
- Removed the redundant `cmd/show` entrypoint; the module root is now the only
  executable entrypoint.
- Added a package comment for the root command so pkg.go.dev can show an
  overview instead of only a directory listing.
- Declared the executable entrypoints as Windows-only so package discovery uses
  the correct build context for the WMI-based CLI.
- Updated the build/install scripts, README, and architecture notes to use the
  module root as the primary entrypoint.
- Added a pkg.go.dev badge to the README.
- Added the standard MIT license under Eduard Khiaev so pkg.go.dev can identify
  the module as redistributable.

## Why

The module previously contained executable code only under `cmd/show`. As a
result, installing the module path itself failed because the root was not a Go
package, and pkg.go.dev could only display the module's directory tree.

## Publishing notes

These improvements appear on pkg.go.dev only after the changes are committed,
pushed, and published under a new semantic-version tag.
