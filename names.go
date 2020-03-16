package namecase

import (
	"bytes"
	"unicode/utf8"
	"unsafe"
)

// namecase is a namecase made up of words
type namecase []word

// parseNamecase parse the namecase string
func parseNamecase(name string) namecase {
	if name == "" {
		return nil
	}
	names := make([]word, 0, len(name)/2+1)

	wordCase := otherWordCase
	prevRuneCase := otherRuneCase
	currRuneCase := otherRuneCase
	prevWordIndex := 0
	currRuneIndex := 0

	appends := func() {
		if prevWordIndex != currRuneIndex {
			names = append(names, word{
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
		case splitRuneCase, otherRuneCase:
			if prevRuneCase != currRuneCase {
				appends()
			}
			wordCase = wordKind(currRuneCase)
		case upperRuneCase:
			if prevRuneCase != currRuneCase {
				appends()
				wordCase = wordKind(currRuneCase)
			}
		case lowerRuneCase:
			switch prevRuneCase {
			case otherRuneCase, splitRuneCase, eofRuneCase:
				appends()
				wordCase = wordKind(currRuneCase)
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
func (n namecase) FormatSnake(wordKind wordKind, split string) string {
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
func (n namecase) FormatCamel(firstKind wordKind, otherKind wordKind) string {
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
			buf.WriteString(word.convert(otherKind))
		}
	}
	bytes := buf.Bytes()
	return *(*string)(unsafe.Pointer(&bytes))
}
