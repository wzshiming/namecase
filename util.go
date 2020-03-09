package namecase

import (
	"strings"
	"unicode"
)

func title(s string) string {
	f := true
	return strings.Map(func(r rune) rune {
		if f {
			f = false
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, s)
}

func lower(s string) string {
	return strings.Map(unicode.ToLower, s)
}

func upper(s string) string {
	return strings.Map(unicode.ToUpper, s)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]struct{}{}

func init() {
	for _, i := range initialisms {
		commonInitialisms[upper(i)] = struct{}{}
		commonInitialisms[lower(i)] = struct{}{}
		commonInitialisms[title(i)] = struct{}{}
	}
}

var initialisms = []string{
	"ACL",
	"API",
	"ASCII",
	"CPU",
	"CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML",
	"HTTP",
	"HTTPS",
	"ID",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM",
	"RHS",
	"RPC",
	"SLA",
	"SMTP",
	"SQL",
	"SSH",
	"TCP",
	"TLS",
	"TTL",
	"UDP",
	"UI",
	"UID",
	"UUID",
	"URI",
	"URL",
	"UTF",
	"VM",
	"XML",
	"XMPP",
	"XSRF",
	"XSS",
}
