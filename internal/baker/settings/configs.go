package settings

type StringOverrider func(string) string
type TypeWrapper func(t string, nullable, array bool) string

type LanguageSettings interface {
	GetTypeConverter() StringOverrider
	GetTypeWrapper() TypeWrapper
	GetPublicFieldNameConverter() StringOverrider
	GetPublicStructNameConverter() StringOverrider
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
