package config

type StringOverrider func(string) string

type LanguageConfig interface {
	GetTypeConverter() StringOverrider
	GetPublicFieldNameConverter() StringOverrider
	GetPublicStructNameConverter() StringOverrider
	GetCompleteMessage() string
}

var configs map[string]LanguageConfig = make(map[string]LanguageConfig)

func GetConfig(lang string) (LanguageConfig, bool) {
	val, exists := configs[lang]
	return val, exists
}

func SetConfig(lang string, config LanguageConfig) {
	configs[lang] = config
}
