package whois

import "testing"

func TestParseRecordsPreservesObjectsDuplicatesAndContinuations(t *testing.T) {
	raw := []byte("% comment\r\ninetnum: 1.1.1.0 - 1.1.1.255\r\ndescr: first\r\ndescr: second\r\naddress: line one\r\n         line two\r\n\r\nroute: 1.1.1.0/24\r\norigin: AS13335\r\n")
	response := ParseRecords("whois.apnic.net", raw)
	if len(response.Records) != 2 {
		t.Fatalf("got %d records, want 2", len(response.Records))
	}
	if got := response.Records[0].First("inetnum"); got != "1.1.1.0 - 1.1.1.255" {
		t.Fatalf("inetnum = %q", got)
	}
	if got := response.Records[0].First("address"); got != "line one line two" {
		t.Fatalf("continued address = %q", got)
	}
	if got := response.Records[1].First("origin"); got != "AS13335" {
		t.Fatalf("origin = %q", got)
	}
}

func TestNormalizeARINAndRPSLFields(t *testing.T) {
	raw := []byte("NetRange: 8.8.8.0 - 8.8.8.255\nCIDR: 8.8.8.0/24\nNetName: GOGL\nOrganization: Google LLC (GOGL)\nCountry: US\n\nOrgAbuseEmail: network-abuse@google.com\n")
	summary := Normalize("8.8.8.8", ParseRecords("whois.arin.net", raw))
	if summary.Range != "8.8.8.0 - 8.8.8.255" || summary.Prefix != "8.8.8.0/24" {
		t.Fatalf("unexpected network summary: %+v", summary)
	}
	if summary.Organization != "Google LLC (GOGL)" || summary.AbuseEmail != "network-abuse@google.com" {
		t.Fatalf("unexpected organization summary: %+v", summary)
	}
	if summary.Registry != "ARIN" {
		t.Fatalf("registry = %q", summary.Registry)
	}
}
