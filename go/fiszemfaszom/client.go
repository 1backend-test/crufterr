package fiszemfaszom

import(
	"github.com/1backend/go-client"
)

var Token string


func GetHi() (string, error) {
	var ret string
	return ret, goclient.New(Token).Call("crufterr", "fiszemfaszom", "GET", "/hi", map[string]interface{}{  }, &ret)
}

func GetImportedHi() (string, error) {
	var ret string
	return ret, goclient.New(Token).Call("crufterr", "fiszemfaszom", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func PostInputExample(rect Rectangle, unit string) (Output, error) {
	var ret Output
	return ret, goclient.New(Token).Call("crufterr", "fiszemfaszom", "POST", "/input-example", map[string]interface{}{ "rect": rect, "unit": unit }, &ret)
}

func GetSqlExample() error {
	var ret 
	return ret, goclient.New(Token).Call("crufterr", "fiszemfaszom", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

