package namecase

type runeKind uint8

const (
	_ runeKind = iota
	lowerRuneCase
	upperRuneCase
	splitRuneCase
	ohterRuneCase
	eofRuneCase
)

func getKind(r rune) runeKind {
	switch r {
	case 0:
		return eofRuneCase
	case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return upperRuneCase
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
		return lowerRuneCase
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return ohterRuneCase
	case ' ', '-', '_':
		return splitRuneCase
	default:
		return ohterRuneCase
	}
}

type wordKind uint8

const (
	_ wordKind = iota
	lowerWordCase
	upperWordCase
	titleWordCase
	initialismsWordCase
	ohterWordCase
)
