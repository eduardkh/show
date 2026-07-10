package whois

import "strings"

// ParseRecords parses the colon-delimited object format used by the RIRs.
func ParseRecords(server string, raw []byte) Response {
	response := Response{Server: server, Raw: raw}
	var record Record
	flush := func() {
		if len(record.Fields) > 0 {
			response.Records = append(response.Records, record)
			record = Record{}
		}
	}
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	for _, line := range strings.Split(text, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			continue
		}
		if strings.HasPrefix(trimmed, "%") || strings.HasPrefix(trimmed, "#") {
			continue
		}
		name, value, found := strings.Cut(line, ":")
		if found && strings.TrimSpace(name) != "" && !strings.ContainsAny(strings.TrimSpace(name), " \t") {
			record.Fields = append(record.Fields, Field{Name: strings.ToLower(strings.TrimSpace(name)), Value: strings.TrimSpace(value)})
			continue
		}
		if len(record.Fields) > 0 && len(line) > 0 && (line[0] == ' ' || line[0] == '\t') {
			last := &record.Fields[len(record.Fields)-1]
			last.Value = strings.TrimSpace(last.Value + " " + trimmed)
		}
	}
	flush()
	return response
}
