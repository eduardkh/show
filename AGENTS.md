# AGENTS.md

## Repo Purpose

Windows network troubleshooting CLI (`show`) built in Go with Cobra.

## Fast Orientation

- Entrypoint: `cmd/show/main.go`
- Command package: `internal/cli`
- Architecture notes: `docs/ARCHITECTURE.md`
- Last migration log: `docs/SESSION_NOTES_2026-02-22.md`

## Conventions

- Keep executable entrypoints under `cmd/<binary>`.
- Keep implementation under `internal/...` unless it must be public.
- Avoid changing command names/flags without documenting in README + session notes.
- Prefer small, behavior-preserving refactors.

## Required Updates When Refactoring

If you move files or rename packages:

1. Update imports and package names.
2. Update `build.bat` and `install.bat`.
3. Update README structure/build instructions.
4. Add or append a dated file in `docs/` describing what changed and why.

## Verification Checklist

Run before finishing:

```powershell
go fmt ./...
go test ./...
go build ./cmd/show
```

If tests are absent, explicitly note that in your handoff.
