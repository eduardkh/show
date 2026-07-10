# show

`show` is a Windows-focused network troubleshooting CLI built with Cobra.

## Requirements

- Go 1.26.0+
- Windows (commands rely on WMI and Windows DNS cache classes)

## Build

```powershell
go build -o show.exe ./cmd/show
```

## Install (local repo)

```powershell
go install ./cmd/show
```

## Install (from GitHub)

```powershell
go install github.com/eduardkh/show/cmd/show@latest
```

## Shell Completion (PowerShell)

```powershell
show completion powershell | Out-String | Invoke-Expression
```

To make completion permanent, add the same command to your `$PROFILE`.

## Project Structure

```text
cmd/show/main.go        # binary entrypoint
internal/cli/*.go       # Cobra command tree and handlers
internal/whois/*.go     # WHOIS transport, referral, parsing, normalization
build.bat               # local build helper
install.bat             # local install + completion helper
docs/                   # architecture and session notes
```

This layout follows standard Go CLI conventions:
- `cmd/<binary>` for application entrypoints
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
go build ./cmd/show

# install
go install ./cmd/show
```

See `docs/ARCHITECTURE.md` and `docs/SESSION_NOTES_2026-02-22.md` for maintainers and future agent sessions.

