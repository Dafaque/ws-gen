package settings

import "text/template"

type StringOverrider func(string) string
type TypeWrapper func(t string, nullable, array bool) string

type LanguageSettings interface {
	GetTypeConverter() StringOverrider
	GetTypeWrapper() TypeWrapper
	GetPublicFieldNameConverter() StringOverrider
	GetPublicStructNameConverter() StringOverrider
	GetEncodingPackage() StringOverrider
	GetSpecialFuncs() template.FuncMap
	GetCompleteMessage() string
}

var configs map[string]LanguageSettings = make(map[string]LanguageSettings)

func GetConfig(lang string) (LanguageSettings, bool) {
	val, exists := configs[lang]
	return val, exists
}

func SetConfig(lang string, config LanguageSettings) {
	configs[lang] = config
}
