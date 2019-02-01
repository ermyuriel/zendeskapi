package zendeskapi

import (
	"encoding/json"
	"net/http"
	"os"
)

type ProfileRequest struct {
	Data Profile `json:"profile"`
}

type Profile struct {
	Source      string      `json:"source"`
	Type        string      `json:"type"`
	Identifiers interface{} `json:"identifiers"`
	Attributes  interface{} `json:"attributes"`
}

func CreateProfile(source, typ string, identifiers, attributes interface{}) (*ObjectResponse, error, *ErrorResponse) {

	path := "/api/cdp/v2/profile"

	pr := ProfileRequest{Profile{Source: source, Type: typ, Identifiers: identifiers, Attributes: attributes}}
	ts, err := bufferJSON(pr)

	if err != nil {
		return nil, err, nil
	}
	r, _ := http.NewRequest("POST", os.Getenv("ZENDESK_URL")+path, ts)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err, nil
	}

	if resp.StatusCode != 202 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, nil, er

	}

	orr := &ObjectResponse{}

	err = json.NewDecoder(resp.Body).Decode(orr)

	if err != nil {

		return nil, err, nil
	}

	return orr, nil, nil

}
