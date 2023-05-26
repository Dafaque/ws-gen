package golang

import (
	"fmt"
)

var defaultString string = `
import "errors"
//FIXME Encoding %s not supported
type unimplementedEncoder struct{}
func (ue unimplementedEncoder) Unmarshal([]byte, interface{}) error {
	return errors.New("Unmarshal() is not emplemented")
}
func (ue unimplementedEncoder) Marshal(interface{}) ([]byte, error) {
	return nil, errors.New("Marshal() is not emplemented")
}
var encoder unimplementedEncoder
`

func getEncodingPackage(e string) string {
	switch e {
	case "json":
		return "import encoder \"encoding/json\""
	case "msgpack":
		return "import encoder \"github.com/vmihailenco/msgpack/v5\""
	default:
		return fmt.Sprintf(defaultString, e)
	}
}
