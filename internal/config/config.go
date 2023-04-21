package config

import (
	"fmt"
	"os"
	"path"

	"github.com/Dafaque/ws-gen/internal/baker/settings"

	"gopkg.in/yaml.v3"
)

var availableEncodingFormats []string = []string{"json"}

type (
	Config struct {
		Models   []model           `yaml:"models"`
		Encoding string            `yaml:"encoding"`
		Init     initParams        `yaml:"init"`
		FullPath string            `yaml:"-"`
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
	initParams struct {
		Presented bool `yaml:"-"`
		Params    []struct {
			Name     string `yaml:"name"`
			Optional bool   `yaml:"optional"`
		} `yaml:"params"`
	}
	internal struct {
		GenerateMode   string
		LanguageConfig settings.LanguageSettings
	}
)

func GetConfig(spec, config string) (*Config, error) {

	specFile, err := os.ReadFile(spec)
	if err != nil {
		return nil, err
	}
	println(string(specFile), spec)
	var conf Config
	err = yaml.Unmarshal(specFile, &conf)
	if err != nil {
		return nil, err
	}
	if !checkEncodingAvailable(conf.Encoding) {
		return nil, fmt.Errorf("\"%s\" encoding format is not supported yet", conf.Encoding)
	}

	configFile, err := os.ReadFile(config)
	if err != nil {
		return nil, fmt.Errorf("err load config file %v.", err)
	}

	if err = yaml.Unmarshal(configFile, &conf.Custom); err != nil {
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
