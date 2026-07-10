package whois

import "strings"

func Normalize(query string, response Response) Summary {
	s := Summary{Query: query, Registry: RegistryForHost(response.Server), Server: response.Server}
	for _, r := range response.Records {
		fill(&s.Range, r.First("netrange", "inetnum"))
		fill(&s.Prefix, r.First("cidr", "route", "route6", "inetrev"))
		fill(&s.Name, r.First("netname"))
		fill(&s.Organization, r.First("organization", "orgname", "org-name", "owner"))
		fill(&s.Country, r.First("country"))
		fill(&s.ASN, r.First("originas", "origin", "aut-num"))
		fill(&s.ASName, r.First("as-name"))
		fill(&s.AbuseEmail, r.First("orgabuseemail", "abuse-mailbox"))
	}
	s.ASN = strings.TrimSpace(strings.TrimPrefix(strings.ToUpper(s.ASN), "AS"))
	return s
}

func fill(destination *string, value string) {
	if *destination == "" && value != "" {
		*destination = value
	}
}
