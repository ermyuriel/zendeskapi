package zendeskapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RelationshipRequest struct {
	Data interface{} `json:"data"`
}

type RelationshipTypeCreate struct {
	Key    string      `json:"key"`
	Source interface{} `json:"source"`
	Target interface{} `json:"target"`
}

type RelationshipRecordCreate struct {
	RelationshipType string      `json:"relationship_type"`
	Source           interface{} `json:"source"`
	Target           interface{} `json:"target"`
}

type RelationshipSearchResponse struct {
	Data  []Relationship `json:"data"`
	Links struct {
		Previous interface{} `json:"previous"`
		Next     interface{} `json:"next"`
	} `json:"links"`
}

type Relationship struct {
	ID     string `json:"id"`
	Target string `json:"target"`
	Ref    string `json:"ref"`
}

func CreateRelationshipType(source interface{}, key string, target interface{}) (error, *ErrorResponse) {
	path := "/api/custom_resources/relationship_types"

	cr := RelationshipRequest{Data: RelationshipTypeCreate{Key: key, Source: source, Target: target}}

	ts, err := bufferJSON(cr)
	if err != nil {
		return err, nil
	}
	r, _ := http.NewRequest("POST", os.Getenv("ZENDESK_URL")+path, ts)
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

func CreateRelationshipRecord(source interface{}, relationshipType string, target interface{}) (error, *ErrorResponse) {
	path := "/api/custom_resources/relationships"

	cr := RelationshipRequest{Data: RelationshipRecordCreate{RelationshipType: relationshipType, Source: fmt.Sprintf("%s", source), Target: fmt.Sprintf("%s", target)}}

	ts, err := bufferJSON(cr)

	if err != nil {
		return err, nil
	}
	r, _ := http.NewRequest("POST", os.Getenv("ZENDESK_URL")+path, ts)
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

func DeleteRelationshipRecord(id string) (error, *ErrorResponse) {
	path := fmt.Sprintf("/api/custom_resources/relationships/%s", id)

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

func ListObjectRelationships(objectID string, relationshipType string) ([]Relationship, error, *ErrorResponse) {

	path := fmt.Sprintf("/api/custom_resources/resources/%s/relationships/%s", objectID, relationshipType)

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

	rsr := &RelationshipSearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(rsr)

	if err != nil {

		return nil, err, nil
	}

	return rsr.Data, nil, nil
}

func ListRelationshipsByType(relationshipType string) ([]Relationship, error, *ErrorResponse) {

	path := fmt.Sprintf("/api/custom_resources/relationships?type=%s", relationshipType)

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

	rsr := &RelationshipSearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(rsr)

	if err != nil {

		return nil, err, nil
	}

	return rsr.Data, nil, nil
}
