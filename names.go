package namecase

import (
	"strings"
)

// Name is a name made up of words
type Name []Word

// ParseName parse the name string
func ParseName(name string) Name {
	names := Name{}
	var currWord []rune

	prev := ohterRuneCase
	curr := ohterRuneCase

	clear := func() {
		currWord = currWord[:0]
	}

	splitCurr := func() {
		switch len(currWord) {
		case 0:
		case 1:
			wordKind := ohterWordCase
			switch getKind(currWord[0]) {
			case lowerRuneCase:
				wordKind = lowerWordCase
			case upperRuneCase:
				wordKind = upperWordCase
			}
			names = append(names, Word{
				word: string(currWord),
				kind: wordKind,
			})
			clear()
		default:
			wordKind := ohterWordCase
			switch getKind(currWord[0]) {
			case lowerRuneCase:
				wordKind = lowerWordCase
			case upperRuneCase:
				switch getKind(currWord[len(currWord)-1]) {
				case lowerRuneCase:
					wordKind = titleWordCase
				case upperRuneCase:
					wordKind = upperWordCase
				}
			}
			names = append(names, Word{
				word: string(currWord),
				kind: wordKind,
			})
			clear()
		}
	}
	splitPrev := func() {
		prevRune := currWord[len(currWord)-1]
		currWord = currWord[:len(currWord)-1]
		splitCurr()
		currWord = append(currWord, prevRune)
	}

	step := func(r rune) {
		curr = getKind(r)
		switch prev {
		case splitRuneCase:
			if curr != splitRuneCase {
				clear()
			}
		case ohterRuneCase:
			if curr != ohterRuneCase {
				splitCurr()
			}
		case lowerRuneCase:
			if curr != lowerRuneCase {
				splitCurr()
			}
		case upperRuneCase:
			switch curr {
			case lowerRuneCase:
				splitPrev()
			case ohterRuneCase, splitRuneCase, eofRuneCase:
				splitCurr()
			}
		}
		currWord = append(currWord, r)
		prev = curr
	}

	for _, r := range []rune(name) {
		step(r)
	}
	step(0)
	return names
}

func (n Name) FormatSnake(wordKind wordKind, split string) string {
	sum := make([]string, 0, len(n))
	for _, word := range n {
		sum = append(sum, word.convert(wordKind))
	}
	return strings.Join(sum, split)
}

func (n Name) FormatCamel(firstKind wordKind, ohterKind wordKind) string {
	if len(n) == 0 {
		return ""
	}
	sum := make([]string, 0, len(n))
	sum = append(sum, n[0].convert(firstKind))
	for _, word := range n[1:] {
		sum = append(sum, word.convert(ohterKind))
	}
	return strings.Join(sum, "")
}
