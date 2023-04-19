package dart

import (
	"strings"
	"unicode"
)

func normaizeStructName(n string) string {
	var b strings.Builder
	b.WriteRune(unicode.ToUpper(rune(n[0])))
	b.WriteString(n[1:])
	return b.String()
}
