package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime/debug"

	"github.com/Dafaque/wsgen/assets"
	"github.com/Dafaque/wsgen/internal/baker"
	bconfig "github.com/Dafaque/wsgen/internal/baker/config"
	_ "github.com/Dafaque/wsgen/internal/baker/config/dart"
	_ "github.com/Dafaque/wsgen/internal/baker/config/golang"
	"github.com/Dafaque/wsgen/internal/config"
)

var (
	flagLang     string
	flagGenerate string
	flagSpec     string

	flagVer bool
)

func main() {
	flag.StringVar(&flagLang, "l", "undefined", "target language")
	flag.StringVar(&flagGenerate, "g", "all", "which source files generate: client, server, all")
	flag.StringVar(&flagSpec, "s", "./wsgen.yml", "path to spec file")
	flag.BoolVar(&flagVer, "v", false, "show version")
	flag.Parse()
	if flagVer {
		info, _ := debug.ReadBuildInfo()
		println(info.Main.Path, info.Main.Version)
		println(info.GoVersion)
		println(info.Main.Sum)
		return
	}
	cfg, err := config.GetConfig(flagSpec)
	if err != nil {
		panic(err)
	}

	templatesPath := path.Join("templates", flagLang)

	if err := os.RemoveAll(cfg.FullPath); err != nil {
		panic(err)
	}

	langConfig, exists := bconfig.GetConfig(flagLang)
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
