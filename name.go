package nomenclature

import (
	"regexp"
	"strings"
)

var humpUpper = regexp.MustCompile(`_*[A-Z]`)

var snakeLower = regexp.MustCompile(`(_+|^)[a-z]`)

func Hump2Snake(s string) string {
	return toSnake(toHump(s))
}

func Snake2Hump(s string) string {
	return toHump(toSnake(s))
}

func toSnake(s string) string {
	return strings.ToLower(strings.Trim(humpUpper.ReplaceAllStringFunc(s, func(d string) string {
		return "_" + strings.Trim(d, "_")
	}), "_"))
}

func toHump(s string) string {
	return snakeLower.ReplaceAllStringFunc(s, func(d string) string {
		d = strings.TrimLeft(d, "_")
		if len(d) == 0 {
			return ""
		}
		return strings.ToUpper(d[:1]) + strings.ToLower(d[1:])
	})

}
