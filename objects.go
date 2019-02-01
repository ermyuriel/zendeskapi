package zendeskapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ObjectRecordCreate struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

type ObjectTypeCreate struct {
	Key    string       `json:"key"`
	Schema ObjectSchema `json:"schema"`
}

type ObjectSchema struct {
	Properties interface{}
	Required   []string
}

type ObjectRequest struct {
	Data interface{} `json:"data"`
}

type ObjectResponse struct {
	Data Object `json:"data"`
}

type Object struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	TypeVersion int         `json:"type_version"`
	Attributes  interface{} `json:"attributes"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type ObjectSearchResponse struct {
	Data  []Object `json:"data"`
	Links struct {
		Previous interface{} `json:"previous"`
		Next     interface{} `json:"next"`
	} `json:"links"`
}

func CreateObjectRecord(t string, attributes interface{}) (*ObjectResponse, error, *ErrorResponse) {

	path := "/api/custom_resources/resources"

	o := ObjectRecordCreate{Type: t, Attributes: attributes}
	or := ObjectRequest{Data: o}

	ts, err := bufferJSON(or)

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

	if resp.StatusCode != 201 {
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

func CreateObjectType(structure interface{}) (error, *ErrorResponse) {
	path := "/api/custom_resources/resource_types"
	schema := StructToSchema(structure)

	r, _ := http.NewRequest("POST", os.Getenv("ZENDESK_URL")+path, bytes.NewBuffer([]byte(schema)))
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return err, nil
	}

	if resp.StatusCode != 201 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, er

	}

	return nil, nil

}

func ListObjectsByType(objectType string) ([]Object, error, *ErrorResponse) {

	path := fmt.Sprintf("/api/custom_resources/resources?type=%s", objectType)

	r, err := http.NewRequest("GET", os.Getenv("ZENDESK_URL")+path, nil)
	if err != nil {
		log.Println(err)

		return nil, err, nil
	}
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err, nil
	}

	if resp.StatusCode != 200 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, nil, er

	}

	rsr := &ObjectSearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(rsr)

	if err != nil {

		return nil, err, nil
	}

	return rsr.Data, nil, nil
}

func DeleteObjectRecord(id string) (error, *ErrorResponse) {
	path := fmt.Sprintf("/api/custom_resources/resources/%s", id)

	r, _ := http.NewRequest("DELETE", os.Getenv("ZENDESK_URL")+path, nil)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return err, nil
	}

	if resp.StatusCode != 204 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, er

	}

	return nil, nil

}
