package whois

import (
	"fmt"
	"strings"
)

func ParseCymru(query string, raw []byte) (Summary, error) {
	var header []string
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		columns := splitPipe(line)
		if header == nil {
			header = columns
			continue
		}
		values := make(map[string]string, len(header))
		for i, name := range header {
			if i < len(columns) {
				values[normalizeColumn(name)] = columns[i]
			}
		}
		asn := strings.TrimSpace(strings.TrimPrefix(strings.ToUpper(values["as"]), "AS"))
		return Summary{Query: query, Prefix: values["bgp prefix"], Country: values["cc"], ASN: asn, ASName: values["as name"], Registry: values["registry"], Server: CymruHost}, nil
	}
	return Summary{}, fmt.Errorf("no Team Cymru data row received")
}

func splitPipe(line string) []string {
	parts := strings.Split(line, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
func normalizeColumn(column string) string {
	return strings.Join(strings.Fields(strings.ToLower(column)), " ")
}
