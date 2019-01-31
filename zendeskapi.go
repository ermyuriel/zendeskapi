package zendeskapi

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func authenticateRequest(r *http.Request) {

	toEncode := []byte(fmt.Sprintf("%s/token:%s", os.Getenv("ZENDESK_USER"), os.Getenv("ZENDESK_TOKEN")))
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(toEncode)))
	base64.StdEncoding.Encode(encoded, toEncode)
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", string(encoded)))
}
