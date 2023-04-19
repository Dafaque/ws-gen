package config

import (
	"fmt"
	"os"
	"path"

	bconfig "github.com/Dafaque/wsgen/internal/baker/config"

	"gopkg.in/yaml.v3"
)

var availableEncodingFormats []string = []string{"json"}

type (
	Config struct {
		FullPath string            `yaml:"-"`
		Models   []model           `yaml:"models"`
		Encoding string            `yaml:"encoding"`
		Custom   map[string]string `yaml:"-"`
		Internal internal          `yaml:"-"`
	}
	model struct {
		Name   string  `yaml:"name"`
		Fields []field `yaml:"fields"`
	}
	field struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}
	internal struct {
		GenerateMode   string
		LanguageConfig bconfig.LanguageConfig
	}
)

func GetConfig(p string) (*Config, error) {

	cfg, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(cfg, &conf)
	if err != nil {
		return nil, err
	}
	if !checkEncodingAvailable(conf.Encoding) {
		return nil, fmt.Errorf("%s packing format is not supported yet", conf.Encoding)
	}

	dir := path.Dir(p)
	customCfg, err := os.ReadFile(path.Join(dir, "config.wsgen.yml"))
	if err != nil {
		return nil, fmt.Errorf("Warning! config.wsgen.yml not found at %s.", dir)
	}

	if err = yaml.Unmarshal(customCfg, &conf.Custom); err != nil {
		return nil, err
	}
	conf.FullPath = path.Join(conf.Custom["root"], conf.Custom["package"])

	return &conf, err
}

func checkEncodingAvailable(p string) bool {
	for _, pf := range availableEncodingFormats {
		if p == pf {
			return true
		}
	}
	return false
}
