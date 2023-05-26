package baker

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/dafaque/ws-gen/internal/baker/settings/common"
	"github.com/dafaque/ws-gen/internal/config"
)

func bake(lang, fp, file string, conf *config.Config) error {
	tmpl := template.New(fp)

	var funcs template.FuncMap = make(template.FuncMap)
	lc := conf.Internal.LanguageConfig
	funcs["fconv"] = lc.GetPublicFieldNameConverter()
	funcs["sconv"] = lc.GetPublicStructNameConverter()
	funcs["snake"] = common.ToSnakeCase
	funcs["get_param"] = common.MakeGetParamFunc(conf.Custom)
	funcs["tconv"] = common.MakeParseTypeFunc(
		lc.GetPublicStructNameConverter(),
		lc.GetTypeConverter(),
		lc.GetTypeWrapper(),
	)
	funcs["enc"] = common.GetEncoder(lang)
	funcs["islist"] = common.IsList
	funcs["isenum"] = common.IsEnum
	for name, fn := range lc.GetSpecialFuncs() {
		funcs[name] = fn
	}
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
