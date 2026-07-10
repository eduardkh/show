package cli

import (
	"strings"
	"testing"
)

func TestWhoisHostUsesLinuxCompatibleShorthand(t *testing.T) {
	whoisCmd.InitDefaultHelpFlag()
	host := whoisCmd.Flags().ShorthandLookup("h")
	if host == nil || host.Name != "host" {
		t.Fatalf("-h resolves to %v, want host flag", host)
	}
	help := whoisCmd.Flags().Lookup("help")
	if help == nil || help.Shorthand != "" {
		t.Fatalf("help shorthand = %q, want none", help.Shorthand)
	}
}

func TestCompleteWhoisHostIncludesAutoAndRegistries(t *testing.T) {
	completions, _ := completeWhoisHost(nil, nil, "")
	joined := strings.Join(completions, "\n")
	for _, expected := range []string{"auto\t", "whois.arin.net\t", "whois.cymru.com\t"} {
		if !strings.Contains(joined, expected) {
			t.Errorf("completion does not include %q", expected)
		}
	}
}
