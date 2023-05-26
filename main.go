package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime/debug"

	"github.com/dafaque/ws-gen/assets"
	"github.com/dafaque/ws-gen/internal/baker"
	"github.com/dafaque/ws-gen/internal/baker/settings"
	_ "github.com/dafaque/ws-gen/internal/baker/settings/dart"
	_ "github.com/dafaque/ws-gen/internal/baker/settings/golang"
	"github.com/dafaque/ws-gen/internal/config"
)

var (
	flagLang     string
	flagGenerate string
	flagSpec     string
	flagConfig   string

	flagVer bool
)

func main() {
	flag.StringVar(&flagLang, "l", "undefined", "target language")
	flag.StringVar(&flagGenerate, "g", "all", "which source files generate: client, server, all")
	flag.StringVar(&flagSpec, "s", "wsgen.spec.yml", "path to spec file")
	flag.StringVar(&flagConfig, "c", "wsgen.config.yml", "path to config file")
	flag.BoolVar(&flagVer, "v", false, "show version")
	flag.Parse()
	if flagVer {
		info, _ := debug.ReadBuildInfo()
		println(info.Main.Path, info.Main.Version)
		println(info.GoVersion)
		println(info.Main.Sum)
		return
	}
	cfg, err := config.GetConfig(flagSpec, flagConfig)
	if err != nil {
		panic(err)
	}

	templatesPath := path.Join("templates", flagLang)

	if err := os.RemoveAll(cfg.FullPath); err != nil {
		panic(err)
	}

	langConfig, exists := settings.GetConfig(flagLang)
	if !exists {
		panic(fmt.Errorf("no config for language %s", flagLang))
	}

	cfg.Internal.GenerateMode = flagGenerate
	cfg.Internal.LanguageConfig = langConfig

	if err := baker.BakeFiles(flagLang, templatesPath, assets.Templates, cfg); err != nil {
		panic(err)
	}
	println(langConfig.GetCompleteMessage())
}
