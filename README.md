# show

[![Go Reference](https://pkg.go.dev/badge/github.com/eduardkh/show.svg)](https://pkg.go.dev/github.com/eduardkh/show)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

`show` is a Windows-focused network troubleshooting CLI built with Cobra.

## Requirements

- Go 1.26.0+
- Windows (commands rely on WMI and Windows DNS cache classes)

## Build

```powershell
go build -o show.exe .
```

## Install (local repo)

```powershell
go install .
```

## Install (from GitHub)

```powershell
go install github.com/eduardkh/show@latest
```

## Shell Completion (PowerShell)

```powershell
show completion powershell | Out-String | Invoke-Expression
```

To make completion permanent, add the same command to your `$PROFILE`.

## Project Structure

```text
main.go                 # module-root entrypoint used by the short install path
doc.go                  # package documentation shown by pkg.go.dev
internal/cli/*.go       # Cobra command tree and handlers
internal/whois/*.go     # WHOIS transport, referral, parsing, normalization
build.bat               # local build helper
install.bat             # local install + completion helper
docs/                   # architecture and session notes
```

The command implementation is organized with a thin module-root entrypoint:

- the module-root entrypoint enables the shortest `go install` command
- `internal/...` for non-public implementation details

## Common Usage

```powershell
show ip external
show ip interface brief
show ip calc 192.168.1.1/25
show ip whois 8.8.8.8
show ip whois --raw 8.8.8.8
show ip whois --host whois.arin.net 8.8.8.8
show ip info 8.8.8.8
show ip abuseipdb 8.8.8.8
show webhook
show timestamp --epoch
show dnscache
```

`show ip whois` selects the authoritative Regional Internet Registry through
IANA by default. The `--host` override supports shell completion; use
`show ip whois --host <Tab>` after enabling completion. `--raw` prints the
registry response without normalization.

## Development

```powershell
go fmt ./...
go test ./...
go build .

# install
go install .
```

See `docs/ARCHITECTURE.md` and `docs/SESSION_NOTES_2026-02-22.md` for maintainers and future agent sessions.

```powershell
# example release workflow
git add .
git commit -m "Add root install path and package metadata"

git tag -a v0.1.0 -m "Release v0.1.0"

git push origin HEAD
git push origin v0.1.0

go install github.com/eduardkh/show@v0.1.0
go install github.com/eduardkh/show@latest
```
