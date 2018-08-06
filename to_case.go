package namecase

// ToUpperSpace to XX YY ZZ
func ToUpperSpace(s string) string {
	return ParseName(s).FormatSnake(upperWordCase, " ")
}

// ToLowerSpace to xx yy zz
func ToLowerSpace(s string) string {
	return ParseName(s).FormatSnake(lowerWordCase, " ")
}

// ToUpperStrike to XX-YY-ZZ
func ToUpperStrike(s string) string {
	return ParseName(s).FormatSnake(upperWordCase, "-")
}

// ToLowerStrike to xx-yy-zz
func ToLowerStrike(s string) string {
	return ParseName(s).FormatSnake(lowerWordCase, "-")
}

// ToUpperSnake to XX_YY_ZZ
func ToUpperSnake(s string) string {
	return ParseName(s).FormatSnake(upperWordCase, "_")
}

// ToLowerSnake to xx_yy_zz
func ToLowerSnake(s string) string {
	return ParseName(s).FormatSnake(lowerWordCase, "_")
}

// ToPascal to XxYyZz
func ToPascal(s string) string {
	return ToLowerHump(s)
}

// ToCamel to xxYyZz
func ToCamel(s string) string {
	return ToLowerHump(s)
}

// ToUpperHump to XxYyZz
func ToUpperHump(s string) string {
	return ParseName(s).FormatCamel(titleWordCase, titleWordCase)
}

// ToLowerHump to xxYyZz
func ToLowerHump(s string) string {
	return ParseName(s).FormatCamel(lowerWordCase, titleWordCase)
}

// ToUpperHumpInitialisms to XxYyZzID
func ToUpperHumpInitialisms(s string) string {
	return ParseName(s).FormatCamel(initialismsWordCase, initialismsWordCase)
}

// ToLowerHumpInitialisms to xxYyZzID
func ToLowerHumpInitialisms(s string) string {
	return ParseName(s).FormatCamel(lowerWordCase, initialismsWordCase)
}
