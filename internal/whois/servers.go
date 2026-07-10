package whois

import (
	"context"
	"fmt"
	"strings"
)

const (
	AutoHost  = "auto"
	IANAHost  = "whois.iana.org"
	CymruHost = "whois.cymru.com"
)

var KnownHosts = []struct{ Host, Registry, Description string }{
	{AutoHost, "AUTO", "Automatically select the authoritative registry"},
	{"whois.arin.net", "ARIN", "North America"},
	{"whois.ripe.net", "RIPE", "Europe, Middle East, and Central Asia"},
	{"whois.apnic.net", "APNIC", "Asia Pacific"},
	{"whois.lacnic.net", "LACNIC", "Latin America and Caribbean"},
	{"whois.afrinic.net", "AFRINIC", "Africa"},
	{CymruHost, "CYMRU", "ASN and BGP summary"},
	{IANAHost, "IANA", "Registry referral service"},
}

func DiscoverHost(ctx context.Context, client Client, query string) (string, error) {
	raw, err := client.Query(ctx, IANAHost, 43, query)
	if err != nil {
		return "", fmt.Errorf("discover registry: %w", err)
	}
	response := ParseRecords(IANAHost, raw)
	for _, record := range response.Records {
		candidate := strings.ToLower(record.First("whois", "refer"))
		if candidate == "" {
			continue
		}
		if !IsKnownRegistryHost(candidate) {
			return "", fmt.Errorf("IANA returned unsupported WHOIS referral %q", candidate)
		}
		return candidate, nil
	}
	return "", fmt.Errorf("IANA response did not contain a WHOIS referral")
}

func IsKnownRegistryHost(host string) bool {
	for _, known := range KnownHosts {
		if known.Host != AutoHost && strings.EqualFold(known.Host, host) {
			return true
		}
	}
	return false
}
func RegistryForHost(host string) string {
	for _, known := range KnownHosts {
		if strings.EqualFold(known.Host, host) {
			return known.Registry
		}
	}
	return strings.ToUpper(host)
}
