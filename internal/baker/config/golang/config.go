package golang

import (
	"fmt"
	"wsgen/internal/baker/config"
)

type Config struct{}

func (c Config) GetTypeConverter() config.StringOverrider {
	return normalizeType
}
func (c Config) GetPublicFieldNameConverter() config.StringOverrider {
	return normaizeName
}
func (c Config) GetPublicStructNameConverter() config.StringOverrider {
	return normaizeName
}
func (c Config) GetCompleteMessage() string {
	var str string
	str += fmt.Sprintln("Run go mod tidy to install dependencies")
	return str
}
