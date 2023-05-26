package common

import "github.com/dafaque/ws-gen/internal/baker/settings"

func MakeGetParamFunc(cfg map[string]string) settings.StringOverrider {
	return func(key string) string {
		return cfg[key]
	}
}
