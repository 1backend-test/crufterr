package saxasxx

import(
	"github.com/1backend/go-client"
)

var Token string


func GetImportedHi() (string, error) {
	var ret string
	return ret, goclient.New(Token).Call("crufterr", "saxasxx", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func GetSqlExample() error {
	var ret 
	return ret, goclient.New(Token).Call("crufterr", "saxasxx", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

func PostInputExample(rect Rectangle, unit string) (Output, error) {
	var ret Output
	return ret, goclient.New(Token).Call("crufterr", "saxasxx", "POST", "/input-example", map[string]interface{}{ "rect": rect, "unit": unit }, &ret)
}

func GetHi() (string, error) {
	var ret string
	return ret, goclient.New(Token).Call("crufterr", "saxasxx", "GET", "/hi", map[string]interface{}{  }, &ret)
}

