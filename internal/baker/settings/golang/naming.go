package golang

import (
	"strings"
	"unicode"
)

func normaizeName(n string) string {
	if strings.ToLower(n) == "id" {
		return "ID"
	}
	var b strings.Builder
	b.WriteRune(unicode.ToUpper(rune(n[0])))
	b.WriteString(n[1:])
	return b.String()
}
