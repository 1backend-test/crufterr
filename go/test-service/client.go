package test-service

import(
	"github.com/1backend/go-client"
)

var Token string

type Test struct {
	Name	string
	Foods	[]string
}


func GetHi(howManyTimes int64, allCool bool) (string, error) {
	var ret string
	return ret, goclient.New(Token).Call("crufterr", "test-service", "GET", "/hi", map[string]interface{}{ "howManyTimes": howManyTimes, "allCool": allCool }, &ret)
}

func GetImportedHi() error {
	var ret 
	return ret, goclient.New(Token).Call("crufterr", "test-service", "GET", "/imported-hi", map[string]interface{}{  }, &ret)
}

func GetSqlExample() error {
	var ret 
	return ret, goclient.New(Token).Call("crufterr", "test-service", "GET", "/sql-example", map[string]interface{}{  }, &ret)
}

