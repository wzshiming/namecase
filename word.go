package namecase

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
