//go:build windows

// Command show is a Windows network troubleshooting CLI.
//
// Install the latest release with:
//
//	go install github.com/eduardkh/show@latest
//
// Show provides commands for inspecting network interfaces, routes, DNS cache,
// public and local IP information, WHOIS records, OUI data, and timestamps.
// Run "show --help" after installation to see the complete command tree.
package main
