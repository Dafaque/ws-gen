package dart

import (
	"fmt"
	"text/template"

	"github.com/dafaque/ws-gen/internal/baker/settings"
	"github.com/dafaque/ws-gen/internal/baker/settings/common"
)

type Config struct{}

func (c Config) GetTypeConverter() settings.StringOverrider {
	return convertType
}
func (c Config) GetTypeWrapper() settings.TypeWrapper {
	return wrapType
}

func (c Config) GetPublicFieldNameConverter() settings.StringOverrider {
	return common.Nop
}
func (c Config) GetPublicStructNameConverter() settings.StringOverrider {
	return normaizeStructName
}
func (c Config) GetCompleteMessage() string {
	var str string
	str += fmt.Sprintln("Now run `dart pub add web_socket_channel`")
	return str
}
func (c Config) GetSpecialFuncs() template.FuncMap {
	return template.FuncMap{
		"listtypecast": wrapListTypeCast,
		"reftoname":    refToEnumName,
	}
}
