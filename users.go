package zendeskapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type User struct {
	ID                   int64         `json:"id"`
	URL                  string        `json:"url"`
	Name                 string        `json:"name"`
	Email                string        `json:"email"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	TimeZone             string        `json:"time_zone"`
	IanaTimeZone         string        `json:"iana_time_zone"`
	Phone                interface{}   `json:"phone"`
	SharedPhoneNumber    interface{}   `json:"shared_phone_number"`
	Photo                interface{}   `json:"photo"`
	LocaleID             int           `json:"locale_id"`
	Locale               string        `json:"locale"`
	OrganizationID       interface{}   `json:"organization_id"`
	Role                 string        `json:"role"`
	Verified             bool          `json:"verified"`
	ExternalID           interface{}   `json:"external_id"`
	Tags                 []interface{} `json:"tags"`
	Alias                interface{}   `json:"alias"`
	Active               bool          `json:"active"`
	Shared               bool          `json:"shared"`
	SharedAgent          bool          `json:"shared_agent"`
	LastLoginAt          interface{}   `json:"last_login_at"`
	TwoFactorAuthEnabled bool          `json:"two_factor_auth_enabled"`
	Signature            interface{}   `json:"signature"`
	Details              interface{}   `json:"details"`
	Notes                interface{}   `json:"notes"`
	RoleType             interface{}   `json:"role_type"`
	CustomRoleID         interface{}   `json:"custom_role_id"`
	Moderator            bool          `json:"moderator"`
	TicketRestriction    string        `json:"ticket_restriction"`
	OnlyPrivateComments  bool          `json:"only_private_comments"`
	RestrictedAgent      bool          `json:"restricted_agent"`
	Suspended            bool          `json:"suspended"`
	ChatOnly             bool          `json:"chat_only"`
	DefaultGroupID       interface{}   `json:"default_group_id"`
	ReportCsv            bool          `json:"report_csv"`
	UserFields           interface{}   `json:"user_fields"`
}

type UserCreate struct {
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Phone      interface{} `json:"phone"`
	Verified   bool        `json:"verified"`
	UserFields interface{} `json:"user_fields"`
}

type UserRequest struct {
	User UserCreate `json:"user"`
}

type UserSearchResponse struct {
	Users        []User      `json:"users"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Count        int         `json:"count"`
}

type UserResponse struct {
	User User `json:"user"`
}

func CreateUser(name, email string) (*UserResponse, error, *ErrorResponse) {

	u := UserCreate{Name: name, Email: email, Verified: true}

	cu := UserRequest{User: u}
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

	resp, err := http.DefaultClient.Do(r)

	if err != nil {

		return nil, err, nil
	}

	if resp.StatusCode != 200 {
		er := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(er)
		return nil, nil, er

	}

	sur := &UserSearchResponse{}

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
