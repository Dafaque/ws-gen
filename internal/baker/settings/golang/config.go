package golang

import (
	"fmt"
	"text/template"

	"github.com/Dafaque/ws-gen/internal/baker/settings"
)

type Config struct{}

func (c Config) GetTypeConverter() settings.StringOverrider {
	return convertType
}
func (c Config) GetTypeWrapper() settings.TypeWrapper {
	return wrapType
}
func (c Config) GetPublicFieldNameConverter() settings.StringOverrider {
	return normaizeName
}
func (c Config) GetPublicStructNameConverter() settings.StringOverrider {
	return normaizeName
}
func (c Config) GetCompleteMessage() string {
	var str string
	str += fmt.Sprintln("Run go mod tidy to install dependencies")
	return str
}
func (c Config) GetEncodingPackage() settings.StringOverrider {
	return getEncodingPackage
}
func (c Config) GetSpecialFuncs() template.FuncMap {
	return nil
}
