package whois

import "strings"

type Field struct{ Name, Value string }
type Record struct{ Fields []Field }

func (r Record) First(names ...string) string {
	for _, name := range names {
		for _, field := range r.Fields {
			if strings.EqualFold(field.Name, name) && field.Value != "" {
				return field.Value
			}
		}
	}
	return ""
}

type Response struct {
	Server  string
	Raw     []byte
	Records []Record
}

type Summary struct {
	Query, Range, Prefix, Name, Organization, Country string
	ASN, ASName, AbuseEmail, Registry, Server         string
}

func (s *Summary) MergeMissing(other Summary) {
	if s.Range == "" {
		s.Range = other.Range
	}
	if s.Prefix == "" {
		s.Prefix = other.Prefix
	}
	if s.Name == "" {
		s.Name = other.Name
	}
	if s.Organization == "" {
		s.Organization = other.Organization
	}
	if s.Country == "" {
		s.Country = other.Country
	}
	if s.ASN == "" {
		s.ASN = other.ASN
	}
	if s.ASName == "" {
		s.ASName = other.ASName
	}
	if s.AbuseEmail == "" {
		s.AbuseEmail = other.AbuseEmail
	}
}
