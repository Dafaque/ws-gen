package baker

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Dafaque/ws-gen/internal/baker/settings/common"
	"github.com/Dafaque/ws-gen/internal/config"
)

func bake(lang, fp, file string, conf *config.Config) error {
	tmpl := template.New(fp)

	var funcs template.FuncMap = make(template.FuncMap)

	funcs["fconv"] = conf.Internal.LanguageConfig.GetPublicFieldNameConverter()
	funcs["sconv"] = conf.Internal.LanguageConfig.GetPublicStructNameConverter()
	funcs["snake"] = common.ToSnakeCase
	funcs["get_param"] = common.MakeGetParamFunc(conf.Custom)
	funcs["tconv"] = common.MakeParseTypeFunc(
		conf.Internal.LanguageConfig.GetTypeConverter(),
		conf.Internal.LanguageConfig.GetTypeWrapper(),
	)
	tmpl.Funcs(funcs)
	tmpl, errParse := tmpl.Parse(file)
	if errParse != nil {
		return errParse
	}
	{
		components := strings.Split(fp, "/")
		for idx, c := range components {
			if c == "templates" {
				components = components[idx+2:]
				break
			}
		}
		filename := components[len(components)-1]
		fnComponents := strings.Split(filename, ".")
		components[len(components)-1] = strings.Join(
			fnComponents[:len(fnComponents)-1], ".")
		fp = path.Join(components...)
		fp = path.Join(conf.FullPath, fp)
	}
	fmt.Println(fp)
	if errCreateDir := os.MkdirAll(filepath.Dir(fp), 0770); errCreateDir != nil {
		return errCreateDir
	}
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	err = tmpl.Execute(f, *conf)
	if err != nil {
		return err
	}
	return nil
}
