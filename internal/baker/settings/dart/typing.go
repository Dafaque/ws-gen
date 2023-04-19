package dart

import "github.com/Dafaque/ws-gen/internal/baker/settings"

func convertType(t string) string {
	switch t {
	case settings.DataTypeInt8,
		settings.DataTypeInt16,
		settings.DataTypeInt32,
		settings.DataTypeInt64,
		settings.DataTypeInt,
		settings.DataTypeUnsignedInt8,
		settings.DataTypeUnsignedInt16,
		settings.DataTypeUnsignedInt32,
		settings.DataTypeUnsignedInt64,
		settings.DataTypeUnsignedInt:
		return "int"
	case settings.DataTypeFloat,
		settings.DataTypeFloat32,
		settings.DataTypeFloat64:
		return "double"
	case settings.DataTypeString:
		return "String"
	default:
		panic("unknown type " + t)
	}
}

func wrapType(dt string, nullable, array bool) string {
	str := dt
	if nullable {
		str += "?"
	}
	if array {
		str = "List<" + str + ">"
	}
	return str
}
