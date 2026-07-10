# WHOIS Refactor — 2026-07-10

## What changed

- Extracted WHOIS networking and parsing from `internal/cli` into
  `internal/whois`.
- Added automatic authoritative registry discovery through IANA.
- Added `--host`/`-h`, `--port`/`-p`, `--raw`, and `--timeout` flags.
- Added shell completion for the known WHOIS servers and the default `auto`
  selection.
- Added a lossless raw response plus ordered record/field model for the
  colon-delimited formats returned by ARIN, RIPE, APNIC, LACNIC, and AFRINIC.
- Added header-driven parsing for compact and verbose Team Cymru output.
- The normalized view uses the authoritative RIR response and, when available,
  fills missing ASN/BGP fields from Team Cymru.

## Why

WHOIS registries do not share one output schema. Keeping transport, raw data,
generic record parsing, registry normalization, and CLI presentation separate
allows provider-specific behavior without forcing users to select a registry.

## Compatibility

`show ip whois 8.8.8.8` remains the primary invocation. It now discovers the
authoritative registry instead of querying only Team Cymru. Use `--raw` for the
unmodified registry response or `--host` for an explicit server override.
