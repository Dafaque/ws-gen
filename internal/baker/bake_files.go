package baker

import (
	"io/fs"
	"path"
	"wsgen/assets"
	"wsgen/internal/config"
)

const (
	genModeAll    = "all"
	genModeServer = "server"
	genModeClient = "client"
)

func BakeFiles(lang, currentPath string, fileSystem fs.ReadDirFS, conf *config.Config) error {
	dir, e := fileSystem.ReadDir(currentPath)
	if e != nil {
		return e
	}
	genMode := conf.Internal.GenerateMode
	skipCheckGenMode := genMode == genModeAll

	for _, en := range dir {
		enName := en.Name()
		fp := path.Join(currentPath, enName)
		if en.IsDir() {
			if !skipCheckGenMode {
				switch enName {
				case genModeServer, genModeClient:
					if genMode != enName {
						continue
					}
				}
			}
			err := BakeFiles(lang, fp, fileSystem, conf)
			if err != nil {
				return err
			}
			continue
		}
		file, err := assets.Templates.ReadFile(fp)
		if err != nil {
			return err
		}
		if err := bake(lang, fp, string(file), conf); err != nil {
			return err
		}
	}
	return nil
}
