package golang

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
		settings.DataTypeUnsignedInt,
		settings.DataTypeFloat32,
		settings.DataTypeFloat64,
		settings.DataTypeString:
		return t
	case settings.DataTypeFloat:
		return settings.DataTypeFloat64
	default:
		panic("unknown type " + t)
	}
}

func wrapType(dt string, nullable, array bool) string {
	str := dt
	if nullable {
		str = "*" + str
	}
	if array {
		str = "[]" + str
	}
	return str
}
