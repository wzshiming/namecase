package namecase

import (
	"strings"
)

// Word Is the word and kind
type Word struct {
	kind wordKind
	word string
}

func (w Word) String() string {
	return w.word
}

func (w Word) convert(kind wordKind) string {
	switch kind {
	case upperWordCase:
		return w.Upper()
	case lowerWordCase:
		return w.Lower()
	case titleWordCase:
		return w.Title()
	case initialismsWordCase:
		return w.Initialisms()
	default:
		return w.word
	}
}

// Initialisms If it's an acronym, it's all capital letters.
func (w Word) Initialisms() string {
	if word := upper(w.word); commonInitialisms[word] {
		return word
	}
	return w.Title()
}

// Title returns capitalize the first letter of other lower case.
func (w Word) Title() string {
	switch w.kind {
	case lowerWordCase, upperWordCase:
		return title(string(w.word))
	default:
		return w.word
	}
}

// Lower returns all lowercase.
func (w Word) Lower() string {
	switch w.kind {
	case upperWordCase, titleWordCase:
		return lower(string(w.word))
	default:
		return w.word
	}
}

// Upper returns all uppercase.
func (w Word) Upper() string {
	switch w.kind {
	case lowerWordCase, titleWordCase:
		return upper(string(w.word))
	default:
		return w.word
	}
}

func title(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func lower(s string) string {
	return strings.ToLower(s)
}

func upper(s string) string {
	return strings.ToUpper(s)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF":   true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}
