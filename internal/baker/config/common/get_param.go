package common

import "wsgen/internal/baker/config"

func MakeGetParamFunc(cfg map[string]string) config.StringOverrider {
	return func(key string) string {
		return cfg[key]
	}
}
