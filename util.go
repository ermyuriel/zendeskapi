package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
)

func printPrettyResponse(r *http.Response) {

	d, _ := httputil.DumpResponse(r, true)

	log.Println(string(d))

}

func printPrettyRequest(r *http.Request) {

	d, _ := httputil.DumpRequest(r, true)

	log.Println(string(d))

}
func printPrettyStruct(d interface{}) {

	jsn, _ := json.MarshalIndent(d, " ", "")

	log.Println(string(jsn))

}

func bufferJSON(s interface{}) (*bytes.Buffer, error) {

	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer([]byte(j)), nil

}
