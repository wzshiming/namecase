package namecase

import (
	"bytes"
	"unicode/utf8"
	"unsafe"
)

// Name is a name made up of words
type Name []Word

// ParseName parse the name string
func ParseName(name string) Name {
	if name == "" {
		return nil
	}
	names := make(Name, 0, len(name)/2+1)

	wordCase := ohterWordCase
	prevRuneCase := ohterRuneCase
	currRuneCase := ohterRuneCase
	prevWordIndex := 0
	currRuneIndex := 0

	appends := func() {
		if prevWordIndex != currRuneIndex {
			names = append(names, Word{
				kind: wordCase,
				word: name[prevWordIndex:currRuneIndex],
			})
			prevWordIndex = currRuneIndex
		}
	}

	for currRuneIndex != len(name) {
		r, s := utf8.DecodeRuneInString(name[currRuneIndex:])
		currRuneCase = getKind(r)

		switch currRuneCase {
		case ohterRuneCase:
			if prevRuneCase != ohterRuneCase {
				appends()
			}
			wordCase = ohterWordCase
		case splitRuneCase:
			if prevRuneCase != splitRuneCase {
				appends()
			}
			wordCase = splitWordCase
		case upperRuneCase:
			if prevRuneCase != upperRuneCase {
				appends()
				wordCase = upperWordCase
			}
		case lowerRuneCase:
			switch prevRuneCase {
			case ohterRuneCase, splitRuneCase, eofRuneCase:
				appends()
				wordCase = lowerWordCase
			}
			switch wordCase {
			case upperWordCase:
				wordCase = titleWordCase
			}
		}
		prevRuneCase = currRuneCase
		currRuneIndex += s
	}
	appends()
	return names
}

// FormatSnake returns format to snake
func (n Name) FormatSnake(wordKind wordKind, split string) string {
	if len(n) == 0 {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	for _, word := range n {
		if word.kind == splitWordCase {
			continue
		}

		if buf.Len() != 0 {
			buf.WriteString(split)
		}
		buf.WriteString(word.convert(wordKind))
	}
	bytes := buf.Bytes()
	return *(*string)(unsafe.Pointer(&bytes))
}

// FormatCamel returns format to camel
func (n Name) FormatCamel(firstKind wordKind, ohterKind wordKind) string {
	if len(n) == 0 {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	for _, word := range n {
		if word.kind == splitWordCase {
			continue
		}

		if buf.Len() == 0 {
			buf.WriteString(word.convert(firstKind))
		} else {
			buf.WriteString(word.convert(ohterKind))
		}
	}
	bytes := buf.Bytes()
	return *(*string)(unsafe.Pointer(&bytes))
}
