package golang

import "wsgen/internal/baker/config"

func init() {
	config.SetConfig("go", Config{})
}
