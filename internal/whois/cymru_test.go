package whois

import "testing"

func TestParseCymruCompact(t *testing.T) {
	raw := []byte("AS | IP | AS Name\n15169 | 8.8.8.8 | GOOGLE - Google LLC, US\n")
	summary, err := ParseCymru("8.8.8.8", raw)
	if err != nil {
		t.Fatal(err)
	}
	if summary.ASN != "15169" || summary.ASName != "GOOGLE - Google LLC, US" {
		t.Fatalf("unexpected summary: %+v", summary)
	}
}

func TestParseCymruVerbose(t *testing.T) {
	raw := []byte("AS | IP | BGP Prefix | CC | Registry | Allocated | AS Name\n13335 | 1.1.1.1 | 1.1.1.0/24 | AU | apnic | 2011-08-11 | CLOUDFLARENET\n")
	summary, err := ParseCymru("1.1.1.1", raw)
	if err != nil {
		t.Fatal(err)
	}
	if summary.Prefix != "1.1.1.0/24" || summary.Country != "AU" || summary.Registry != "apnic" {
		t.Fatalf("unexpected summary: %+v", summary)
	}
}
