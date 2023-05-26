package common

import (
	"fmt"
	"log"

	"github.com/Dafaque/ws-gen/assets"
	"github.com/Dafaque/ws-gen/internal/baker/settings"
)

func GetEncoder(lang string) settings.StringOverrider {
	return func(e string) string {
		return getEncoder(lang, e)
	}
}

func getEncoder(lang, e string) string {
	f, err := assets.Templates.ReadFile(
		fmt.Sprintf("templates/codesnippets/%s/encoding/%s", lang, e),
	)
	if err != nil {
		log.Printf(
			"Warning! %s encoding code snippet for lang %s not found",
			e, lang,
		)
		f, err = assets.Templates.ReadFile(
			fmt.Sprintf("templates/codesnippets/%s/encoding/unsupported", lang),
		)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf(string(f), e)
	}
	return string(f)
}
