package main

import "time"

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

type SearchUserResponse struct {
	Users        []User      `json:"users"`
	NextPage     interface{} `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Count        int         `json:"count"`
}

type UserResponse struct {
	User User `json:"user"`
}
type RelationshipCreate struct {
	Data Relationship `json:"data"`
}

type Relationship struct {
	Key    string      `json:"key"`
	Source interface{} `json:"source"`
	Target interface{} `json:"target"`
}

type ErrorResponse struct {
	Errors []struct {
		Code   string `json:"code"`
		Status string `json:"status"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"errors"`
}

type ObjectRecordCreate struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

type ObjectRequest struct {
	Data ObjectRecordCreate `json:"data"`
}

type ObjectResponse struct {
	Data struct {
		ID          string      `json:"id"`
		Type        string      `json:"type"`
		TypeVersion int         `json:"type_version"`
		Attributes  interface{} `json:"attributes"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
	} `json:"data"`
}