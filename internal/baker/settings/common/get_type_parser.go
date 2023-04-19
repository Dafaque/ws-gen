package common

import (
	"strings"

	"github.com/Dafaque/ws-gen/internal/baker/settings"
)

func MakeParseTypeFunc(override settings.StringOverrider, wrap settings.TypeWrapper) settings.StringOverrider {
	return func(s string) string {
		return parseType(s, override, wrap)
	}
}

func parseType(dt string, override settings.StringOverrider, wrap settings.TypeWrapper) string {
	var raw = dt
	var nullable, array bool
	if nullable = strings.HasSuffix(dt, "?"); nullable {
		raw = raw[:len(raw)-1]
	}
	if array = strings.HasPrefix(dt, "..."); array {
		raw = raw[3:]
	}
	raw = override(raw)
	raw = wrap(raw, nullable, array)
	return raw
}
