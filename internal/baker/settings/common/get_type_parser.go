package common

import (
	"strings"

	"github.com/Dafaque/ws-gen/internal/baker/settings"
)

func MakeParseTypeFunc(modelOverride settings.StringOverrider, override settings.StringOverrider, wrap settings.TypeWrapper) settings.StringOverrider {
	return func(s string) string {
		return parseType(s, modelOverride, override, wrap)
	}
}

func parseType(
	dt string,
	mo settings.StringOverrider,
	o settings.StringOverrider,
	w settings.TypeWrapper,
) string {
	var raw = dt
	if strings.HasPrefix(raw, "$") {
		raw = raw[1:]
		return mo(raw)
	}
	var nullable, array bool
	if nullable = strings.HasSuffix(dt, "?"); nullable {
		raw = raw[:len(raw)-1]
	}
	if array = strings.HasPrefix(dt, "..."); array {
		raw = raw[3:]
	}
	raw = o(raw)
	raw = w(raw, nullable, array)
	return raw
}
