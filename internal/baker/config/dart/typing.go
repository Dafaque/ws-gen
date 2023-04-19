package dart

import "wsgen/internal/baker/config"

func normalizeType(t string) string {
	switch t {
	case config.DataTypeInt8,
		config.DataTypeInt16,
		config.DataTypeInt32,
		config.DataTypeInt64,
		config.DataTypeInt:
		return "int"
	case config.DataTypeFloat,
		config.DataTypeFloat32,
		config.DataTypeFloat64:
		return "double"
	case config.DataTypeString:
		return "String"
	default:
		panic("unknown type " + t)
	}
}
