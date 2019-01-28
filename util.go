package zendeskapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
)

func printPrettyResponse(r *http.Response) {

	d, _ := httputil.DumpResponse(r, true)

	log.Println(string(d))

}

func printPrettyStruct(d interface{}) {

	jsn, _ := json.MarshalIndent(d, " ", "")

	log.Println(string(jsn))

}
