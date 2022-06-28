package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


// book controllers need Unmarshal data. 
// this code will help with that
// especially for the createBook function

func ParseBody(r * http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err ==nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}