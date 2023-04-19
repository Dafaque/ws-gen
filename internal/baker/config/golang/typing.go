package golang

import "wsgen/internal/baker/config"

func normalizeType(t string) string {
	switch t {
	case config.DataTypeInt8,
		config.DataTypeInt16,
		config.DataTypeInt32,
		config.DataTypeInt64,
		config.DataTypeInt,
		config.DataTypeFloat32,
		config.DataTypeFloat64,
		config.DataTypeString:
		return t
	case config.DataTypeFloat:
		return config.DataTypeFloat64
	default:
		panic("unknown type " + t)
	}
}
