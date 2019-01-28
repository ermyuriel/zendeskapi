package zendeskapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//edouard.maire@luuna.mx

//m1YaEHA9zD2cYhDi2zkxkFp1Ro5oy8xYWHrkq9P7

//Authorization: ""

func init() {

	godotenv.Load()

}

func authenticateRequest(r *http.Request) {

	toEncode := []byte(fmt.Sprintf("%s/token:%s", os.Getenv("ZENDESK_USER"), os.Getenv("ZENDESK_TOKEN")))
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(toEncode)))
	base64.StdEncoding.Encode(encoded, toEncode)
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", string(encoded)))
}

func createUser(u *User) (*UserResponse, error) {

	cu := UserRequest{User: *u}
	j, err := json.Marshal(cu)
	if err != nil {
		return nil, err
	}
	path := "/api/v2/users.json"

	r, _ := http.NewRequest("POST", os.Getenv("ZENDESK_URL")+path, bytes.NewBuffer([]byte(j)))
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err
	}

	cur := &UserResponse{}

	_ = json.NewDecoder(resp.Body).Decode(cur)

	return cur, nil

}

func getUser(id int64) (*UserResponse, error) {

	path := fmt.Sprintf("/api/v2/users/%v.json", id)

	r, _ := http.NewRequest("GET", os.Getenv("ZENDESK_URL")+path, nil)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err
	}

	cur := &UserResponse{}

	err = json.NewDecoder(resp.Body).Decode(cur)

	if err != nil {

		return nil, err
	}

	return cur, nil
}

func searchUser(value string) ([]User, error) {

	path := fmt.Sprintf("/api/v2/users/search.json?query=%s", value)

	r, _ := http.NewRequest("GET", os.Getenv("ZENDESK_URL")+path, nil)
	authenticateRequest(r)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err
	}

	sur := &SearchUserResponse{}

	err = json.NewDecoder(resp.Body).Decode(sur)

	if err != nil {

		return nil, err
	}

	return sur.Users, nil
}

func updateUser(email string, newData *User) error {

	targets, err := searchUser(email)

	if err != nil {
		return err

	}

	if len(targets) == 0 {
		return errors.New("No users found")
	}

	if len(targets) > 1 {
		return errors.New("Ambiguous user search")
	}

	return nil

}

func createRelationshipType() {}
