package dart

import "fmt"

var defaultString string = `
//FIXME Encoding %s not supported
class encoder {
	static dynamic encode(dynamic) => throw "decoded is not implemented";
	static dynamic decode(dynamic) => throw "decoded is not implemented";
}
`

func getEncodingPackage(e string) string {
	switch e {
	case "json":
		return "import 'dart:convert' show json;\nvar encoder = json;"
	default:
		return fmt.Sprintf(defaultString, e)
	}
}
