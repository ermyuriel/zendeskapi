package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	godotenv.Load()

}

func main() {

	/* u := UserCreate{Name: "test_" + time.Now().Format("2006_01_02_03.04.05"), Email: time.Now().Format("2006_01_02_03.04.05") + "@luuna.mx", Verified: true}

	printPrettyStruct(u)
	CreateUser(&u)
	m := map[string]interface{}{"id": "1", "name": "cosa"}
	CreateObjectRecord("test_object", m)

	*/

}

func authenticateRequest(r *http.Request) {

	toEncode := []byte(fmt.Sprintf("%s/token:%s", os.Getenv("ZENDESK_USER"), os.Getenv("ZENDESK_TOKEN")))
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(toEncode)))
	base64.StdEncoding.Encode(encoded, toEncode)
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", string(encoded)))
}

func CreateUser(u *UserCreate) (*UserResponse, error, *ErrorResponse) {

	cu := UserRequest{User: *u}
	ts, err := bufferJSON(cu)
	if err != nil {
		return nil, err, nil
	}
	path := "/api/v2/users.json"

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

	cur := &UserResponse{}

	_ = json.NewDecoder(resp.Body).Decode(cur)

	return cur, nil, nil

}

func GetUser(id int64) (*UserResponse, error, *ErrorResponse) {

	path := fmt.Sprintf("/api/v2/users/%v.json", id)

	r, _ := http.NewRequest("GET", os.Getenv("ZENDESK_URL")+path, nil)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err, nil
	}

	cur := &UserResponse{}

	err = json.NewDecoder(resp.Body).Decode(cur)

	if err != nil {

		return nil, err, nil
	}

	if resp.StatusCode != 201 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, nil, er

	}

	return cur, err, nil
}

func SearchUser(searchValue string) ([]User, error, *ErrorResponse) {

	path := fmt.Sprintf("/api/v2/users/search.json?query=%s", url.QueryEscape(searchValue))

	r, _ := http.NewRequest("GET", os.Getenv("ZENDESK_URL")+path, nil)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	printPrettyRequest(r)

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err, nil
	}

	printPrettyResponse(resp)

	if resp.StatusCode != 200 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, nil, er

	}

	sur := &SearchUserResponse{}

	err = json.NewDecoder(resp.Body).Decode(sur)

	if err != nil {

		return nil, err, nil
	}

	return sur.Users, nil, nil
}

func UpdateUser(searchValue string, newData *User) (error, *ErrorResponse) {

	targets, err, _ := SearchUser(searchValue)

	if err != nil {
		return err, nil

	}

	if len(targets) == 0 {
		return errors.New("No users found"), nil
	}

	if len(targets) > 1 {
		return errors.New("Ambiguous user search"), nil
	}

	return nil, nil

}

func CreateRelationshipType(source interface{}, key string, target interface{}) (error, *ErrorResponse) {
	path := "/api/custom_resources/relationship_types"

	cr := RelationshipCreate{Data: Relationship{Key: key, Source: source, Target: target}}

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

	if resp.StatusCode != 200 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, er

	}

	return nil, nil

}

func SetRelationship(source interface{}, key string, target interface{}) (error, *ErrorResponse) {
	path := "/api/custom_resources/relationship_types"

	cr := RelationshipCreate{Data: Relationship{Key: key, Source: source, Target: target}}

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

func CreateObjectRecord(t string, attributes map[string]interface{}) (*ObjectResponse, error, *ErrorResponse) {

	path := "/api/custom_resources/resources"

	j, err := json.Marshal(attributes)
	if err != nil {

		return nil, err, nil
	}
	log.Println(string(j))
	o := ObjectRecordCreate{Type: t, Attributes: string(j)}
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
